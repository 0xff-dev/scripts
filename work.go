/*
	使用无缓冲的通道构建池
*/
package concurrentmode

import (
	"fmt"
	"sync"
	"time"
)


type Worker interface {
	Task()
}


type Pools struct {
	work chan Worker
	wg1 sync.WaitGroup
}


func PoolsNew(maxGoroutines int) *Pools {
	p := Pools{
		work: make(chan Worker),
	}
	p.wg1.Add(maxGoroutines)
	for i := 0; i<maxGoroutines; i ++ {
		go func(){
			defer p.wg1.Done()
			for w := range p.work {
				w.Task()
			}
		}()
	}
	return &p
}


func (p *Pools) Run(w Worker) {
	p.work <- w
}

func (p *Pools) Shutdown() {
	close(p.work)
	p.wg1.Wait()
}


var names = []string {
	"lisa",
	"bob",
	"bbbb",
	"cccc",
	"dddd",
}
type namePrint struct {
	name string
}

func (np *namePrint) Task() {
	fmt.Printf("Name -> %s\n", np.name)
	time.Sleep(time.Second)
}

func TestWorker() {
	ps := PoolsNew(2)
	var wg5 sync.WaitGroup
	wg5.Add(100 * len(names))
	// 每个name会创建100哥goroutin送到ps里, 形成竞争
	for i := 0; i<100; i ++ {
		for _, name := range names {
			np := namePrint{
				name: name,
			}
			go func() {
				ps.Run(&np)
				wg5.Done()
			}()
		}
	}
	wg5.Wait()
	ps.Shutdown()
}