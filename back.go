package probigo

// Back implementations represent the minimal functionality that a key-value
// store (or data structure server) should provide to back the probabilistic
// data structures defined in this package.  For simplicity, all set
// elements are assumed to be of type []byte.
//
// Generally speaking, the user is not expected to provide their own
// implementation.  By default, the probabilistic data structures are
// Redis backed.  User defined backings can be useful for unit testing.
type Back interface {
	// Get the value for the input key.
	Get(string) ([]byte, error)
	// Set the given key to the value, with an optional expiration in seconds.
	Set(string, []byte, ...int) error
	// Get the value associated to the key and set it to the input value
	// with an optional expiration in seconds.
	GetSet(string, []byte, ...int) ([]byte, error)
	Exists(string) (bool, error)
}
