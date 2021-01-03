package middlewares

import (
	"golang-rate-limiter/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RateLimiterMiddleware struct {
	rateLimiter interfaces.RateLimiterInterface
}

func NewRateLimiterMiddleware(rateLimiter interfaces.RateLimiterInterface) *RateLimiterMiddleware {
	return &RateLimiterMiddleware{
		rateLimiter: rateLimiter,
	}
}

func (r *RateLimiterMiddleware) CheckIPLimit(c *gin.Context) {
	response, err := r.rateLimiter.CheckExceed(c.ClientIP())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message:": "IntervalServerError",
		})
		return
	}

	if response.IsExceed {
		c.HTML(http.StatusTooManyRequests, "error.html", nil)
		c.Abort()
		return
	}

	c.Set("ip_request_count", response.CurrentCount)

	c.Next()
}
