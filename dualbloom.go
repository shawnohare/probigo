package probigo

import (
	"hash"
	"hash/fnv"
	"log"
	"strconv"
	"sync"
)

// Fix the redis data structure code that a dual bloom filter uses.
const dbfStruct = redisHash

// DualBloomFilter instances provide a probabilistic answer to
// set membership tests with no false positives, but some degree of false negatives.
// That is, the filter will never report that it seen an element when
// when it has in fact not, but might report not having seen an element that
// it already has.  These properties are dual to a usual bloom filter,
// in that the DualBloomFilter has the properties of a bloom filter for
// the complement (in some abstract universal set) of elements it contains.
//
// Redis backing is provided by appropriately namespacing keys with the
// DualBloomFilter's ID.  For instance, a DualBloomFilter with an ID of "example"
// will insert the element x (of type []byte) into the Redis key
// id:hash(x), where hash is a configurable hash function.  DualBloomFilters
// support element expiration by utilizing the builtin Redis key expiration
// functionality.
type DualBloomFilter struct {
	capacity uint
	id       string
	expiry   int
	hashPool *sync.Pool
	store    Storer
}

// index in dual bloom filter cache for the specified element.
func (s *DualBloomFilter) index(element []byte) uint64 {
	hash := s.hashPool.Get().(hash.Hash64)
	hash.Write(element)
	i := hash.Sum64() % uint64(s.capacity)
	hash.Reset()
	s.hashPool.Put(hash)
	return i
}

// key computes the hash key identifier for the input elemeent.
func (s DualBloomFilter) key(element []byte) string {
	i := strconv.FormatUint(s.index(element), 10)
	return i
	// return strings.Join([]string{s.id, i}, ":")
}

// Has performs a probabilistic set membership test on the specified element.
// If true, the filter definitely contains the element.  If false,
// the filter may or may not contain the element.
func (s DualBloomFilter) Has(element []byte) (bool, error) {
	field := s.key(element)

	old, err := getFieldBytes(s.store, dbfStruct, s.id, field)
	log.Printf("probigo: dbf old value: %#v", old)
	if err != nil {
		return false, err
	}
	return bytesEqual(old, element), nil
}

// Add an element to the dual bloom filter.
func (s DualBloomFilter) Add(element []byte) error {
	field := s.key(element)
	_, err := setField(s.store, dbfStruct, s.id, field, element)
	return err
}

// SetHashFactory sets the hashing function factory used in the filter.
func (s *DualBloomFilter) SetHashFactory(h func() hash.Hash64) {
	s.hashPool = &sync.Pool{New: func() interface{} { return h() }}
}

// SetStore replaces the current data structure store.
func (s *DualBloomFilter) SetStore(store *store.Store) {
	s.store = store
}

// NewDualBloomFilter will create a new dual bloom filter using the provided
// data store.  If the store provided is nil, the user can later
func NewDualBloomFilter(id string, capacity uint, expiry int, store Storer) *DualBloomFilter {
	return &DualBloomFilter{
		id:       id,
		capacity: capacity,
		expiry:   expiry,
		store:    store,
		hashPool: &sync.Pool{
			New: func() interface{} { return fnv.New64a() },
		},
	}
}

// Expiration in hashes / sets etc. Hash is bf_id
// bf_id: hash(x):element
// bf_id:hash(x):ts:timestamp_value

// bf_id =  map[string]string{
// 	hash(x): element, // []byte
// 	has(x)_ts: timestamp, time.Time
// }

// Another option is to have top-level keys with expiries, appropriately namespaced.
// In the hash, or set, etc, have the key -> value pair.
// In the top level cache, have key with an expiry.  First, check if key exists
// in main Redis cache.  If so, get from the appropriate data structure. Otherwise
// we have to perform two gets, compare times, etc.

// Then, the Store interface could have an
// type Element struct {
// 	Data   []byte
// 	Key    string
// 	Expiry int
// 	Type   int // top, hash, set, hll, etc.
// }

// type Store interface {
// 	Exists(*Element) (bool, error)
// 	Get(*Element) (*Element, error)
// 	Set(*Element) error
// 	Del(*Element) error
// }
