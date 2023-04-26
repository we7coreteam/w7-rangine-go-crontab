package crontab

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/we7coreteam/w7-rangine-go-support/src/facade"
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
	crontabServer := NewDefaultServer()

	logger, err := facade.GetLoggerFactory().Channel("default")
	if err == nil {
		var log cron.Logger
		if facade.GetConfig().GetString("app.env") == "debug" {
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

	facade.RegisterServer(crontabServer)
}
