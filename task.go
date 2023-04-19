package crontab

type Task interface {
	Run()
}
