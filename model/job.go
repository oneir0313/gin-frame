package model

type Job struct {
	Schedule string `json:"schedule" binding:"required" `
	Message  string `json:"message" binding:"required"`
}
