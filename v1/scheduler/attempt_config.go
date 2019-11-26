package scheduler

import "time"

// AttemptConfig defines a single attempt
type AttemptConfig struct {
	// PoolType indicates the order that this attempt should iterate connectors
	PoolType PoolType

	// MethodSelector indicates how payment method should be selected for this attempt
	MethodSelector MethodSelector

	// ConnectorLimit is a maximum number of connectors to process within an attempt per method
	ConnectorLimit int32

	// ConnectorLimit is a maximum number of methods to be attempt per method
	MethodLimit int

	// CascadeDelay is the duration to wait between each cascade
	CascadeDelay *time.Duration

	// OverridePoolConnectorIDs will use this connectors instead of the ones in the pool
	OverridePoolConnectorIDs []string
}
