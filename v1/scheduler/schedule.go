package scheduler

import "time"

// Schedule is a single schedule item for a single attempt
type Schedule struct {
	// AttemptConfig is the configuration to use when processing this schedule
	AttemptConfig AttemptConfig `json:"attemptConfig" yaml:"attemptConfig" validate:"required,dive"`

	// TimeDelay [deprecated - prefer delay days] is the amount of time to wait before processing after TimeDelayOrigin
	TimeDelay time.Duration `json:"timeDelay,omitempty" yaml:"timeDelay" validate:"gte=0"`

	// DAY CONFIG

	// DelayDays is the number of days to delay the schedule by
	DelayDays int `json:"delayDays,omitempty" yaml:"delayDays" validate:"gte=0"`

	// DayOfMonth is the day of the month to process the schedule
	DayOfMonth int `json:"dayOfMonth,omitempty" yaml:"dayOfMonth" validate:"min=0,max=31"`

	// DayOfWeek is the day of the week to process the schedule (1 = Monday) - Sunday moved to 7 to leave 0 as an ignored value
	DayOfWeek int `json:"dayOfWeek,omitempty" yaml:"dayOfWeek" validate:"min=0,max=7"`

	// DayType indicates the type of day to process the schedule
	DayType DayType `json:"dayType,omitempty" yaml:"dayType" validate:"oneof='' weekday weekend"`

	// DayShift indicates the direction to take when syncing to Days
	DayShift DayShift `json:"daySync,omitempty" yaml:"daySync" validate:"oneof='' forward backward"`

	// TIME CONFIG

	// TimeDelayOrigin specifies when the time origin is based from
	TimeDelayOrigin AttemptOriginType `json:"timeDelayOrigin,omitempty" yaml:"timeDelayOrigin" validate:"oneof=initialisation last-failure"`

	// TimeDelaySync indicates the direction in time that the schedule should move when syncing to TimeSyncHour
	TimeDelaySync TimeDelaySync `json:"timeDelaySync,omitempty" yaml:"timeDelaySync" validate:"oneof=None Earliest Latest Closest"`

	// TimeSyncHour is an hour designation (0-23) i.e 2 == 2AM. Ignored if TimeDelaySync is set to None
	TimeSyncHour int `json:"timeSyncHour,omitempty" yaml:"timeSyncHour" validate:"min=0,max=23"`

	// TimeWindowHours is the number of hours grace for a transaction to be processed within.  Allowing for a transaction to be processed within X hours of the requested delay & sync hour e.g. 10am > 11am with a TimeWindowHours of 1
	TimeWindowHours int `json:"timeWindowHours,omitempty" yaml:"timeWindowHours" validate:"gte=0"`

	// TimeSyncZone indicates the timezone that the TimeSyncHour is relative to. Ignored if TimeDelaySync is set to None
	TimeSyncZone TimeZone `json:"timeSyncZone,omitempty" yaml:"timeSyncZone" validate:"oneof=ULT UTC CIT"`
}
