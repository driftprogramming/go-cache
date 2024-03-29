package example

import (
	gocache "github.com/driftprogramming/go-cache/core"
	log "github.com/sirupsen/logrus"
)

func StartApp() {
	RegisterAllCacheKeys()
	cacheInstance := gocache.GetInstance()
	books, ok := cacheInstance.GetOrSet(KeyBooks, func(cache *gocache.CacheKey) (interface{}, error) {
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
