package store

import (
	"sync"
)

type Store interface {
	SET(string, interface{}) bool
	GET(string) any
	DEL(string) bool
}

type KeyValStore struct {
	mu  sync.Mutex
	Map map[string]interface{}
}

func (kvs *KeyValStore) SET(key string, val interface{}, time *string) bool {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()
	kvs.Map[key] = val
	return true
}

func (kvs *KeyValStore) GET(key string) interface{} {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()
	return kvs.Map[key]
}

func (kvs *KeyValStore) DEL(key string) bool {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()
	if _ , ok := kvs.Map[key]; ok {
		delete(kvs.Map, key)
		return true
    }
	return true
}

var (
	keyValStoreInstance *KeyValStore
	once                sync.Once
)

func NewKeyValueStore() *KeyValStore {
	once.Do(func() {
		keyValStoreInstance = &KeyValStore{
			Map: make(map[string]interface{}),
		}
	})
	return keyValStoreInstance
}
