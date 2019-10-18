package main

import (
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

// Create a map to hold the rate limiters for each visitor and a mutex.
var visitors = make(map[string]*rate.Limiter)
var mtx sync.Mutex

// Create a new rate limiter and add it to the visitors map, using the
// IP address as the key.
func addVisitor(ip string) *rate.Limiter {
	limiter := rate.NewLimiter(2, 5)
	mtx.Lock()
	visitors[ip] = limiter
	mtx.Unlock()
	return limiter
}

// Retrieve and return the rate limiter for the current visitor if it
// already exists. Otherwise call the addVisitor function to add a
// new entry to the map.
func getVisitor(ip string) *rate.Limiter {
	mtx.Lock()
	limiter, exists := visitors[ip]
	mtx.Unlock()
	if !exists {
		return addVisitor(ip)
	}
	return limiter
}

func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Call the getVisitor function to retreive the rate limiter for
		// the current user.
		limiter := getVisitor(r.RemoteAddr)
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
