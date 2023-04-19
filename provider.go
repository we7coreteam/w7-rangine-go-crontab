package crontab

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/we7coreteam/w7-rangine-go-support/src/provider"
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
	provider.Abstract
}

func (provider *Provider) Register() {
	logger, err := provider.GetLoggerFactory().Channel("default")
	if err == nil {
		var log cron.Logger
		if provider.GetConfig().GetString("app.env") == "debug" {
			log = cron.VerbosePrintfLogger(Logger{
				log: logger,
			})
		} else {
			log = cron.PrintfLogger(Logger{
				log: logger,
			})
		}
		cron.WithLogger(log)(GetCrontab())
	}

	provider.RegisterServer(&Server{})
}
