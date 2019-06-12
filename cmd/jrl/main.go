package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jlucktay/rate-limit/pkg/ratelimit"
	log "github.com/sirupsen/logrus"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okHandler)

	log.Fatal(http.ListenAndServe(":8080", limit(mux)))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK) // 200
	w.Header().Set("Content-Type", "text/plain")

	if _, errWrite := w.Write([]byte("OK\n")); errWrite != nil {
		log.Fatal(errWrite)
	}
}

func limit(next http.Handler) http.Handler {
	log.Print("setting up Limiter...")
	limiter := ratelimit.New(1*time.Minute, 3, log.TraceLevel)
	log.Print("Limiter setup complete!")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("JRL-ID")
		if id == "" {
			http.Error(
				w,
				fmt.Sprintf(
					"The 'JRL-ID' header was not specified on the request: %s",
					http.StatusText(http.StatusBadRequest),
				),
				http.StatusBadRequest,
			)
			return
		}

		if !limiter.Allow(id) {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
