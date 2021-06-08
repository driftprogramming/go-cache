package gocache

import "time"

const (
	OneDay                      = 24 * time.Hour
	OneHour                     = 1 * time.Hour
	OneMinute                   = 1 * time.Minute
	FiveMinutes                 = 5 * time.Minute
	defaultCacheExpiration      = OneDay
	defaultCacheCleanupInterval = OneHour
)
