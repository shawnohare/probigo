// Package store provides a structure that exposes typed versions of
// a redigo connection's Do method.
package store

import (
	"strings"

	"github.com/garyburd/redigo/redis"
)

// Interface
// type Interface interface {
// 	// Do(string, ...interface{}) (interface{}, error)
// 	// Get(string) (interface{}, error)
// 	// Set(string, interface{}, int) error
// 	// Delete(string) error
// 	// GetBytes(string) ([]byte, error)
// 	// SetBytes(string, []byte, int) error
// 	// HGet(string, string) (interface{}, error)
// 	// HGetBytes(string, string) ([]byte, error)
// 	// HSet(string, string, interface{}) error
// 	// HDelete(string, string) error
// }

// Store represents a data structure store.  It merely wraps the Doer
// interface it is initialized with to provide some additional type-safe
// versions of the Do method.
type Store struct {
	conn Conn
}

func (s *Store) Do(cmd string, args ...interface{}) (interface{}, error) {
	return s.conn.Do(cmd, args...)
}

func (s *Store) Get(key string) (interface{}, bool, error) {
	return get(s.conn, key)
}

func (s *Store) Set(key string, value interface{}) error {
	_, err := set(s.conn, key, value)
	return err
}

func (s *Store) SetBytes(key string, value []byte) error {
	return s.Set(key, value)
}

func (s *Store) GetBytes(key string) ([]byte, bool, error) {
	return getBytes(s.conn, key)
}

func (s *Store) Exists(key string) (bool, error) {
	return exists(s.conn, key)
}

func (s *Store) HSet(key string, field string, value interface{}) error {
	_, err := setField(s.conn, "HSET", key, field, value)
	return err
}

func (s *Store) HSetBytes(key string, field string, value []byte) error {
	return s.HSet(key, field, value)
}

func (s *Store) HGet(key string, field string) (interface{}, bool, error) {
	return getField(s.conn, "HGET", key, field)
}

func (s *Store) HGetBytes(key string, field string) ([]byte, bool, error) {
	return getFieldBytes(s.conn, "HGET", key, field)
}

func (s *Store) HExists(key string, field string) (bool, error) {
	// TODO
	return false, nil
}

func exists(d Conn, key string) (bool, error) {
	return redis.Bool(d.Do("EXISTS", key))
}

// get a top level key from the Store.
func get(d Conn, key string) (interface{}, bool, error) {
	v, err := d.Do("GET", key)
	if err != nil {
		return nil, false, err
	}
	return v, v == nil, nil
}

func getBytes(d Conn, key string) ([]byte, bool, error) {
	v, ok, err := get(d, key)
	if err != nil || !ok {
		return nil, ok, err
	}
	// Otherwse, v != nil and there is no error.
	bs, err := redis.Bytes(v, err)
	return bs, ok, err
}

// set a top level key .
func set(d Conn, key string, value interface{}) (interface{}, error) {
	return d.Do("SET", key, value)
}

// isFieldExpired checks to see if the top-level key:field:ex id exists.
// Key is typically the identifier for some data structure, such as a hash
// hyperloglog, etc... If the composite top level key does not exist,
// then the field is expired.
func isFieldExpired(conn Conn, key string, field string) (bool, error) {
	topkey := strings.Join([]string{key, field, "ex"}, ":")
	return exists(conn, topkey)
}

// fieldExists is responsible for handling
func fieldExists(conn Conn, cmd string, key string, field string) (bool, error) {
	return redis.Bool(conn.Do(cmd, key, field))
}

// getField get's a field value.
func getField(conn Conn, cmd string, key string, field string) (interface{}, bool, error) {
	v, err := conn.Do(cmd, key, field)
	if err != nil {
		return nil, false, err
	}
	return v, v == nil, nil
}

// getFieldBytes cast the field value as a byte slice.
func getFieldBytes(conn Conn, cmd string, key string, field string) ([]byte, bool, error) {
	v, ok, err := getField(conn, cmd, key, field)
	if err != nil || !ok {
		return nil, ok, err
	}
	bs, err := redis.Bytes(v, err)
	return bs, ok, err
}

// getField get's a field value from a Store struct.
func setField(conn Conn, cmd string, key string, field string, value interface{}) (interface{}, error) {
	return conn.Do(cmd, key, field, value)
}

// New promotes a Conn instance to a Store.
func New(connection Conn) *Store {
	return &Store{conn: connection}
}
