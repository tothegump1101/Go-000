package limiter

import (
	"sync"
	"time"
)

type SlidingWindow interface {
	Allow() bool
	Incr(i int64)
	Sum(now time.Time) int64
}

func NewSlidingWindow(capacity int64) *slidingWindow {
	return &slidingWindow{
		capacity: capacity,
		window:   make(map[int64]*bucket),
	}
}

func (s *slidingWindow) Allow() bool {
	now := time.Now()
	return s.Sum(now) < s.capacity
}

func (s *slidingWindow) Sum(now time.Time) int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	var sum int64

	for k, b := range s.window {
		if k >= now.Unix()-10 {
			sum += b.value
		}
	}
	return sum
}

func (s *slidingWindow) Incr(i int64) {
	if i == 0 {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	b := s.getCurrent()
	b.value += 1
	s.delOldBucket()
}

type slidingWindow struct {
	capacity int64
	window   map[int64]*bucket
	mu       sync.Mutex
}

type bucket struct {
	value int64
}

func (s *slidingWindow) getCurrent() *bucket {
	now := time.Now().Unix()
	var b *bucket
	var ok bool
	if b, ok = s.window[now]; !ok {
		b = &bucket{}
		s.window[now] = b
	}
	return b
}

func (s *slidingWindow) delOldBucket() {
	now := time.Now().Unix()
	for k := range s.window {
		if k < now {
			delete(s.window, k)
		}
	}
}
