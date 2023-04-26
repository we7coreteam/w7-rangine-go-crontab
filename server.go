package crontab

import (
	"github.com/robfig/cron/v3"
	"github.com/we7coreteam/w7-rangine-go-support/src/server"
)

type Server struct {
	server.Server

	Cron *cron.Cron
}

func NewDefaultServer() *Server {
	return &Server{
		Cron: cron.New(cron.WithParser(cron.NewParser(
			cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
		))),
	}
}

func (s *Server) GetServerName() string {
	return "crontab"
}

func (s *Server) GetOptions() map[string]string {
	return nil
}

func (s *Server) RegisterTask(rule string, task Task) {
	_, err := s.Cron.AddFunc(rule, task.Run)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Start() {
	s.Cron.Run()
}
