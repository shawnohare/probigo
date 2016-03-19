package probigo

import (
	"testing"

	"github.com/shawnohare/probigo/mock"
	"github.com/stretchr/testify/assert"
)

func TestDualBloomFilterAdd(t *testing.T) {
	s := NewDualBloomFilter("test", 10, 10, mock.New())
	x := []byte("element")
	err := s.Add(x)
	assert.NoError(t, err)
}

func TestDualBloomFilterHas(t *testing.T) {
	store := mock.New()
	s := NewDualBloomFilter("test", 10000, 10, store)
	x := []byte("element")
	s.Add(x)
	// t.Logf("Store looks like: %#v", store)
	ok, err := s.Has(x)
	assert.NoError(t, err)
	assert.True(t, ok)

	ok, err = s.Has([]byte("non-member element"))
	assert.NoError(t, err)
	assert.False(t, ok)
}

// func ExampleDualBloomFilter_Has() {
// 	// Initialize a new dual bloom filter with
// 	s := NewDualBloomFilter("test", 10000, 10, NewRedis(NewRedisPool(nil)))
// 	x := []byte("example_element")
// 	s.Add(x)
// 	ok, _ := s.Has(x)
// 	fmt.Printf("Filter contains %s? %t", x, ok)
// 	// Output:
// 	// Filter contains example_element? true

// }
