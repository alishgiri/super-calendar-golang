package services

import (
	"time"

	"github.com/go-co-op/gocron/v2"
)

func ScheduleCronJob(task gocron.Task) {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	_, err = s.NewJob(gocron.DurationJob(10*time.Second), task)
	if err != nil {
		panic(err)
	}

	s.Start()
}
