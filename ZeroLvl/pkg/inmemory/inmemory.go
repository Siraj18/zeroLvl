package inmemory

import (
	"errors"
	"sync"
)

type Cache struct {
	sync.Mutex
	items map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.Lock()

	defer c.Unlock()

	c.items[key] = value

}

func (c *Cache) Get(key string) interface{} {
	c.Lock()

	defer c.Unlock()

	item, found := c.items[key]

	if !found {
		return nil
	}

	return item
}

func (c *Cache) Delete(key string) error {
	c.Lock()

	defer c.Unlock()

	_, found := c.items[key]

	if !found {
		return errors.New("item not found")
	}

	delete(c.items, key)

	return nil
}
