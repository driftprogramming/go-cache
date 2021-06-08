# Usage

## What

This repo is a wrapper with [go-cache](https://github.com/patrickmn/go-cache). We use `gocache.GetOrSet` to set or get
cache thread safely. It means that when multiple threads/go-routing call `gocache.GetOrSet`, the cache will be set only
once. Usually it means we just call database to retrieve data ONLY one time. This is very useful when the first time
startup the application or when cache expired. It reduces the pressure on database.

#### Install

```
go get -u github.com/driftprogramming/go-cache
```

#### STEP 1: Register all your cache keys

```go
package example

import gocache "go-cache/core"

const (
	KeyOrders  = "KEY_ORDERS"
	KeyUsers   = "KEY_USERS"
	KeyBooks   = "KEY_BOOKS"
	KeyModules = "KEY_MODULES"
	KeyConfig  = "KEY_CONFIG"
)

func RegisterAllCacheKeys() {
	registry := gocache.GetRegistryInstance()
	registry.Register(&gocache.CacheKey{Key: KeyOrders, Expire: gocache.OneDay})
	registry.Register(&gocache.CacheKey{Key: KeyUsers, Expire: gocache.OneMinute})
	registry.Register(&gocache.CacheKey{Key: KeyBooks, Expire: gocache.OneHour})
	registry.Register(&gocache.CacheKey{Key: KeyModules, Expire: gocache.FiveMinutes})
	registry.Register(&gocache.CacheKey{Key: KeyConfig, Expire: gocache.OneMinute})
}

```

#### STEP 2: Use `gocache.GetOrSet` to get or set cache thread safely.

```go
package example

import (
	gocache "go-cache/core"

	log "github.com/sirupsen/logrus"
)

type Book struct {
	name string
}

func StartApp() {
	RegisterAllCacheKeys()
	books, ok := gocache.GetOrSet(KeyBooks, func(cache *gocache.CacheKey) (interface{}, error) {
		// you also can call database to retrieve books here.
		return []Book{
			{name: "i love coding"},
			{name: "coding is great"},
		}, nil
	})

	if ok {
		log.Info("Got books and also set cache.")
		log.Info(books)
	} else {
		log.Info("No books found in cache also in database maybe")
	}
}
```

example refer to `./example`