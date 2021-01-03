package ratelimiter

import (
	"golang-rate-limiter/internal/interfaces"
	"golang-rate-limiter/internal/models"
)

type RateLimiter struct {
	config           RateLimiterConfig
	memcachedService interfaces.MemcachedServiceInterface
}

type RateLimiterConfig struct {
	PeriodSecond int32
	LimitCount   uint64
}

func NewRateLimiter(config RateLimiterConfig, memcachedService interfaces.MemcachedServiceInterface) *RateLimiter {
	return &RateLimiter{
		config:           config,
		memcachedService: memcachedService,
	}
}

func (r *RateLimiter) CheckExceed(ip string) (models.RateLimiterResponse, error) {
	latestValue, err := r.memcachedService.IncrementOrAdd(ip, 1, r.config.PeriodSecond)
	if err != nil {
		return models.RateLimiterResponse{}, err
	}

	return models.RateLimiterResponse{
		CurrentCount: latestValue,
		IsExceed:     latestValue > r.config.LimitCount,
	}, nil
}
