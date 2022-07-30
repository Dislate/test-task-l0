package cache

import (
	"sync"

	"github.com/Dislate/test-task-l0/models"
)

type Cache struct {
	sync.RWMutex
	items map[string]interface{}
}

func New() *Cache {
	items := make(map[string]interface{})
	cache := Cache{items: items}

	return &cache
}

func (c *Cache) Set(value models.Order) {
	c.Lock()
	defer c.Unlock()

	uid := value.Order_uid
	c.items[uid] = value
}

func (c *Cache) Get(uid string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, found := c.items[uid]

	if !found {
		return nil, false
	}

	return item, true
}
