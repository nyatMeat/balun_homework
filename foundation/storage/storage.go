package storage

import (
	"balun_homework_1/foundation/logger"
	"context"
)

type InMemoryStorage struct {
	e   *InMemoryEngine
	log *logger.Logger
}

func NewInMemoryStorage(e *InMemoryEngine, log *logger.Logger) *InMemoryStorage {
	return &InMemoryStorage{
		e:   e,
		log: log,
	}
}

func (ms *InMemoryStorage) Get(ctx context.Context, key string) (string, bool, error) {
	ms.log.Debug(ctx, "[InMemoryStorage::Get] Try to get value", key)

	v, ok := ms.e.Get(key)

	if !ok {
		ms.log.Info(ctx, "[InMemoryStorage::Get] Value not found", key)
	}

	return v, ok, nil
}

func (ms *InMemoryStorage) Set(ctx context.Context, key string, value string) (bool, error) {
	ms.log.Debug(ctx, "[InMemoryStorage::Set] Try to set value", key)

	ms.e.Set(key, value)

	ms.log.Debug(ctx, "[InMemoryStorage::Set] Value set to the key", key)

	return true, nil
}

func (ms *InMemoryStorage) Delete(ctx context.Context, key string) (bool, error) {
	ms.log.Debug(ctx, "[InMemoryStorage::Set] Try to set value", key)

	ms.e.Del(key)

	return true, nil
}
