package cache

import (
	"sync"

	"github.com/Dislate/test-task-l0/models"
)

type Cache struct {
	sync.RWMutex
	items map[int]interface{}
}

func New() *Cache {
	items := make(map[int]interface{})
	cache := Cache{items: items}

	return &cache
}

func (c *Cache) Set(key int, value models.Order) {
	c.Lock()
	defer c.Unlock()

	c.items[key] = value
}

func (c *Cache) Get(key int) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, found := c.items[key]

	if !found {
		return nil, false
	}

	return item, true
}
