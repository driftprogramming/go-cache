package example

import gocache "github.com/driftprogramming/go-cache/core"

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
