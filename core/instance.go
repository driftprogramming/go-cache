package gocache

import (
	"sync"

	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

var cacheInstance *cache.Cache
var onceCreateCacheInstance sync.Once

func GetInstance() *cache.Cache {
	onceCreateCacheInstance.Do(func() {
		cacheInstance = cache.New(defaultCacheExpiration, defaultCacheCleanupInterval)
	})

	return cacheInstance
}

type CallbackIfCacheMissing func(cache *CacheKey) (interface{}, error)

func GetOrSet(key string, callback CallbackIfCacheMissing) (interface{}, bool) {
	keyRegistry := GetRegistryInstance()
	cacheInstance := GetInstance()
	cache := keyRegistry.Get(key)
	if cache == nil {
		log.Panicf("Cache key %v not found in registry, you need to Register it first", key)
		return nil, false
	}

	_, ok := cacheInstance.Get(cache.Key)
	if !ok { // if cache missing, set cache thread safely
		cache.Sync.Lock()
		defer cache.Sync.Unlock()
		_, okAgain := cacheInstance.Get(cache.Key)
		if !okAgain {
			log.Infof("Cache missing for %s, setting cache with callback ...", cache.Key)
			value, err := callback(cache)
			if err == nil {
				log.Infof("Set cache for %s successfully", cache.Key)
				cacheInstance.Set(cache.Key, value, cache.Expire)
			} else {
				log.Errorf("Set cache for %s failed: %v", cache.Key, err)
			}
		}
	}

	return cacheInstance.Get(cache.Key)
}
