package limiter

import (
	"sync/atomic"
	"testing"
	"time"
)

var incr int64

func run(t *testing.T, l *FixedWindow) {
	for i := 0; i < 1000; i++ {
		if l.Allow() {
			atomic.AddInt64(&incr, 1)
			t.Log("hello: ", incr)
		}
		time.Sleep(time.Second)
	}
}
