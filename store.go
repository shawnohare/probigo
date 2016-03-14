package probigo

// Define some data structure types typically supported by Store interfaces.
const (
	// Top level element
	Top = iota + 1
	Hash
	Set
	OrderedSet
	HyperLogLog
)

// Data is an object that a Data Store manipulates.
type Data struct {
	// Element is the underlying object pushed to a probabilistic data structure.
	Element []byte
	// Key associated to the data.  This is usually computed by a probabilistic
	// data structure, and not set by the user.
	Key string
	// Expiry indicates when, in seconds, the element should expire. (Optional)
	Expiry int
	// Type represents what type of data structure the store should use to
	// store this element.
	Type int
}

// Store implementations represent the minimal functionality that a data
// structure store should provide to back the probabilistic
// data structures defined in this package.  For simplicity, all set
// elements are assumed to be of type []byte.
//
// Generally speaking, the user is not expected to provide their own
// implementation.  By default, the probabilistic data structures are
// Redis backed.  User defined backings can be useful for unit testing.
type Store interface {
	Exists(*Data) (bool, error)
	Get(*Data) (*Data, error)
	Set(*Data) error
	Del(*Data) error
}

// type Store interface {
// 	// Get the value for the input key.
// 	Get(string) ([]byte, error)
// 	// Set the given key to the value, with an optional expiration in seconds.
// 	Set(string, []byte, int) error
// 	// Get the value associated to the key and set it to the input value
// 	// with an optional expiration in seconds.
// 	GetSet(string, []byte, int) ([]byte, error)
// 	Delete(string) error
// 	Exists(string) (bool, error)
// }
