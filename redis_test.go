package probigo

// import (
// 	"github.com/stretchr/testify/assert"
// 	"testing"
// )

// func TestRedisSet(t *testing.T) {
// 	if testing.Short() {
// 		t.SkipNow()
// 	}

// 	r := NewRedis(NewRedisPool(nil))
// 	key := "test"
// 	x := []byte("element")
// 	err := r.Set(key, x, 0)
// 	assert.NoError(t, err)
// 	err = r.Set(key, x, 10)
// 	assert.NoError(t, err)
// }

// func TestRedisGet(t *testing.T) {
// 	if testing.Short() {
// 		t.SkipNow()
// 	}

// 	r := NewRedis(NewRedisPool(nil))
// 	key := "test"
// 	x := []byte("element")
// 	r.Set(key, x, 10)
// 	v, err := r.Get(key)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, v)
// 	assert.Equal(t, x, v)
// }

// func TestRedisDelete(t *testing.T) {
// 	if testing.Short() {
// 		t.SkipNow()
// 	}

// 	r := NewRedis(NewRedisPool(nil))
// 	key := "test"
// 	err := r.Delete(key)
// 	assert.NoError(t, err)
// }

// func TestRedisSetThenDelete(t *testing.T) {
// 	if testing.Short() {
// 		t.SkipNow()
// 	}

// 	r := NewRedis(NewRedisPool(nil))
// 	key := "test"
// 	x := []byte("element")

// 	r.Set(key, x, 10)
// 	v, _ := r.Get(key)
// 	assert.NotEmpty(t, v)
// 	r.Delete(key)
// 	v2, _ := r.Get(key)
// 	assert.Empty(t, v2)
// }
