package scheduler

import "time"

// Schedule is a single schedule item for a single attempt
type Schedule struct {
	// AttemptConfig is the configuration to use when processing this schedule
	AttemptConfig AttemptConfig `json:"attemptConfig" yaml:"attemptConfig" validate:"required,dive"`

	// TimeDelay is the amount of time to wait before processing after TimeDelayOrigin
	TimeDelay time.Duration `json:"timeDelay" yaml:"timeDelay" validate:"gte=0"`

	// TimeDelayOrigin specifies when the time origin is based from
	TimeDelayOrigin AttemptOriginType `json:"timeDelayOrigin" yaml:"timeDelayOrigin" validate:"oneof=initialisation last-failure"`

	// TimeDelaySync indicates the direction in time that the schedule should move when syncing to TimeSyncHour
	TimeDelaySync TimeDelaySync `json:"timeDelaySync" yaml:"timeDelaySync" validate:"oneof=None Earliest Latest Closest"`

	// TimeSyncHour is an hour designation (0-23) i.e 2 == 2AM. Ignored if TimeDelaySync is set to None
	TimeSyncHour int `json:"timeSyncHour" yaml:"timeSyncHour" validate:"min=0,max=23"`

	// TimeWindowHours is the number of hours grace for a transaction to be processed within.  Allowing for a transaction to be processed within X hours of the requested delay & sync hour e.g. 10am > 11am with a TimeWindowHours of 1
	TimeWindowHours int `json:"timeWindowHours" yaml:"timeWindowHours" validate:"gte=0"`

	// TimeSyncZone indicates the timezone that the TimeSyncHour is relative to. Ignored if TimeDelaySync is set to None
	TimeSyncZone TimeZone `json:"timeSyncZone" yaml:"timeSyncZone" validate:"oneof=ULT UTC CIT"`

	// DayOfMonth is the day of the month to process the schedule
	DayOfMonth int `json:"dayOfMonth" yaml:"dayOfMonth" validate:"min=0,max=28"`

	// DayOfWeek is the day of the week to process the schedule (1 = Monday) - Sunday moved to 7 to leave 0 as an ignored value
	DayOfWeek int `json:"dayOfWeek" yaml:"dayOfWeek" validate:"min=0,max=7"`
}
