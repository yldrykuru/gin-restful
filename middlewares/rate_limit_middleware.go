package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimitMiddleware(limit int, duration time.Duration) gin.HandlerFunc {
	rateLimiter := ratelimit.NewBucketWithRate(float64(limit), int64(limit))

	return func(c *gin.Context) {
		if rateLimiter.TakeAvailable(1) <= 0 {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		c.Next()
	}
}
