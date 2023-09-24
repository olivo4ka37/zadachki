package CacheStructure

import (
	"errors"
	"fmt"
	"sync"
)

func main() {
	var x1, x2, x3 Order

	x1.Uid = "X1"
	x1.Info = "IT IS X1"
	x2.Uid = "X2"
	x2.Info = "IT IS X2"
	x3.Uid = "X3"
	x3.Info = "IT IS X3"
	Cache := NewCache()
	fmt.Println(Cache)
	Cache.Set(&x1)
	Cache.Set(&x2)
	Cache.Set(&x3)
	fmt.Println(Cache.Cache["X1"])
}

type Order struct {
	Uid  string
	Info string
}

type MemoryCache struct {
	sync.RWMutex
	Cache map[string]*Order
}

func NewCache() *MemoryCache {
	items := make(map[string]*Order)
	cache := &MemoryCache{
		Cache: items,
	}
	return cache
}

func (c *MemoryCache) Set(value *Order) {

	c.Lock()
	defer c.Unlock()

	c.Cache[value.Uid] = value
}

func (c *MemoryCache) Get(key string) (*Order, error) {

	c.RLock()

	defer c.RUnlock()

	item, found := c.Cache[key]

	// ключ не найден
	if !found {
		return nil, errors.New("Cant get element from MemoryCache")
	}

	return item, nil
}
