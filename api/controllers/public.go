package controllers

import (
	"fmt"
	"net/http"
	"time"

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

func (r *PublicController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong " + fmt.Sprint(time.Now().Unix()))
}
