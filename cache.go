package cache

import (
	"time"
)

type cacheEntry struct{
	value string
	isInfinit bool
	deadline time.Time
}

type Cache struct {
	data map[string]cacheEntry
}

func NewCache() Cache {
	return Cache{data: make(map[string]cacheEntry)}
}

func (c Cache) Get(key string) (string, bool) {
	if _, ok := c.data[key]; !ok { return "", false }
	if !c.data[key].isInfinit && time.Now().After(c.data[key].deadline) {return "", false}

	return c.data[key].value, true
}

func (c *Cache) Put(key, value string) {
	c.data[key] = struct {
		value string
		isInfinit bool
		deadline time.Time
	}{
		value: value,
		isInfinit: true,
		deadline: time.Now(),
	}
}

func (c Cache) Keys() []string {
	var res []string
	
	for k, _ := range c.data {
		res = append(res, k)
	}

	return res
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.Put(key, value)
	t := c.data[key]
	t.isInfinit = false
	t.deadline = deadline
	c.data[key] = t
}	