package gocache_test

import (
	"fmt"
	"sync"
	"testing"

	gocache "github.com/driftprogramming/go-cache/core"
	log "github.com/sirupsen/logrus"
)

var count = 0

func TestInstance_UsageExample(t *testing.T) {
	t.Parallel()
	registry := gocache.GetRegistryInstance()
	instance := gocache.GetInstance()
	registry.Register(&gocache.CacheKey{Key: "xyz", Expire: gocache.OneMinute})
	v := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		v.Add(1)
		go func() {
			job(instance)
			v.Done()
		}()
	}

	v.Wait()
	log.Info(fmt.Sprintf("count = %v", count))
	if count != 1 {
		log.Error("The CallbackIfCacheMissing should ba called ONLY once no matter how go-routing we have.")
	}
}

func job(instance *gocache.CacheInstance) {
	result, _ := instance.GetOrSet("xyz", func(cache *gocache.CacheKey) (interface{}, error) {
		count++
		log.Info("Got cache key : " + cache.Key)
		log.Info(fmt.Sprintf("Got cache expire : %v", cache.Expire))
		return "999", nil
	})

	log.Info(fmt.Sprintf("Go result : %v", result))
}
