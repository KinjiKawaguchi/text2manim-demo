package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	limit    int
	interval time.Duration
	requests []time.Time
	mu       sync.Mutex
}

func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:    limit,
		interval: interval,
		requests: make([]time.Time, 0, limit),
	}
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if len(rl.requests) < rl.limit {
		rl.requests = append(rl.requests, now)
		return true
	}

	if now.Sub(rl.requests[0]) > rl.interval {
		rl.requests = append(rl.requests[1:], now)
		return true
	}

	return false
}
