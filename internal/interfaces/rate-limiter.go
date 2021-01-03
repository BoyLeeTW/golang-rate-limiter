package interfaces

import "golang-rate-limiter/internal/models"

type RateLimiterInterface interface {
	CheckExceed(key string) (models.RateLimiterResponse, error)
}
