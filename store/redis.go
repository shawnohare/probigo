package store

import "github.com/garyburd/redigo/redis"

// redisConn is a Redis backed implementation of the Conn interface.
type redisConn struct {
	pool *redis.Pool
}

// Do wraps the redigo redis.Conn Do method by fetching a connection from
// the connection pool, defering connection closure, and issuing the
// appropriate command.
func (r *redisConn) Do(cmd string, args ...interface{}) (interface{}, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return conn.Do(cmd, args...)
}

// NewRedis creates a Redis backed Store instance.
func NewRedis(p *redis.Pool) *Store {
	rd := &redisConn{pool: p}
	return New(rd)
}
