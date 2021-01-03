package router

import (
	"golang-rate-limiter/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	engine              *gin.Engine
	normalController    interfaces.NormalControllerInterface
	rateLimitController interfaces.RateLimiterMiddlewareInterface
}

func New(rateLimiterMiddleware interfaces.RateLimiterMiddlewareInterface, normalController interfaces.NormalControllerInterface) *Router {
	r := gin.Default()
	r.LoadHTMLGlob("internal/view/*")

	r.GET("/", rateLimiterMiddleware.CheckIPLimit, normalController.Get)

	return &Router{
		engine:              r,
		normalController:    normalController,
		rateLimitController: rateLimiterMiddleware,
	}
}

func (r *Router) Run(port string) error {
	server := http.Server{
		Handler: r.engine,
		Addr:    port,
	}

	return server.ListenAndServe()
}
