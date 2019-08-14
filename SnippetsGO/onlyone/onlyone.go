package onlyone

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type OnlyOne struct {
	m    sync.Mutex
	done uint32
}

// 常见的函数签名类型
type FuncType1 func() error
type FuncType2 func() (interface{}, error)

func (o *OnlyOne) Do(f FuncType1) error {
	syncError := fmt.Errorf("We are working on the task! please try again later")
	if atomic.LoadUint32(&o.done) == 1 {
		return syncError
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		atomic.StoreUint32(&o.done, 1)
		defer atomic.StoreUint32(&o.done, 0)
		err := f()
		return err
	}
	return syncError
}

func (o *OnlyOne) DoType2(f FuncType2) (interface{}, error) {
	syncError := fmt.Errorf("We are working on the task! please try again later")
	if atomic.LoadUint32(&o.done) == 1 {
		return nil, syncError
	}
	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		atomic.StoreUint32(&o.done, 1)
		defer atomic.StoreUint32(&o.done, 0)
		out, err := f()
		return out, err
	}
	return nil, syncError
}
