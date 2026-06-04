package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viettrung2103/bookmark-management-lesson/internal/service"
)

const passwordLength = 12

// GenPass represents the genpass handler
type GenPass interface {
	GeneratePassword(c *gin.Context)
}
type genPassHandler struct {
	genPassService service.GenPass
}

// NewGenPass creates a new genpass handler
func NewGenPass(genPassSvc service.GenPass) GenPass {
	return &genPassHandler{
		genPassService: genPassSvc,
	}
}

// GeneratePassword generates a new password
// @Summary Generate a new password
// @Description Generate a new password
// @Tags password
// @Success 200 {object} string "12345678"
// @Router /genpass [get]
func (s *genPassHandler) GeneratePassword(c *gin.Context) {
	pass, err := s.genPassService.GeneratePassword(passwordLength)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Err"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"password": pass})
}
