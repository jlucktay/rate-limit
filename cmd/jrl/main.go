package main

import (
	"log"
	"net/http"

	"github.com/jlucktay/rate-limit/pkg/ratelimit"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	log.Fatal(http.ListenAndServe(":8080", limit(mux)))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // 200
	w.Header().Set("Content-Type", "text/plain")

	if _, errWrite := w.Write([]byte("OK")); errWrite != nil {
		log.Fatal(errWrite)
	}
}

func limit(next http.Handler) http.Handler {
	limiter := &ratelimit.Limiter{}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
