package scheduler

import "time"

// Schedule is a single schedule item for a single attempt
type Schedule struct {
	// AttemptConfig is the configuration to use when processing this schedule
	AttemptConfig AttemptConfig `json:"attemptConfig" yaml:"attemptConfig" validate:"required,dive"`

	// TimeDelay is the amount of time to wait before processing after TimeDelayOrigin
	TimeDelay time.Duration `json:"timeDelay" yaml:"timeDelay" validate:"required,min=0,max=1000"`

	// TimeDelayOrigin specifies when the time origin is based from
	TimeDelayOrigin AttemptOriginType `json:"timeDelayOrigin" yaml:"timeDelayOrigin" validate:"required,oneof=initialisation last-failure"`

	// TimeDelaySync indicates the direction in time that the schedule should move when syncing to TimeSyncHour
	TimeDelaySync TimeDelaySync `json:"timeDelaySync" yaml:"timeDelaySync" validate:"required,oneof=Earliest Latest Closest"`

	// TimeSyncHour is an hour designation (1-24) i.e 2 == 2AM | where less than 1 indicates that this value is not set
	TimeSyncHour int `json:"timeSyncHour" yaml:"timeSyncHour" validate:"required,min=0,max=24"`

	// TimeSyncZone indicates the timezone that the TimeSyncHour is relative to
	TimeSyncZone TimeZone `json:"timeSyncZone" yaml:"timeSyncZone" validate:"required,oneof=ULT UTC"`
}

// AttemptConfig defines a single attempt
type AttemptConfig struct {
	// PoolType indicates the order that this attempt should iterate connectors
	PoolType PoolType `json:"poolType" yaml:"poolType" validate:"required,oneof=single failover cascade"`

	// MethodSelector indicates how payment method should be selected for this attempt
	MethodSelector MethodSelector `json:"methodSelector" yaml:"methodSelector" validate:"required,oneof=primary backup all all-backup"`

	// ConnectorLimit is a maximum number of connectors to process within an attempt per method
	ConnectorLimit int32 `json:"connectorLimit" yaml:"connectorLimit" validate:"required,min=0,max=1000"`

	// MethodLimit is a maximum number of methods to be attempt per method
	MethodLimit int `json:"methodLimit" yaml:"methodLimit" validate:"required,min=0,max=1000"`

	// CascadeDelay is the duration to wait between each cascade
	CascadeDelay *time.Duration `json:"cascadeDelay" yaml:"cascadeDelay" validate:"min=0"`

	// OverridePoolConnectorIDs will use this connectors instead of the ones in the pool
	OverridePoolConnectorIDs []string `json:"overridePoolConnectorIDs" yaml:"overridePoolConnectorIDs"`
}

// ScheduleRefund is a single schedule item for a single refund attempt
type ScheduleRefund struct {
	// TimeDelay is the amount of time to wait since the last refund attempt
	TimeDelay time.Duration `json:"timeDelay" yaml:"timeDelay" validate:"required,min=0"`
}
