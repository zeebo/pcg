package pcg

import (
	"sync"
	"sync/atomic"
)

// a poor man's thread id. use the fact that sync.Pool has some affinity to return
// a counter that should stay the same between calls. it's up to you to turn this
// counter into something useful.

var (
	tidCounter uint64
	tidPool    = sync.Pool{New: func() interface{} { return atomic.AddUint64(&tidCounter, 1) }}
)

func tid() (v uint64) {
	x := tidPool.Get()
	tidPool.Put(x)
	v, _ = x.(uint64)
	return v
}
