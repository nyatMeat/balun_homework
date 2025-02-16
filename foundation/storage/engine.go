package storage

import (
	"fmt"
)

type MapEngine struct {
	data map[string]string
}

func NewMapEngine() *MapEngine {
	return &MapEngine{data: make(map[string]string)}
}

func (e *MapEngine) Set(key, value string) {
	e.data[key] = value
}

func (e *MapEngine) Del(key string) {
	delete(e.data, key)
}

func (e *MapEngine) Get(key string) (string, bool) {
	v, ok := e.data[key]

	if !ok {
		return "", false
	}

	return fmt.Sprintf("%v", v), true
}
