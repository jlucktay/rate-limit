package ratelimit

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// Limiter is the core of the package, and provides rate-limiting functionality with its attached methods.
type Limiter struct {
	duration time.Duration
	requests uint
	visitors map[string]*visitor

	sync.Mutex
}

// visitor aids Limiter in keeping track of individual visitors, storing each of their visits.
type visitor struct {
	seen []time.Time
}

// New constructs a new Limiter with the given configuration options.
func New(d time.Duration, r uint) *Limiter {
	l := &Limiter{
		duration: d,
		requests: r,
		visitors: make(map[string]*visitor),
	}
	return l
}

// Allow will check whether or not the given visitor is allowed to make any more requests right now.
func (l *Limiter) Allow(visitor string) bool {
	log.Printf("[Allow()] visitor ID: '%s'", visitor)

	v := l.getVisitor(visitor)
	log.Printf("[Allow()] visitor: '%#v'", v)

	if uint(len(v.seen)) > l.requests {
		return false
	}

	return true
}

func (l *Limiter) getVisitor(id string) *visitor {
	l.Lock()
	visitor, exists := l.visitors[id]
	if !exists {
		l.Unlock()
		return l.addVisitor(id)
	}

	visitor.seen = append(visitor.seen, time.Now())
	l.Unlock()
	return visitor
}

func (l *Limiter) addVisitor(id string) *visitor {
	l.Lock()
	l.visitors[id] = &visitor{
		seen: []time.Time{time.Now()},
	}
	l.Unlock()

	return l.visitors[id]
}
