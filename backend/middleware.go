package main

import (
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Um limiter por IP
var (
	limiters = make(map[string]*rate.Limiter)
	mu       sync.Mutex
)

func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	if limiter, exists := limiters[ip]; exists {
		return limiter
	}

	// 10 requisições por segundo
	limiter := rate.NewLimiter(rate.Every(time.Second), 10)
	limiters[ip] = limiter
	return limiter
}

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		if !getLimiter(ip).Allow() {
			respondWithError(w, http.StatusTooManyRequests, "Muitas requisições, tente novamente em breve")
			return
		}

		next.ServeHTTP(w, r)
	})
}