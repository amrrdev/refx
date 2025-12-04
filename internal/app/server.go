package app

import (
	"github.com/amrrdev/refx/internal/url"
	"github.com/gin-gonic/gin"
)

func NewServer(urlHandler *url.Handler) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	url.RegisterRoutes(api, urlHandler)

	return r
}
