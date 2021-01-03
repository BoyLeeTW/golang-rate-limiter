package main

import (
	"golang-rate-limiter/internal/controllers"
	"golang-rate-limiter/internal/middlewares"
	"golang-rate-limiter/internal/pkg"
	ratelimiter "golang-rate-limiter/internal/rate-limiter"
	"golang-rate-limiter/internal/router"
	"log"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	memcachedClient := memcache.New(os.Getenv("MEMCACHED_ADDRESS"))

	if err := memcachedClient.Ping(); err != nil {
		log.Fatal("new memcached client failed")
	}

	memcachedService := pkg.NewMemcachedService(memcachedClient)

	rateLimiter := ratelimiter.NewRateLimiter(ratelimiter.RateLimiterConfig{
		PeriodSecond: 60,
		LimitCount:   60,
	}, memcachedService)

	rateLimiterController := middlewares.NewRateLimiterMiddleware(rateLimiter)
	normalController := controllers.NewNormalController()

	router := router.New(rateLimiterController, normalController)

	router.Run(":8080")
}
