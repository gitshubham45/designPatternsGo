package limiter

import (
	"sync"
	"time"
)

// here we are tracking count of request
type FixedWindowRateLimiter struct {
	mu sync.Mutex
	requests map[string]int // user -> request count
	windows map[string]time.Time // user -> window start time
	limit int
	window time.Duration
}

func NewFixedWindowRateLimiter(limit int,window time.Duration) *FixedWindowRateLimiter{
	return &FixedWindowRateLimiter{
		requests: make(map[string]int),
		windows: make(map[string]time.Time),
		limit: limit,
		window: window,
	}
}

func (rl *FixedWindowRateLimiter) Allow(key string) bool{
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	start , ok := rl.windows[key]

	if !ok || now.Sub(start) >= rl.window {
		rl.windows[key] = now
		rl.requests[key] = 1
		return true
	}

	if rl.requests[key] < rl.limit {
		rl.requests[key]++
		return true
	}

	return false
}