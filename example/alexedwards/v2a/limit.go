package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"golang.org/x/time/rate"
)

// Create a new rate limiter and add it to the visitors map, using the IP address as the key.
func addVisitor(ip string, rem *remember) *rate.Limiter {
	limiter := rate.NewLimiter(2, 5)
	rem.Lock()
	rem.visitors[ip] = limiter
	rem.Unlock()
	return limiter
}

// Retrieve and return the rate limiter for the current visitor if it already exists. Otherwise call the addVisitor
// function to add a new entry to the map.
func getVisitor(ip string, rem *remember) *rate.Limiter {
	rem.Lock()
	limiter, exists := rem.visitors[ip]
	rem.Unlock()
	if !exists {
		return addVisitor(ip, rem)
	}
	return limiter
}

type remember struct {
	visitors map[string]*rate.Limiter
	sync.Mutex
}

func limit(next http.Handler) http.Handler {
	// Create a map to hold the rate limiters for each visitor and a mutex.
	rem := &remember{
		visitors: make(map[string]*rate.Limiter),
	}
	// var visitors = make(map[string]*rate.Limiter)
	// var mtx sync.Mutex

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Weeding out requests from localhost on cycling ports.

		// index := strings.LastIndex(r.RemoteAddr, ":")
		// split := strings.Split(r.RemoteAddr, ":")
		// fmt.Fprintf(os.Stderr, "r.RemoteAddr/index/sliced/split: '%s'/%d/'%#v'/'%#v'\n",
		// 	r.RemoteAddr,
		// 	index,
		// 	r.RemoteAddr[:index],
		// 	split,
		// )

		visitorIP := r.RemoteAddr[:strings.LastIndex(r.RemoteAddr, ":")]
		fmt.Fprintf(os.Stderr, "visitorIP: '%s'\n", visitorIP)
		// Finished weeding.

		// Call the getVisitor function to retrieve the rate limiter for the current user.
		limiter := getVisitor(visitorIP, rem)

		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
