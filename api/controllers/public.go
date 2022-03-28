package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PublicController struct {
}

func NewPublicController() PublicController {
	return PublicController{}
}

func (r *PublicController) Health(ctx *gin.Context) {
	ctx.String(http.StatusOK, "I'm here!")
}
