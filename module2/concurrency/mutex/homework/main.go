package main

import (
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Cache struct {
	data map[string]interface{}
	init bool
	sync.Mutex
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

	c.Lock()
	c.data[key] = v
	c.Unlock()

	return nil
}

func (c *Cache) Get(key string) interface{} {
	if !c.init {
		return nil
	}
	return c.data[key]
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
