package hw04_lru_cache //nolint:golint,stylecheck
import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*listItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	actualItem := cacheItem{key, value}

	item, isFound := c.items[key]
	if isFound {
		c.queue.MoveToFront(item)
		item.Value = actualItem
	} else {
		item = c.queue.PushFront(actualItem)
		c.items[key] = item
		if c.queue.Len() > c.capacity {
			lastItem := c.queue.Back()
			c.queue.Remove(lastItem)
			delete(c.items, lastItem.Value.(cacheItem).Key)
		}
	}
	return isFound
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(value)
		return value.Value.(cacheItem).Value, true
	}
	return nil, ok
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, item := range c.items {
		c.queue.Remove(item)
	}
	c.items = make(map[Key]*listItem, c.capacity)
}

type cacheItem struct {
	Key
	Value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*listItem, capacity),
	}
}
