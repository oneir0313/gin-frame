package controllers

import (
	"gin-frame/model"
	"gin-frame/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/robfig/cron/v3"
)

type JobController struct {
	service    services.JobService
}

func NewJobController(service services.JobService) JobController {
	return JobController{
		service: service,
	}
}

func (r *JobController) AddJob(ctx *gin.Context) {
	var req model.Job
	if err := ctx.ShouldBindWith(&req, binding.JSON); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	entryID, err := r.service.AddJob(req)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"entryID": entryID})
}

func (r *JobController) GetJobs(ctx *gin.Context) {
	jobs := r.service.GetJobs()
	if len(jobs) == 0 {
		ctx.JSON(http.StatusOK, []string{})
		return
	}
	ctx.JSON(http.StatusOK, jobs)
}

func  (r *JobController) DeleteJob(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = r.service.Remove(cron.EntryID(id))
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
	}
	ctx.Status(http.StatusOK)
}
