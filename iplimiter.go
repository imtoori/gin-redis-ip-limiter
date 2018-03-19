package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func CreateRateLimiterMiddleware(redisClient *redis.Client, key *string, limit int, slidingWindowNanoseconds int64) gin.HandlerFunc {

	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprint("error init redis", err.Error()))
	}

	if key == nil {
		*key = ""
	}

	return func(c *gin.Context) {
		now := time.Now().UnixNano()
		userCntKey := fmt.Sprint(c.ClientIP(), ":", key)

		redisClient.ZRemRangeByScore(userCntKey,
			"0",
			fmt.Sprint(now-(slidingWindowNanoseconds))).Result()

		reqs, _ := redisClient.ZRange(userCntKey, 0, -1).Result()

		if len(reqs) >= limit {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"status":  http.StatusTooManyRequests,
				"message": "too many request",
			})
			return
		}

		c.Next()
		redisClient.ZAddNX(userCntKey, redis.Z{Score: float64(now), Member: float64(now)})
		redisClient.Expire(userCntKey, time.Duration(slidingWindowNanoseconds)*time.Second)
	}

}
