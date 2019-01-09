/*
   使用有缓冲的通道实现资源池,管理任意数量的goroutine之间共享及独立使用的资源。
   共享数据库连接池，内存缓冲区
 */
package concurrentmode

import (
	"errors"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Pool struct {
	mutex sync.Mutex // Lock, Unlock
	resources chan io.Closer // 资源的关闭方法？？？，多次调用就是自己的实现结果，行为未知
	factory func() (io.Closer, error)
	closed bool
}

type dbConnection struct {
	id int32
}

var ErrPoolClosed = errors.New("pool Close!")


func PoolNew(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small")
	}
	return &Pool {
		factory: fn,
		resources: make(chan io.Closer, size),
	}, nil
}
func (p *Pool) Acquire() (io.Closer, error) {
	// 在池里获取资源
	select {
	case r, ok := <- p.resources:
		log.Println("Acquire: ", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire", "New Resource")
		return p.factory()
	}
}
func (p *Pool) Release(r io.Closer) {
	// 使用一个资源后放入池中
	p.mutex.Lock()
	{
		// 池子关闭， 则销毁这个资源
		if p.closed {
			r.Close()
			return
		}
		select {
		case p.resources <- r:
			log.Println("Release: ", "In Queue")
		default:
			// 队列满了， 关闭该chan
			log.Println("Resource: ", "Closed")
			r.Close()
		}
	}
	p.mutex.Unlock()
}

func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.closed {
		return
	}
	p.closed = true
	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}


// 测试
var idCounter int32
const (
	maxGoroutine = 25
	pooledResources = 2
)

func (db *dbConnection) Close() error {
	log.Println("db connection: ", db.id)
	return nil
}
// 池子使用的工厂创建函数
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}, nil
}
func preformQuries(query int, p *Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer p.Release(conn)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("QID[%d] CID[%d]", query, conn.(*dbConnection).id)
}
func TestPool() {
	var wg sync.WaitGroup
	// goroutine数量
	wg.Add(maxGoroutine)
	p, err := PoolNew(createConnection, pooledResources)
	if err != nil {
		log.Println("Create connection fail")
	}
	for query := 0; query<maxGoroutine; query ++ {
		go func(q int){
			defer wg.Done()
			preformQuries(q, p)
		}(query)
	}
	wg.Wait()
	log.Println("db pools close")
	p.Close()
}