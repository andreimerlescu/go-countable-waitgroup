package countable_waitgroup

import (
	"sync"
	"sync/atomic"
)

type CountableWaitGroup struct {
	wg      sync.WaitGroup
	cnt     int64
	stopped atomic.Bool
}

type IWaitGroup interface {
	Add(i int)
	Done()
	Count() int64
	IsPending() bool
	PreventAdd()
	CanAdd() bool
}

func (cwg *CountableWaitGroup) Add(i int) {
	if !cwg.CanAdd() {
		return
	}
	atomic.AddInt64(&cwg.cnt, int64(i))
	cwg.wg.Add(i)
}

func (cwg *CountableWaitGroup) Done() {
	atomic.AddInt64(&cwg.cnt, -1)
	cwg.wg.Done()
}

func (cwg *CountableWaitGroup) Count() int64 {
	return atomic.LoadInt64(&cwg.cnt)
}

func (cwg *CountableWaitGroup) Wait() {
	cwg.wg.Wait()
}

func (cwg *CountableWaitGroup) IsPending() bool {
	return atomic.LoadInt64(&cwg.cnt) > 0
}

func (cwg *CountableWaitGroup) PreventAdd() {
	cwg.stopped.Store(true)
}

func (cwg *CountableWaitGroup) CanAdd() bool {
	return !cwg.stopped.Load()
}
