package ratelimit

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const slidingWindowFactor = 60

// Limiter is the core of the package, and provides rate-limiting functionality with its attached methods.
type Limiter struct {
	duration time.Duration
	requests uint
	visitors map[string]*visitor

	sync.Mutex
}

type window int64

// visitor aids Limiter in keeping track of individual visitors, storing their visits in a map where the keys are
// sliding windows and the values are the number of visits within that window.
type visitor struct {
	seen map[window]uint

	sync.Mutex
}

// New constructs a new Limiter with the given configuration options.
func New(d time.Duration, r uint) *Limiter {
	l := &Limiter{
		duration: d,
		requests: r,
		visitors: make(map[string]*visitor),
	}

	go l.cleanupVisitors()

	return l
}

// Allow will check whether or not the given visitor is allowed to make any more requests right now.
func (l *Limiter) Allow(visitor string) bool {
	log.Printf("[Allow()] visitor ID: '%s'", visitor)

	v := l.getVisitor(visitor)
	log.Printf("[Allow()] visitor: '%#v'", v)

	if v.countVisits() > l.requests {
		// This visitor has hit their limit.
		return false
	}

	return true
}

// getVisitor checks for pre-existing visitors with the same ID, and either returns them directly, or calls a create
// method accordingly. For visitors that have been here before, their visit count is incremented appropriately.
func (l *Limiter) getVisitor(id string) *visitor {
	l.Lock()

	visitor, exists := l.visitors[id]
	if !exists {
		l.Unlock()
		return l.addVisitor(id)
	}

	l.visited(visitor)
	l.Unlock()

	return visitor
}

// addVisitor will create a new visitor with one visit logged against a time windows based on right now.
func (l *Limiter) addVisitor(id string) *visitor {
	l.Lock()
	defer l.Unlock()

	l.visitors[id] = &visitor{
		seen: map[window]uint{
			timeWindow(time.Now()): 1,
		},
	}

	return l.visitors[id]
}

// visited will add the new visit, if the visitor hasn't already hit their limit. Here, the map is modified without
// explicit lock/unlock calls, as the only calling function getVisitor has already done so.
func (l *Limiter) visited(v *visitor) {
	// Add this visit to the appropriate time window.
	if v.countVisits() <= l.requests {
		w := timeWindow(time.Now())
		v.seen[w]++
	}
}

// prune removes old visits from before the Limiter's duration window.
func (l *Limiter) prune(v *visitor) {
	// Capture current time window.
	w := timeWindow(time.Now())

	// Express Limiter's duration in same terms.
	d := durationWindow(l.duration)

	v.Lock()
	defer v.Unlock()

	for visitWindow := range v.seen {
		// Roll off old time windows that are past the oldest one that we are concerned with.
		if w-visitWindow > d {
			delete(v.seen, visitWindow)
		}
	}
}

// cleanupVisitors starts running on its own goroutine when a new Limiter is constructed, and prunes old visit windows
// at regular intervals, as well as the visitors themselves when they no longer have any visits within the duration
// that the Limiter is keeping track of.
func (l *Limiter) cleanupVisitors() {
	cleanupInterval := l.duration.Nanoseconds() / slidingWindowFactor

	for {
		time.Sleep(time.Duration(cleanupInterval) * time.Nanosecond)

		l.Lock()
		for id, v := range l.visitors {
			l.prune(v)

			if v.countVisits() == 0 {
				delete(l.visitors, id)
			}
		}
		l.Unlock()
	}
}

// countVisits tallies up all tracked visits for the given visitor.
func (v *visitor) countVisits() (total uint) {
	v.Lock()
	defer v.Unlock()

	for _, visitCount := range v.seen {
		total += visitCount
	}
	return
}

// timeWindow divides timestamps by a constant factor to give us keys for tracking our sliding windows.
func timeWindow(t time.Time) window {
	return window(t.UnixNano() / 1e9 / slidingWindowFactor)
}

// durationWindows converts durations in the same manner as timeWindow() for instances where the two need to be
// compared against each other.
func durationWindow(d time.Duration) window {
	return window(d.Nanoseconds() / 1e9 / slidingWindowFactor)
}
