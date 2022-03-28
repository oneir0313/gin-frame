package routes

import (
	"gin-frame/api/controllers"
)

type JobRoutes struct {
	handler       controllers.Handler
	jobController controllers.JobController
}

func (r JobRoutes) Setup() {
	r.handler.Gin.PUT("/job", r.jobController.AddJob)
	r.handler.Gin.GET("/jobs", r.jobController.GetJobs)
	r.handler.Gin.DELETE("/job/:id", r.jobController.DeleteJob)
}

func NewJobRoutes(
	handler controllers.Handler,
	jobController controllers.JobController,
) JobRoutes {
	return JobRoutes{
		handler:       handler,
		jobController: jobController,
	}
}
