package ratelimit_test

import (
	"testing"
	"time"

	"github.com/jlucktay/rate-limit/pkg/ratelimit"
	"github.com/matryer/is"
)

func TestAllow(t *testing.T) {
	t.Parallel()

	rl := ratelimit.New(1*time.Minute, 3)
	i := is.New(t)
	v := "visitor_id_1"

	i.True(rl.Allow(v))
	i.True(rl.Allow(v))
	i.True(rl.Allow(v))
	i.True(!rl.Allow(v))
	i.True(!rl.Allow(v))
	i.True(!rl.Allow(v))
}
