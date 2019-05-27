package doone

import (
	"sync"
	"sync/atomic"
)

type OnlyOne struct {
	m    sync.Mutex
	done uint32
}

func (o *OnlyOne) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 0 {
		atomic.StoreUint32(&o.done, 1)
		defer atomic.StoreUint32(&o.done, 0)
		f()
	}
}
