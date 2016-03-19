# probigo
Probabilistic data structures in Go with Redis (and more generic) backends.

# Stores

Anything satisfying the `Store` interface defined in this package can be used
as a backend for the probabilistic data structures.  The `Store` interface
is still a bit unstable, so we recommend using the provided `Redis` store.

## Redis

Redis is the main and default data stucture store this package has in mind.
A user can either provide Redis connection pools or use the package's 
`NewRedisPool` helper function.  For example:
```go
var (
  pool *redis.Pool
  r *Redis
)
var pool is
pool = NewRedisPool(nil)
r = NewRedis(pool)
```
will create a new Redis-backed data store with the package's default 
configuration constants: `RedisNetwork`, `RedisAddress`, etc. 
