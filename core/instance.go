package gocache

import (
	"sync"

	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

type CacheInstance struct {
	*cache.Cache
}

var cacheInstance *CacheInstance
var onceCreateCacheInstance sync.Once

func GetInstance() *CacheInstance {
	onceCreateCacheInstance.Do(func() {
		cacheInstance = &CacheInstance{cache.New(defaultCacheExpiration, defaultCacheCleanupInterval)}
	})

	return cacheInstance
}

type CallbackIfCacheMissing func(cache *CacheKey) (interface{}, error)

func (instance *CacheInstance) GetOrSet(key string, callback CallbackIfCacheMissing) (interface{}, bool) {
	keyRegistry := GetRegistryInstance()
	cache := keyRegistry.Get(key)

	if cache == nil {
		log.Panicf("Cache key %v not found in registry, you need to Register it first", key)
		return nil, false
	}

	_, ok := instance.Get(cache.Key)
	if !ok { // if cache missing, set cache thread safely
		cache.Sync.Lock()
		defer cache.Sync.Unlock()
		_, okAgain := instance.Get(cache.Key)
		if !okAgain {
			log.Infof("Cache missing for %s, setting cache with callback ...", cache.Key)
			value, err := callback(cache)
			if err == nil {
				log.Infof("Set cache for %s successfully", cache.Key)
				instance.Set(cache.Key, value, cache.Expire)
			} else {
				log.Errorf("Set cache for %s failed: %v", cache.Key, err)
			}
		}
	}

	return instance.Get(cache.Key)
}
