package ratelimit

type Limiter struct {
}

func (l *Limiter) Allow() bool {
	return true
}
