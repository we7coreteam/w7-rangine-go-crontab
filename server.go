package crontab

import (
	"github.com/we7coreteam/w7-rangine-go-support/src/server"
)

type Server struct {
	server.Server
}

func (Server) GetServerName() string {
	return "crontab"
}

func (Server) GetOptions() map[string]string {
	return nil
}

func (Server) Start() {
	GetCrontab().Run()
}
