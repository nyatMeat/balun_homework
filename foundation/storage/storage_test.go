package storage

import (
	"balun_homework_1/foundation/logger"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetGetStorage(t *testing.T) {
	t.Parallel()

	e := NewInMemoryEngine()
	e.data["key"] = "value"
	s := NewInMemoryStorage(e, logger.CreateMock())

	ctx := context.Background()

	tests := map[string]struct {
		key             string
		setValue        string
		expectedValue   string
		expectedIsExist bool
	}{
		"get existing key": {
			key:             "key1",
			setValue:        "value",
			expectedValue:   "value",
			expectedIsExist: true,
		},
		"get non existing key": {
			key:             "key2",
			setValue:        "",
			expectedValue:   "",
			expectedIsExist: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if len(test.setValue) > 0 {
				ok, err := s.Set(ctx, test.key, test.setValue)

				assert.True(t, ok)
				assert.Nil(t, err)
			}

			value, exist, err := s.Get(ctx, test.key)
			assert.Equal(t, value, test.expectedValue)
			assert.Equal(t, exist, test.expectedIsExist)
			assert.Nil(t, err)
		})
	}
}

func TestDeleteStorage(t *testing.T) {
	t.Parallel()

	e := NewInMemoryEngine()
	e.data["key"] = "value"
	s := NewInMemoryStorage(e, logger.CreateMock())

	ctx := context.Background()

	tests := map[string]struct {
		key   string
		value string
	}{
		"delete existing key": {
			key: "key",
		},
		"delete non existing key": {
			key: "key1",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			e.Del(test.key)
			value, exist, err := s.Get(ctx, test.key)
			assert.Equal(t, value, "")
			assert.False(t, exist)
			assert.Nil(t, err)
		})
	}
}
