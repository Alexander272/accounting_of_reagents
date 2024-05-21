package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/Alexander272/accounting_of_reagents/backend/internal/config"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/error_bot"
	"github.com/Alexander272/accounting_of_reagents/backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron/v2"
)

type SchedulerService struct {
	cron    gocron.Scheduler
	reagent Reagent
}

func NewSchedulerService(reagent Reagent) *SchedulerService {
	cron, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalf("failed to create new scheduler. error: %s", err.Error())
	}

	return &SchedulerService{
		cron:    cron,
		reagent: reagent,
	}
}

type Scheduler interface {
	Start(conf *config.SchedulerConfig) error
	Stop() error
}

// запуск заданий в cron
func (s *SchedulerService) Start(conf *config.SchedulerConfig) error {
	now := time.Now()
	jobStart := time.Date(now.Year(), now.Month(), now.Day(), conf.StartTime, 0, 0, 0, now.Location())
	if now.Hour() >= conf.StartTime {
		jobStart = jobStart.Add(24 * time.Hour)
	}
	// jobStart := now.Add(1 * time.Minute)
	logger.Info("starting jobs time " + jobStart.Format("02.01.2006 15:04:05"))

	job := gocron.DurationJob(conf.Interval)
	task := gocron.NewTask(s.job)
	jobStartAt := gocron.WithStartAt(gocron.WithStartDateTime(jobStart))

	_, err := s.cron.NewJob(job, task, jobStartAt)
	if err != nil {
		return fmt.Errorf("failed to create new job. error: %w", err)
	}

	//? запуск крона через интервал (по умолчанию день)
	s.cron.Start()
	return nil
}

// остановка заданий в cron
func (s *SchedulerService) Stop() error {
	if err := s.cron.Shutdown(); err != nil {
		return fmt.Errorf("failed to shutdown cron scheduler. error: %w", err)
	}
	return nil
}

func (s *SchedulerService) job() {
	logger.Debug("job was started")

	if err := s.reagent.DeleteOld(context.Background()); err != nil {
		error_bot.Send(
			&gin.Context{
				Request: &http.Request{
					Method: "Get",
					URL:    &url.URL{Host: "cron", Path: "delete-old-reagent"},
				},
			}, err.Error(), nil,
		)
	}

	if err := s.reagent.SendOverdue(context.Background()); err != nil {
		error_bot.Send(
			&gin.Context{
				Request: &http.Request{
					Method: "Get",
					URL:    &url.URL{Host: "cron", Path: "send-overdue"},
				},
			},
			err.Error(), nil,
		)
	}
}
