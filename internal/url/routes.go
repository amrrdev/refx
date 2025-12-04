package url

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	urls := r.Group("/urls")
	{
		urls.GET("/:short_code", h.Redirect)
		urls.POST("", h.CreateShortUrl)
	}
}
