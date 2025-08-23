package main

import (
	"fmt"
	"time"

	"github.com/gitshubham45/designPatternGo/rateLimiter/limiter"
)

type RateLimiter interface {
	Allow(key string) bool
}

func main() {
	rl := limiter.NewSlidingWindowRateLimiter(5, 10*time.Second)

	frl := limiter.NewFixedWindowRateLimiter(10, 10*time.Second)

	key := "user1"
	for i := 0; i < 30; i++ {
		allowed := rl.Allow(key)
		allowedFrl := frl.Allow(key)
		fmt.Printf("Request %d allowed:%v\n", i+1, allowed)
		fmt.Printf("Request %d allowed from frl:%v\n", i+1, allowedFrl)
		time.Sleep(500 * time.Millisecond)
	}
}
