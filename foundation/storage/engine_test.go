package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetGetEngine(t *testing.T) {
	t.Parallel()

	e := NewInMemoryEngine()

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
				e.Set(test.key, test.setValue)
			}

			value, exist := e.Get(test.key)
			assert.Equal(t, value, test.expectedValue)
			assert.Equal(t, exist, test.expectedIsExist)
		})
	}
}

func TestDeleteEngine(t *testing.T) {
	t.Parallel()

	e := NewInMemoryEngine()
	e.data["key"] = "value"

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
			value, exist := e.Get(test.key)
			assert.Equal(t, value, "")
			assert.Equal(t, exist, false)
		})
	}
}
