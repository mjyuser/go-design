package singleton

import (
	"sync"
	"sync/atomic"
)

type Singleton struct {}

var singleton *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}

type OnceDo struct {
	do int32
	sync.Mutex
}

func (o *OnceDo) Do(fn func()) {
	if atomic.LoadInt32(&o.do) == 0 {
		o.getSlow(fn)
	}
}

func (o *OnceDo) getSlow(fn func()) {
	o.Lock()
	defer o.Unlock()

	if o.do == 0 {
		defer atomic.AddInt32(&o.do, 1)
		fn()
	}
}
