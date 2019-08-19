package scheduler

import "time"

type Schedule struct {
	AttemptConfig   AttemptConfig
	TimeDelayOrigin AttemptOriginType
	TimeDelaySync   TimeDelaySync
	TimeDelay       time.Duration
	TimeSyncHour    int
	TimeSyncZone    TimeZone
}
