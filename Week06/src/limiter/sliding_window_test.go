package limiter

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestNewFixedWindow(t *testing.T) {
	l := NewFixedWindow(10)
	t.Log("start")
	ctx := context.Background()
	eg, _ := errgroup.WithContext(ctx)
	for i := 0; i < 30; i++ {
		eg.Go(func() error {
			run(t, l)
			return nil
		})
	}
	_ = eg.Wait()
}

func Test_slidingWindow_Allow(t *testing.T) {
	l := NewSlidingWindow(10)
	t.Error("start")
	ctx := context.Background()
	eg, _ := errgroup.WithContext(ctx)
	for i := 0; i < 30; i++ {
		eg.Go(func() error {
			run1(t, l)
			return nil
		})
	}
	_ = eg.Wait()

}

func run1(t *testing.T, l *slidingWindow) {
	for i := 0; i < 1000; i++ {
		t.Log("sum:", l.Sum(time.Now()))
		if l.Allow() {
			atomic.AddInt64(&incr, 1)
			t.Log("hello: ", incr)
		}
		time.Sleep(time.Second)
	}
}

func Test_slidingWindow_Avg(t *testing.T) {

}

func Test_slidingWindow_Incr(t *testing.T) {

}

func Test_slidingWindow_Sum(t *testing.T) {

}

func Test_slidingWindow_UpdateMax(t *testing.T) {

}

func Test_slidingWindow_delOldBucket(t *testing.T) {

}

func Test_slidingWindow_getCurrent(t *testing.T) {

}
