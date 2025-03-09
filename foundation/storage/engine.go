package storage

import (
	"sync"
)

type InMemoryEngine struct {
	data  map[string]string
	mutex sync.RWMutex
}

func NewInMemoryEngine() *InMemoryEngine {
	return &InMemoryEngine{data: make(map[string]string)}
}

func (e *InMemoryEngine) Set(key, value string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.data[key] = value
}

func (e *InMemoryEngine) Del(key string) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	delete(e.data, key)
}

func (e *InMemoryEngine) Get(key string) (string, bool) {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	v, ok := e.data[key]

	if !ok {
		return "", false
	}

	return v, true
}
