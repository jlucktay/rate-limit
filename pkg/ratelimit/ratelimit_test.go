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
	is := is.New(t)
	v := t.Name()

	is.True(rl.Allow(v))
	is.True(rl.Allow(v))
	is.True(rl.Allow(v))
	is.True(!rl.Allow(v))
	is.True(!rl.Allow(v))
	is.True(!rl.Allow(v))
}

func TestAllowOneMinute(t *testing.T) {
	if testing.Short() {
		t.Skipf("%s takes more than one minute to run", t.Name())
	}

	t.Parallel()

	d := 1 * time.Minute
	rl := ratelimit.New(d, 3)
	is := is.New(t)
	v := t.Name()

	is.True(rl.Allow(v))
	is.True(rl.Allow(v))
	is.True(rl.Allow(v))
	is.True(!rl.Allow(v))

	time.Sleep(1 * d)

	is.True(rl.Allow(v))
	is.True(rl.Allow(v))
	is.True(rl.Allow(v))
	is.True(!rl.Allow(v))
}
