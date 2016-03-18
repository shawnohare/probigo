package mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	s := New()
	_, err := s.set("key", 1)
	assert.NoError(t, err)

	// Fails
	_, err = s.set()
	assert.Error(t, err)
	_, err = s.set("novalue")
	assert.Error(t, err)
	_, err = s.set(1, nil)
	assert.Error(t, err)
}

func TestGet(t *testing.T) {
	s := New()
	s.set("key", 1)
	val, err := s.get("key")
	assert.NoError(t, err)
	assert.Equal(t, int(1), val.(int))

	// Fails
	_, err = s.get()
	assert.Error(t, err)
	_, err = s.get(1)
	assert.Error(t, err)
}

func TestDoSet(t *testing.T) {
	s := New()
	_, err := s.Do("SET", "key", 1)
	assert.NoError(t, err)
}

func TestFset(t *testing.T) {
	s := New()
	_, err := s.fset("key", "field", 1)
	assert.NoError(t, err)

	// Fail arguments
	fails := [][]interface{}{
		{},
		{1},
		{"key"},
		{"key", 1},
	}

	for _, test := range fails {
		_, err := s.fget(test)
		assert.Error(t, err)
	}

}

func TestFget(t *testing.T) {
	s := New()
	s.fset("key", "field", 1)
	val, err := s.fget("key", "field")
	assert.NoError(t, err)
	assert.Equal(t, int(1), val.(int))

	// Fails
	fails := [][]interface{}{
		{},
		{1},
		{"key"},
		{"key", 1},
	}

	for _, test := range fails {
		_, err := s.fget(test)
		assert.Error(t, err)
	}
}
