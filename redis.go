package probigo

import (
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
)

// Default Redis configuration.
const (
	RedisMaxIdle            int    = 3
	RedisIdleTimeoutSeconds int    = 240
	RedisNetwork            string = "tcp"
	RedisAddress            string = ":6379"
)

// Define some data structure types typically supported by Store interfaces.
const (
	redisTop = iota + 1
	redisHash
	redisSet
	redisOrderedSet
	redisHyperLogLog
)

// RedisConfig is a helper struct that allows users to more easily create
// a redis.Pool instance to be used with the built-in Redis backing.
type RedisConfig struct {
	MaxIdle            int
	IdleTimeoutSeconds time.Duration
	Network            string
	Address            string
}

// map of (base command, struct type) pairs to redis command.
var redisCommands = map[int]map[string]string{
	redisTop: map[string]string{
		"get":    "GET",
		"set":    "SET",
		"del":    "DEL",
		"exists": "EXISTS",
	},
	redisHash: map[string]string{
		"get":    "HGET",
		"set":    "HSET",
		"del":    "HDEL",
		"exists": "HEXISTS",
	},
	redisSet: map[string]string{
		"get":    "SISMEMBER",
		"set":    "SADD",
		"del":    "SREM",
		"exists": "SISMEMBER",
	},
	redisOrderedSet: map[string]string{
		"get":    "ZSCORE",
		"set":    "ZADD",
		"del":    "ZREM",
		"exists": "ZSCORE",
	},
}

// cmd converts the base command into the data structure specific Redis version.
func redisCmd(structType int, base string) (string, error) {
	if cmd, ok := redisCommands[structType][base]; ok {
		return cmd, nil
	}
	return "", errors.New(ErrInvalidCommand)
}

// Redis backend for the probabilistic data structures.
// A user can initialize by specifying the Redis pool to use or via the
// package's NewRedisBack constructor.
type Redis struct {
	Pool *redis.Pool
}

func (r Redis) Do(cmd string, args ...interface{}) (interface{}, error) {
	conn := r.Pool.Get()
	return conn.Do(cmd, args...)
}

// func (r redis) args(d *Data) []interface{} {
// 	if args == nil {
// 		return nil
// 	}

// 	var xs []interface{}
// 	if d.ID != "" {
// 		xs = append(xs, d.ID)
// 	}
// 	if d.Key != "" {
// 		xs = append(xs, x.Key)
// 	}
// }

// func (r redis) get(key string) ([]byte, error) {
// 	conn := pool.Get()
// 	return redis.Bytes(conn.Do("GET", key))
// }

// func (r redis) set(key string, value []byte, expiry int) error {
// 	conn := pool.Get()
// 	args := []interface{}{key, value}
// 	if expiry > 0 {
// 		args = append(args, "EX", expires)
// 	}
// 	_, err := conn.Do("SET", args...)
// 	return err
// }

// func (r Redis) del(key string) error {
// 	conn := r.pool.Get()
// 	_, err := conn.Do("DEL", key)
// 	return err
// }

// func (r Redis) exists(key string) (bool, error) {
// 	conn := r.pool.Get()
// 	return redis.Bool(conn.Do("EXISTS", key))
// }

// func (r redis) hget(key string, field string) ([]byte, error) {
// 	conn := r.pool.Get()
// 	return redis.Bytes(conn.Do("HGET", key, value))
// }

// func (r redis) hset(key string, field string, value []byte) error {
// 	conn := r.pool.Get()
// 	_, err := conn.Do("HSET", key, field, value)
// 	return err
// }

// func (r Redis) hdel(key string, field string) error {
// 	conn := r.pool.Get()
// 	_, err := conn.Do("HDEL", key, field)
// 	return err
// }

// func (r Redis) hexists(key string, field string) (bool, error) {
// 	conn := r.pool.Get()
// 	return redis.Bool(conn.Do("HEXISTS", key, field))
// }

// expired checks whether a nested key is expired
// func (r Redis) isExpired(id string, key string) (bool, error) {
// 	topKey := strings.Join([]string{id, key, "ex"}, ":")
// 	ok, err := exists(r.store, topKey)
// 	return !ok, err
// }

// rkey returns an appropriate top level key.
// func (r Redis) rkey(d *Data) string {
// 	var key string
// 	if d.Struct == Top {
// 		key = strings.Join([]string{d.ID, d.Key}, ":")
// 	} else {
// 		key = d.ID
// 	}

// 	return key
// }

// func (r Redis) Get(input interface{}) (interface{}, error) {
// 	var req *request
// 	switch input.(type) {
// 	case *request:
// 		req = input.(*request)
// 	default:
// 		return nil, errors.New(ErrBadGetInput)
// 	}

// 	conn := r.pool.Get()
// 	d := req.data

// 	if d.pType == redisTop {
// 		key := strings.Join([]string{d.id, d.key}, ":")
// 		return conn.Do("GET", key)
// 	}

// 	if req.checkIfExpired {
// 		expired, err := r.isExpired(d.id, d.key)
// 		if err != nil {
// 			return nil, err
// 		}
// 		if expired {
// 			err := r.Del(req)
// 			if err != nil {
// 				return nil, err
// 			}
// 		}
// 	}

// 	cmd := r.cmd(d.pType, "get")
// 	return conn.Do(cmd, d.id, d.key)
// }

// func (r Redis) Exists(input interface{}) (bool, error) {
// 	var req *request
// 	switch input.(type) {
// 	case *request:
// 		req = input.(*request)
// 	default:
// 		return nil, errors.New(ErrBadGetInput)
// 	}

// 	conn := r.pool.Get()
// 	d := req.data
// 	cmd := r.cmd(req.data.pType, "exists")
// 	if d.pType == redisTop {
// 		key := strings.Join([]string{d.id, d.key}, ":")
// 		return redis.Bool(conn.Do(cmd, key))
// 	}
// 	return redis.Bool(conn.Do(cmd, d.id, d.key))
// }

// func (r Redis) Del(key string) error {
// 	return redisDel(r.pool, key)
// }

// func redisDel(pool *redis.Pool, keys ...string) error {
// 	if pool == nil {
// 		return errors.New(ErrNilRedisPool)
// 	}
// 	conn := pool.Get()
// 	args := make([]interface{}, len(keys))
// 	for i, key := range keys {
// 		args[i] = key
// 	}
// 	_, err := conn.Do("DEL", args...)
// 	return err
// }

// func redisSet(pool *redis.Pool, key string, value []byte, expires int) error {
// 	if pool == nil {
// 		return errors.New(ErrNilRedisPool)
// 	}
// 	conn := pool.Get()
// 	args := []interface{}{key, value}
// 	if expires > 0 {
// 		args = append(args, "EX", expires)
// 	}
// 	_, err := conn.Do("SET", args...)
// 	return err
// }

// func redisGetSet(pool *redis.Pool, key string, value []byte) ([]byte, error) {
// 	if pool == nil {
// 		return nil, errors.New(ErrNilRedisPool)
// 	}
// 	conn := pool.Get()
// 	return redis.Bytes(conn.Do("GETSET", key, value))
// }

// func redisGet(pool *redis.Pool, key string) ([]byte, error) {
// 	if pool == nil {
// 		return nil, errors.New(ErrNilRedisPool)
// 	}
// 	conn := pool.Get()
// 	return redis.Bytes(conn.Do("GET", key))
// }

// func redisExists(pool *redis.Pool, key string) (bool, error) {
// 	if pool == nil {
// 		return false, errors.New(ErrNilRedisPool)
// 	}
// 	conn := pool.Get()
// 	return redis.Bool(conn.Do("EXISTS", key))
// }

// NewRedisPool is a helper function that creates a configured Redis pool.
// If the configuration is nil, default values will be used.
func NewRedisPool(in *RedisConfig) *redis.Pool {
	conf := &RedisConfig{
		MaxIdle:            RedisMaxIdle,
		IdleTimeoutSeconds: time.Duration(RedisIdleTimeoutSeconds),
		Network:            RedisNetwork,
		Address:            RedisAddress,
	}

	if in != nil {
		if in.MaxIdle > 0 {
			conf.MaxIdle = in.MaxIdle
		}
		if in.IdleTimeoutSeconds > 0 {
			conf.IdleTimeoutSeconds = in.IdleTimeoutSeconds
		}
		if in.Network != "" {
			conf.Network = in.Network
		}
		if in.Address != "" {
			conf.Address = in.Address
		}
	}

	pool := &redis.Pool{
		MaxIdle:     conf.MaxIdle,
		IdleTimeout: conf.IdleTimeoutSeconds * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(conf.Network, conf.Address)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return pool
}

// NewRedis creates a new Redis backend that implements the Store interface.
func NewRedis(pool *redis.Pool) *Redis {
	return &Redis{Pool: pool}
}
