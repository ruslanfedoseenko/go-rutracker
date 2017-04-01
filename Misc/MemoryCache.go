package Misc

import (
	"reflect"
	"fmt"
	"errors"
	"sync"
)

type MemoryCache struct {
	storage map[string]interface{}
	accessGuard sync.Mutex
}


func NewMemoryCache() *MemoryCache{
	return &MemoryCache{
		storage: make(map[string]interface{}),
	}
}

func (c *MemoryCache) Get(key string, v interface{}) (error) {
	c.accessGuard.Lock()
	defer c.accessGuard.Unlock()
	if val, ok := c.storage[key]; ok {
		inputType := reflect.ValueOf(v).Type()
		cacheValueType := reflect.ValueOf(val).Type()
		if inputType == cacheValueType {
			v = val
			return nil
		} else {
			return errors.New(fmt.Sprintf("Unable to assign %s to %s", cacheValueType, inputType))
		}
	} else {
		return errors.New(fmt.Sprintf("Specified Key Not found: %s", key))
	}
}


func (c *MemoryCache) Put(key string, v interface{}) {
	c.accessGuard.Lock()
	defer c.accessGuard.Unlock()
	c.storage[key] = v
}