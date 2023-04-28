package crontab

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"github.com/we7coreteam/w7-rangine-go-support/src/logger"
	"github.com/we7coreteam/w7-rangine-go-support/src/server"
	"go.uber.org/zap"
)

type Logger struct {
	log *zap.Logger
}

func (l Logger) Printf(format string, v ...any) {
	if l.log != nil {
		l.log.Info(fmt.Sprintf(format, v...))
	}
}

type Provider struct {
	server *Server
}

func (provider *Provider) Register(config *viper.Viper, loggerFactory logger.Factory, serverManager server.Manager) *Provider {
	crontabServer := NewDefaultServer()

	logger, err := loggerFactory.Channel("default")
	if err == nil {
		var log cron.Logger
		if config.GetString("app.env") == "debug" {
			log = cron.VerbosePrintfLogger(Logger{
				log: logger,
			})
		} else {
			log = cron.PrintfLogger(Logger{
				log: logger,
			})
		}
		cron.WithLogger(log)(crontabServer.Cron)
	}
	provider.server = crontabServer

	serverManager.RegisterServer(crontabServer)

	return provider
}

func (provider *Provider) Export() *Server {
	return provider.server
}
