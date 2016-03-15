package probigo

import (
	"strings"

	"github.com/garyburd/redigo/redis"
)

// Store abstracts the redigo Redis connection.
type Store interface {
	Do(string, ...interface{}) (interface{}, error)
}

func keyExists(store Store, key string) (bool, error) {
	return redis.Bool(store.Do("EXISTS", key))
}

// isFieldExpired checks to see if the top-level key:field:ex id exists.
// Key is typically the identifier for some data structure, such as a hash
// hyperloglog, etc... If the composite top level key does not exist,
// then the field is expired.
func isFieldExpired(store Store, key string, field string) (bool, error) {
	topkey := strings.Join([]string{key, field, "ex"}, ":")
	return keyExists(store, topkey)
}
