package storage

import (
	"balun_homework_1/foundation/logger"
	"context"
)

type MapStorage struct {
	e   *MapEngine
	log *logger.Logger
}

func NewMapStorage(e *MapEngine, log *logger.Logger) *MapStorage {
	return &MapStorage{
		e:   e,
		log: log,
	}
}

func (ms *MapStorage) Get(ctx context.Context, key string) (string, bool, error) {
	ms.log.Debug(ctx, "[MapStorage::Get] Try to get value", key)

	v, ok := ms.e.Get(key)

	if !ok {
		ms.log.Info(ctx, "[MapStorage::Get] Value not found", key)
	}

	return v, ok, nil
}

func (ms *MapStorage) Set(ctx context.Context, key string, value string) (bool, error) {
	ms.log.Debug(ctx, "[MapStorage::Set] Try to set value", key)

	ms.e.Set(key, value)

	ms.log.Debug(ctx, "[MapStorage::Set] Value set to the key", key)

	return true, nil
}

func (ms *MapStorage) Delete(ctx context.Context, key string) (bool, error) {
	ms.log.Debug(ctx, "[MapStorage::Set] Try to set value", key)

	ms.e.Del(key)

	return true, nil
}
