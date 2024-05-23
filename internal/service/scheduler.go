package service

import (
	"github.com/jasonlvhit/gocron"
)

type Scheduler struct {
	*gocron.Scheduler
}

func NewScheduler() *Scheduler {
	return &Scheduler{gocron.NewScheduler()}
}
