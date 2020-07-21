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

	// TimeSyncZone indicates the timezone that the TimeSyncHour is relative to. Ignored if TimeDelaySync is set to None
	TimeSyncZone TimeZone `json:"timeSyncZone" yaml:"timeSyncZone" validate:"oneof=ULT UTC"`
}
