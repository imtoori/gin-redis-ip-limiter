# gin-redis-ip-limiter

It works with redis. It's pluggable and all you need is a redis client for every limiter.
It limits path access based on client ip-address.

## Example

```go
package main

import (
  "github.com/Salvatore-Giordano/gin-redis-ip-limiter"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(iplimiter.CreateRateLimiterMiddleware(redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	}), "general", 200, 60*time.Second.Nanoseconds()))
  // ...
  r.Run(":8080")
}
```
**Key**: this is the key used to save data on redis. Data are saved as ip:key.

**Limit**: number of request to accept.

**SlidingWindowNanoseconds**: duration in nanosecond of the sliding window to consider.

