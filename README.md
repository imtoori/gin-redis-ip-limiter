## Description

It works with redis. It's pluggable and all you need is a redis client for every limiter.
It limits path access based on client ip-address.

It's inspired by this article https://engagor.github.io/blog/2017/05/02/sliding-window-rate-limiter-redis/.

It uses this golang redis library https://github.com/go-redis/redis.

## Installation

Just run 

`go get -u github.com/Salvatore-Giordano/gin-redis-ip-limiter/`

## Example

```go
package main

import (
  "github.com/Salvatore-Giordano/gin-redis-ip-limiter"
  "github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()
  r.Use(iplimiter.NewRateLimiterMiddleware(redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	}), "general", 200, 60*time.Second))
  // ...
  r.Run(":8080")
}
```
**Key**: this is the key used to save data on redis. Data are saved as ip:key.

**Limit**: number of request to accept.

**SlidingWindow**: duration of the sliding window to consider.

