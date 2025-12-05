package url

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Redirect(ctx *gin.Context) {
	shortCode := ctx.Param("short_code")

	if shortCode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing short code"})
		return
	}

	longUrl, err := h.service.GetLongUrl(ctx, shortCode)
	if err != nil {
		if err == ErrShortNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "short url not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to redirect"})
		return
	}

	ctx.Redirect(http.StatusFound, longUrl)
}

func (h *Handler) CreateShortUrl(ctx *gin.Context) {
	var body CreateShortUrlBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	row, err := h.service.CreateShortUrl(ctx, body.LongUrl)
	if err != nil {

		if err == ErrLongAlreadyExists {
			ctx.JSON(http.StatusConflict, gin.H{
				"error": "long url already exists",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create short url",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":    "success",
		"short_url": row.ShortCode,
	})
}
