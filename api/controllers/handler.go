package controllers

import (
	"net/http"

	configmanager "gin-frame/lib/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Gin *gin.Engine
}

func NewHandler() Handler {
	engine := gin.Default()
	corsConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"*", "Authorization", "Content-Type", "Origin", "Content-Length"},

		// firefox 和 safari 不支援 *, 所以需要一個一個打，但更好的是要抓 Access-Control-Request-Headers
		// https://stackoverflow.com/questions/54666673/cors-check-fails-for-firefox-but-passes-for-chrome
	}

	engine.Use(cors.New(corsConfig))

	if configmanager.Global.Env == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "this url is not found on this service")
	})
	return Handler{Gin: engine}
}
