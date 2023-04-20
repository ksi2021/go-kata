package main

import (
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Cache struct {
	data  map[string]interface{}
	init  bool
	mutex sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}, 100),
		init: true,
	}
}

func (c *Cache) Set(key string, v interface{}) error {
	if !c.init {
		return fmt.Errorf("cache isnt initialized")
	}

	c.mutex.Lock()
	c.data[key] = v
	c.mutex.Unlock()

	return nil
}

func (c *Cache) Get(key string) (interface{}, bool) {
	if !c.init {
		return nil, false
	}
	c.mutex.RLock()
	val, ok := c.data[key]
	c.mutex.RUnlock()

	return val, ok
}

func main() {
	cache := NewCache()
	keys := []string{
		"programming",
		"is",
		"so",
		"awesome",
		"write",
		"clean",
		"code",
		"use",
		"solid",
		"principles",
	}

	var eg errgroup.Group
	for i := range keys {
		idx := i
		eg.Go(func() error {
			return cache.Set(keys[idx], idx)
		})
	}

	err := eg.Wait()
	if err != nil {
		panic(err)
	}

}
