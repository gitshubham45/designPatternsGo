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

	key := "user1"
	for i := 0; i < 10; i++ {
		allowed := rl.Allow(key)
		fmt.Printf("Request %d allowed:%v\n", i+1, allowed)
		time.Sleep(1 * time.Second)
	}
}
