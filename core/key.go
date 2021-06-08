package gocache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

type CacheKey struct {
	Key    string
	Sync   sync.Mutex
	Expire time.Duration
}

var (
	Default = &CacheKey{Key: "__DEFAULT_KEY__", Expire: cache.DefaultExpiration}
)
