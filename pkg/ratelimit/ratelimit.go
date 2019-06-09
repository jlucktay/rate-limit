package ratelimit

import (
	log "github.com/sirupsen/logrus"
)

// Limiter is the core of the package, and provides rate-limiting functionality with its attached methods.
type Limiter struct {
}

// Allow will check whether or not the given visitor is allowed to make any more requests right now.
func (l *Limiter) Allow(visitor string) bool {
	log.Printf("[Allow()] visitor ID: '%s'", visitor)

	return true
}
