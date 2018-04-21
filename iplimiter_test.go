package iplimiter

import (
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"time"
	"net/http"
)

func TestIpLimiter(t *testing.T) {
	r := gin.Default()

	r.Use(NewRateLimiterMiddleware(redis.NewClient(
		&redis.Options{
			DB:       0,
			Password: "",
			Addr:     "127.0.0.1:6379",
		},
	), "test", 100, time.Second*5))

	r.GET("/", func(c *gin.Context) {
		c.String(200,"OK")
	})

	go r.Run(":9999")

	for i := 0; i < 102; i++ {
		c := &http.Client{}

		resp, e := c.Get("http://127.0.0.1:9999")
		if e != nil {
			t.Error("Error during requests ", e.Error())
			return
		}

		switch {
		case i < 100:
			break;
		case i == 100:
			if resp.StatusCode != 429 {
				t.Error("Threashold break not detected")
			} else {
				time.Sleep(time.Second * 5)
			}
			break;
		case i == 101:
			if resp.StatusCode != 200 {
				t.Error("Unnecessary block")
			}
			break;
		}
	}
}
