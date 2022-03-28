package model

type Job struct {
	Schedule string `json:"schedule" binding:"required" `
	Number   int    `json:"number" binding:"required"`
}
