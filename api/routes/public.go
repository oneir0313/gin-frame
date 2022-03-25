package routes

import (
	"gin-frame/api/controllers"
)

type PublicRoutes struct {
	handler       controllers.Handler
	publicController controllers.PublicController
}

func (r PublicRoutes) Setup() {
	r.handler.Gin.GET("/health", r.publicController.Health)
}

func NewPublicRoutes( handler controllers.Handler, publicController controllers.PublicController ) PublicRoutes {
	return PublicRoutes{
		handler:      handler,
		publicController:     publicController,
	}
}