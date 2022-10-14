package routes

import (
	"gin-frame/api/controllers"
	"io"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type PublicRoutes struct {
	handler          controllers.Handler
	publicController controllers.PublicController
}

func (r PublicRoutes) Setup() {
	r.handler.Gin.GET("/health", r.publicController.Health)
	r.handler.Gin.GET("/ping", logger.SetLogger(
		logger.WithUTC(true),
		logger.WithLogger(func(c *gin.Context, out io.Writer, latency time.Duration) zerolog.Logger {
		  return zerolog.New(out).With().
			Str("foo", "bar").
			Str("path", c.Request.URL.Path).
			Dur("latency", latency).
			Logger()
		}),
	  ), r.publicController.Ping)
}

func NewPublicRoutes(handler controllers.Handler, publicController controllers.PublicController) PublicRoutes {
	return PublicRoutes{
		handler:          handler,
		publicController: publicController,
	}
}
