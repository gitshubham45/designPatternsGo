package limiter

import (
	"sync"
	"time"
)

// here we are tracking time of request
type SlidingWindowRateLimiter struct {
	requests map[string][]time.Time // key = user / IP , value = list of timestamps
	limit    int
	window   time.Duration
	mu       sync.Mutex
}

func NewSlidingWindowRateLimiter(limit int, window time.Duration) *SlidingWindowRateLimiter {
	return &SlidingWindowRateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		window:   window,
	}
}

func (rl *SlidingWindowRateLimiter) Allow(key string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-rl.window) // eg. taking last 10 second

	// filter out timestamp outside the window
	reqs := rl.requests[key]
	var filtered []time.Time
	for _ , t := range reqs {
		if t.After(windowStart){
			filtered = append(filtered, t)
		}
	}

	// check if within limit
	if len(filtered) >= rl.limit{
		return false
	}

	filtered = append(filtered, now)
	rl.requests[key] = filtered
	return true
}
