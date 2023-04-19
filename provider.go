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
		cron.DefaultLogger = cron.VerbosePrintfLogger(Logger{
			log: logger,
		})
	}

	provider.RegisterServer(&Server{})
}
