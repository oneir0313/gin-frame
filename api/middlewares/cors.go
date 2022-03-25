package middlewares

import (
	"gin-frame/api/controllers"

	"github.com/gin-contrib/cors"
)

type CorsMiddleware struct {
	handler controllers.Handler
}

func NewCorsMiddleware(handler controllers.Handler) CorsMiddleware {
	return CorsMiddleware{
		handler: handler,
	}
}

func (m CorsMiddleware) Setup() {
	corsConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"*", "Authorization", "Content-Type", "Origin", "Content-Length"},

		// firefox 和 safari 不支援 *, 所以需要一個一個打，但更好的是要抓 Access-Control-Request-Headers
		// https://stackoverflow.com/questions/54666673/cors-check-fails-for-firefox-but-passes-for-chrome
	}
	m.handler.Gin.Use(cors.New(corsConfig))
}
