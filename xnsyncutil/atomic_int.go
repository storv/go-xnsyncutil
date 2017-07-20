package xnsyncutil

import "sync/atomic"

type AtomicInt int64

// Store int value atomically
func (ai *AtomicInt) Store(v int) {
	atomic.StoreInt64((*int64)(ai), int64(v))
}

// Add int value atomically
func (ai *AtomicInt) Add(delta int) {
	atomic.AddInt64((*int64)(ai), int64(delta))
}

// Load int value atomically
func (ai *AtomicInt) Load() int {
	return int(atomic.LoadInt64((*int64)(ai)))
}

type AtomicBool AtomicInt

// Store bool value atomically
func (ab *AtomicBool) Store(v bool) {
	if v {
		(*AtomicInt)(ab).Store(1)
	} else {
		(*AtomicInt)(ab).Store(0)
	}
}

// Load bool value atomically
func (ab *AtomicBool) Load() bool {
	return (*AtomicInt)(ab).Load() != 0
}
