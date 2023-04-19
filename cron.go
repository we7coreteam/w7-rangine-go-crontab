package crontab

import "github.com/robfig/cron/v3"

var crontab = cron.New(cron.WithSeconds())

func SetCrontab(cron *cron.Cron) {
	crontab = cron
}

func GetCrontab() *cron.Cron {
	return crontab
}

func RegisterTask(rule string, task func()) {
	_, err := crontab.AddFunc(rule, task)
	if err != nil {
		panic(err)
	}
}
