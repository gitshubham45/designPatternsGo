package limiter

import (
	"sync"
	"time"
)

type TokenBucketRateLimiter struct {
	mu           sync.Mutex
	tokens       map[string]float64
	lastRefilled map[string]time.Time
	rate         float64 // tokens per second
	capacity     float64
}

func NewTokenBucketRateLimiter(rate float64, capacity float64) *TokenBucketRateLimiter {
	return &TokenBucketRateLimiter{
		tokens:       make(map[string]float64),
		lastRefilled: make(map[string]time.Time),
		rate:         rate,
		capacity:     capacity,
	}
}

func (tb *TokenBucketRateLimiter) Allow(key string) bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	now := time.Now()
	last, ok := tb.lastRefilled[key]
	if !ok {
		last = now
		tb.tokens[key] = tb.capacity
	}

	elapsed := now.Sub(last).Seconds()
	tb.tokens[key] += elapsed * tb.rate
	if tb.tokens[key] > tb.capacity {
		tb.tokens[key] = tb.capacity
	}

	if tb.tokens[key] < 1 {
		return false
	}

	tb.tokens[key]--
	tb.lastRefilled[key] = now
	return true
}
