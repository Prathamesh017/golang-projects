package cache

import (
	"fmt"
	"time"
)
type Cache struct {
	cacheMap map[string] CacheEntry
}
type CacheEntry struct {
     val [] byte
	 time  time.Time
}

func NewCache() Cache{
	return Cache{
		cacheMap: make(map[string]CacheEntry),
	}
}
	
func GetCache(c Cache ,key string) ([]byte, bool){
	fmt.Println()
	cache,ok := c.cacheMap[key]
	if !ok{
		fmt.Println("Cache not found")
		return nil,false;
	}
	return cache.val, true
}

func AddCache(c *Cache, key string, val []byte) {
	c.cacheMap[key] = CacheEntry{
		val: val,
		time: time.Now(),
	}
}