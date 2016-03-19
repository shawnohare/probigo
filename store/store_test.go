package store

// import (
// 	"testing"

// 	"github.com/shawnohare/probigo/mock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestRedigoMock(t *testing.T) {
// 	store := redigomock.NewConn()
// 	_, err := store.Do("SET", "mykey", "myval")
// 	assert.NoError(t, err)
// 	gotVal, err := redis.String(store.Do("GET", "mykey"))
// 	assert.NoError(t, err)
// 	assert.Equal(t, "mykey", gotVal)
// }

// func newMockStore() *mock.Store {
// 	store := mock.New()
// 	store.Do("")

// }

// func TestSet(t *testing.T) {
// 	s := mock.New()
// 	_, err := set(s, "key", 1)
// 	assert.NoError(t, err)
// }

// func TestSetThenGet(t *testing.T) {
// 	s := mock.New()
// 	set(s, "key", 1)
// 	v, err := get(s, "key")
// 	assert.NoError(t, err)
// 	assert.Equal(t, 1, v.(int))
// }

// func TestGet(t *testing.T) {
// 	s := mock.New()
// 	s.Init() // load in dummy data
// 	vi, err := get(s, "key1Int")
// 	assert.NoError(t, err)
// 	assert.Equal(t, 1, vi.(int))
// }

// func TestGetField(t *testing.T) {
// 	s := mock.New()
// 	s.Init()
// 	v, err := getField(s, redisHash, "hash1", "key1Int")
// 	assert.NoError(t, err)
// 	assert.Equal(t, 1, v.(int))
// }
