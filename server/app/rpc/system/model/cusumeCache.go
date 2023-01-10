package model

import (
	"github.com/zeromicro/go-zero/core/mathx"
	"time"
)

func init() {
	cacheOption = newOptions()
	unstable = mathx.NewUnstable(expiryDeviation)
}

var (
	cacheOption           Options
	unstable              mathx.Unstable
	defaultExpiry         = time.Hour * 24 * 7
	defaultNotFoundExpiry = time.Minute
	// make the expiry unstable to avoid lots of cached items expire at the same time
	// make the unstable expiry to be [0.95, 1.05] * seconds
	expiryDeviation = 0.05
)

type (
	Options struct {
		Expiry         time.Duration
		NotFoundExpiry time.Duration
	}

	Option func(o *Options)
)

func newOptions(opts ...Option) Options {
	var o Options
	for _, opt := range opts {
		opt(&o)
	}

	if o.Expiry <= 0 {
		o.Expiry = defaultExpiry
	}
	if o.NotFoundExpiry <= 0 {
		o.NotFoundExpiry = defaultNotFoundExpiry
	}

	return o
}
func WithExpiry(expiry time.Duration) Option {
	return func(o *Options) {
		o.Expiry = expiry
	}
}

func WithNotFoundExpiry(expiry time.Duration) Option {
	return func(o *Options) {
		o.NotFoundExpiry = expiry
	}
}
