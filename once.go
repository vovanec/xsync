package xsync

import (
	"sync"
	"sync/atomic"
)

// Once is an object that behaves exactly like sync.Once with only exception
// that it does not save state on panic or function failure.
// See http://golang.org/pkg/sync/#Once
type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func() error) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		if err := f(); err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
}
