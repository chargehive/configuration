package scheduler

import "time"

// Schedule is a single schedule item for a single attempt
type Schedule struct {
	// AttemptConfig is the configuration to use when processing this schedule
	AttemptConfig AttemptConfig

	// TimeDelay is the amount of time to wait before processing after TimeDelayOrigin
	TimeDelay time.Duration

	// TimeDelayOrigin specifies when the time origin is based from
	TimeDelayOrigin AttemptOriginType

	// TimeDelaySync indicates the direction in time that the schedule should move when syncing to TimeSyncHour
	TimeDelaySync TimeDelaySync

	// TimeSyncHour is an hour designation (1-24) i.e 2 == 2AM | where less than 1 indicates that this value is not set
	TimeSyncHour int

	// TimeSyncZone indicates the timezone that the TimeSyncHour is relative to
	TimeSyncZone TimeZone
}

func (Schedule) Validate() {

}
