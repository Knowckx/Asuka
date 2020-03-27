package util

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

var SyncError = fmt.Errorf("We are working on the task! please try again later")

func (o *OnlyOne) Do(f FuncType1) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return SyncError
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
	return SyncError
}

func (o *OnlyOne) DoType2(f FuncType2) (interface{}, error) {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil, SyncError
	}
	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		atomic.StoreUint32(&o.done, 1)
		defer atomic.StoreUint32(&o.done, 0)
		out, err := f()
		return out, err
	}
	return nil, SyncError
}

/* ----------------- OnlyOneMu Start ----------------- */
// 不可重入的悲观锁实现 | 不会等待,立即返回 | 基于变量的原子操作实现
const (
	StatusUnlock uint32 = 0
	StatusLock   uint32 = 1
)

type OnlyOneMu struct {
	status uint32
}

func (in *OnlyOneMu) TryLock() bool {
	return atomic.CompareAndSwapUint32(&in.status, StatusUnlock, StatusLock)
}

func (in *OnlyOneMu) Unlock() {
	atomic.StoreUint32(&in.status, StatusUnlock)
}

/* ----------------- OnlyOneMu End   ----------------- */
