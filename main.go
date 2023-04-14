package main

import (
	"sync"
)

type Cache struct {
	Data map[string]interface{}
	sync.Mutex
}

func (c *Cache) Set(key string, data interface{}) {
	c.Lock()
	c.Data[key] = data
	c.Unlock()
}

func (c *Cache) Get(key string) interface{} {
	c.Lock()
	temp := c.Data[key]
	c.Unlock()
	return temp
}

func MakeCache() *Cache {
	return &Cache{Data: make(map[string]interface{})}
}

func main() {
	// cashe := MakeCache()

	// cashe.Set("test", 123)

	// fmt.Println(cashe.Get("test1"))
}
