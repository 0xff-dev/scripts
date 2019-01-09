package concurrentmode

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

const timeOut = 3 * time.Second


type Runner struct {
	interrupt chan os.Signal
	complete chan error
	timeout <- chan time.Time
	tasks []func(int)
}


var ErrTimeout = errors.New("received timeout")
var ErrInterrupt = errors.New("received interrupt")


func New(d time.Duration) (*Runner) {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete: make(chan error),
		timeout: time.After(d), //  使用After返回一个time.Time类型的chan
	}
}
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}
func (r *Runner) getInterrupt() bool {
	select {
	// 从interrupt通道去接受ctrl+c的信号, 没有任何信号接收到的时候， 阻塞除非default
	case <- r.interrupt:
		signal.Stop(r.interrupt) // 停止接受后续的任何信号. 关闭通道.
		return true
	default:
		return false
	}
}
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.getInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()

	// select 多路复用，
	select {
	case  err := <- r.complete:
		return err
	case <- r.timeout:
		return ErrTimeout
	}
}


func TestRunner() {
	run := New(timeOut)
	run.Add(createTask(), createTask(), createTask())
	if err := run.Start(); err != nil {
		switch err {
		case ErrTimeout:
			log.Printf("Terminating due to timeout")
			os.Exit(1)
		case ErrInterrupt:
			log.Printf("Terminating due to interrupt")
			os.Exit(2)
		}
	}
	log.Printf("Process Done")
}


func createTask() func(int){
	return func(id int) {
		log.Printf("Processor - Task %d", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}