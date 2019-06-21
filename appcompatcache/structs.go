package appcompatcache

import "time"

type CacheEntry struct {
	Name      string    `json:"name"`
	Epoch     uint64    `json:"epoch"`
	Timestamp time.Time `json:"timestamp"`
}
