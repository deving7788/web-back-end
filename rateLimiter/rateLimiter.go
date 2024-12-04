package rateLimiter

import (
    "sync"
    "time"
)

type RateLimiter struct {
    mu             sync.Mutex
    requests       map[string][]time.Time
    maxRequests    int
    interval       time.Duration
    cleanInterval  time.Duration
}

func NewRateLimiter(maxRequests int, interval time.Duration, cleanInterval time.Duration) *RateLimiter{
    rateLimiter := &RateLimiter{
        requests:       make(map[string][]time.Time),
        maxRequests:    maxRequests,
        interval:       interval,
        cleanInterval:  cleanInterval,
    }
    go rateLimiter.CleanOldRequests()
    return rateLimiter
}

func (rl *RateLimiter) CleanOldRequests() {
    ticker := time.NewTicker(rl.cleanInterval)
    for range ticker.C {
        rl.mu.Lock()
        for ip, timestamps := range rl.requests {
            var filtered []time.Time
            for _, timestamp := range timestamps {
                if time.Since(timestamp) < rl.interval {
                    filtered = append(filtered, timestamp)
                }
            }
            rl.requests[ip] = filtered
        }
        rl.mu.Unlock()
    }
}

func (rl *RateLimiter) IsAllowed(ip string) bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()

    now:= time.Now()
    rl.requests[ip] = append(rl.requests[ip], now) 

    var validRequests []time.Time
    for _, timeStamp := range rl.requests[ip] {
        if now.Sub(timeStamp) < rl.interval {
            validRequests = append(validRequests, timeStamp)
        }
    }

    rl.requests[ip] = validRequests

    return len(rl.requests[ip]) <= rl.maxRequests
}
