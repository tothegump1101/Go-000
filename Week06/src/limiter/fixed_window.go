package limiter

import (
	"fmt"
	"sync/atomic"
	"time"
)

type SimpleRateLimiter interface {
	Allow() bool
}

type FixedWindow struct {
	window    []int64
	maxPerSec int64
}

func (w *FixedWindow) Allow() bool {
	index := w.getIndex()
	addr := &w.window[index]
	atomic.AddInt64(addr, 1)
	return w.window[index] <= w.maxPerSec
}

func (w *FixedWindow) getIndex() int {
	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), 0, 0, 0, currentTime.Location()).Unix()
	if tickets == 0 {
		panic("ticket is zero")
	}
	ret := int((currentTime.Unix() - startTime) / tickets)
	fmt.Errorf("ret: %+v", ret)
	return ret
}

func NewFixedWindow(maxPerSec int64) *FixedWindow {
	//window := make(map[int64]int64, 60*60)
	window := make([]int64, 60*60)
	currentTime := time.Now()
	begin := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location()).Unix()
	oneHour := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 1, 0, 0, 0, currentTime.Location()).Unix()
	tickets = (oneHour - begin) / (60 * 60)
	return &FixedWindow{
		maxPerSec: maxPerSec,
		window:    window,
	}
}

var tickets int64
