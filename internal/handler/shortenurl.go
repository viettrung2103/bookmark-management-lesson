package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/viettrung2103/bookmark-management-lesson/internal/service"
)

type ShortenUrlHandler interface {
	ShortenUrl(c *gin.Context)
	Redirect(c *gin.Context)
}

type shortenUrlHandler struct {
	shortenUrlService service.ShortenUrl
}

func NewShortenUrlHandler(shortenUrlSvc service.ShortenUrl) ShortenUrlHandler {
	return &shortenUrlHandler{
		shortenUrlService: shortenUrlSvc,
	}
}

type shortenUrlRequest struct {
	URL    string `json:"url" binding:"required,url"`
	Expiry int64  `json:"exp" binding:"required"`
}

type shortenUrlResponse struct {
	Code string `json:"code"`
}

// ShortenUrl shorten the url to code
// @Summary receive the url, return the code
// @Tags link
// @Accept application/json
// @Produce application/json
// @Param request body shortenUrlRequest true "Shorten URL Input payload"
// @Success 200 {object} string
// @Router /v1/links/shorten [post]
func (h *shortenUrlHandler) ShortenUrl(c *gin.Context) {
	request := &shortenUrlRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input error"})
		return
	}

	key, err := h.shortenUrlService.ShortenUrl(c, request.URL, request.Expiry)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": key, "exp": request.Expiry})
}

// Redirect Forward the request to the original url
// @Tags link
// @Accept application/json
// @Produce application/json
// @Param code path string true "code"
// @Success	302
// @Router /v1/links/shorten/{code} [get]
func (h *shortenUrlHandler) Redirect(c *gin.Context) {
	code := c.Param("code")
	if code == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "code is required"})
	}

	// call service to get url from code
	url, err := h.shortenUrlService.GetURL(c, code)
	if err != nil {
		if errors.Is(err, service.UrlNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "url not found"})
			return
		}
		log.Error().Err(err).Str("from", "handler.shortenurl.Redirect").Msg("failed to get url from code")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// redirect to url
	c.Redirect(http.StatusFound, url)

}
