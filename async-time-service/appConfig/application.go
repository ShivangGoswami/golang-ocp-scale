package appConfig

import "time"

type applicationProperties struct {
	AppUrl            string
	SchedulerDelay    time.Duration
	SchedulerInterval time.Duration
}

var Config = applicationProperties{"127.0.0.1:8080", time.Second * 7, time.Second * 23}
