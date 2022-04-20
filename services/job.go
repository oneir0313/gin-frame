package services

import (
	"context"
	"fmt"
	"time"

	"gin-frame/model"

	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
)

type JobService struct {
	cron *cron.Cron
	jobs []jobEntry
}

type jobEntry struct {
	EntryID int `json:"entryID"`
	model.Job
}

func NewJobService() JobService {
	utc8, _ := time.LoadLocation("Asia/Taipei")
	schedule := cron.New(cron.WithLocation(utc8), cron.WithSeconds())
	schedule.Start()

	tmpJobManager := JobService{
		cron: schedule,
	}

	return tmpJobManager
}

func (r *JobService) AddJob(job model.Job) (cron.EntryID, error) {

	entryId, err := r.cron.AddFunc(job.Schedule, func() {
		log.Info().Msgf("Job number is %d", job.Number)
	})

	if err != nil {
		return 0, err
	}

	addJob := jobEntry{
		EntryID: int(entryId),
		Job:     job,
	}

	r.jobs = append(r.jobs, addJob)
	return entryId, nil
}

func (r *JobService) GetJobs() []jobEntry {
	return r.jobs
}

func (r *JobService) GetJob(id int) cron.Entry {
	return r.cron.Entry(cron.EntryID(id))
}

func (r *JobService) Stop() context.Context {
	return r.cron.Stop()
}

func (r *JobService) Remove(id cron.EntryID) error {
	deleteIdx := -1
	for idx, item := range r.jobs {
		if cron.EntryID(item.EntryID) == id {
			deleteIdx = idx
			break
		}
	}
	if deleteIdx == -1 {
		return fmt.Errorf("Job is not found")
	}
	r.cron.Remove(id)
	r.jobs = append(r.jobs[:deleteIdx], r.jobs[deleteIdx+1:]...)
	return nil
}
