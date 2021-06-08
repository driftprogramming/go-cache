package gocache

import (
	"sync"
)

type Registry interface {
	Register(key *CacheKey)
	Get(key string) *CacheKey
}

type RegistryImpl struct {
	keymap sync.Map
}

func (r *RegistryImpl) Get(key string) *CacheKey {
	value, ok := r.keymap.Load(key)
	if ok {
		return value.(*CacheKey)
	} else {
		return nil
	}
}

func (r *RegistryImpl) Register(cacheKey *CacheKey) {
	r.keymap.Store(cacheKey.Key, cacheKey)
}

var onceCreateRegistry sync.Once
var registryInstance Registry

func GetRegistryInstance() Registry {
	onceCreateRegistry.Do(func() {
		registryInstance = &RegistryImpl{}
	})

	return registryInstance
}
