package example

import (
	gocache "github.com/driftprogramming/go-cache/core"
	log "github.com/sirupsen/logrus"
)

func StartApp() {
	RegisterAllCacheKeys()
	books, ok := gocache.GetOrSet(KeyBooks, func(cache *gocache.CacheKey) (interface{}, error) {
		// you also can call database to retrieve cache here.
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
