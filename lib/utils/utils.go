package utils

import (
	"github.com/gin-gonic/gin"
)

func GetRealIP(ctx *gin.Context) string {
	r := ctx.Request
	IPAddress := r.Header.Get("X-Real-IP")
	if IPAddress == "" {
		IPAddress = ctx.ClientIP()
	}
	return IPAddress
}
