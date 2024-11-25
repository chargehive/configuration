package scheduler

import "time"

// AttemptConfig defines a single attempt
type AttemptConfig struct {
	// PoolType indicates the order that this attempt should iterate connectors
	PoolType PoolType `json:"poolType" yaml:"poolType" validate:"oneof=single failover cascade"`

	// MethodSelector indicates how payment method should be selected for this attempt
	MethodSelector MethodSelector `json:"methodSelector" yaml:"methodSelector" validate:"oneof=primary backup all all-backup"`

	// ConnectorRetryType indicates how the attempt should retry (on the same connector), if needed
	ConnectorRetryType ConnectorRetryType `json:"connectorRetryType" yaml:"connectorRetryType" validate:"oneof='' token-pan pan-token"`

	// ConnectorLimit is a maximum number of connectors to process within an attempt per method
	ConnectorLimit int32 `json:"connectorLimit" yaml:"connectorLimit" validate:"min=0,max=1000"`

	// MethodLimit is a maximum number of methods to be attempt per method
	MethodLimit int `json:"methodLimit" yaml:"methodLimit" validate:"min=0"`

	// CascadeDelay is the duration to wait between each cascade
	CascadeDelay *time.Duration `json:"cascadeDelay" yaml:"cascadeDelay" validate:"required,gte=0"`

	// AttemptType indicates what type of transaction to submit to the connector
	AttemptType AttemptType `json:"attemptType,omitempty" yaml:"attemptType,omitempty" validate:"omitempty,oneof=capture auth"`

	// OverridePoolConnectorIDs will use this connectors instead of the ones in the pool
	OverridePoolConnectorIDs []string `json:"overridePoolConnectorIDs,omitempty" yaml:"overridePoolConnectorIDs,omitempty" validate:"dive,lowercase"`

	Prefer3RI bool `json:"prefer3RI,omitempty" yaml:"prefer3RI,omitempty"`

	PreferNetworkToken bool `json:"preferNetworkToken,omitempty" yaml:"preferNetworkToken,omitempty"`

	ShouldTokenize bool `json:"shouldTokenize,omitempty" yaml:"shouldTokenize,omitempty"`

	RecoveryAgentConnectorID string `json:"recoveryAgentConnectorID" yaml:"recoveryAgentConnectorID"`

	ChargeType ChargeType `json:"chargeType" yaml:"chargeType" validate:"oneof='' unscheduled"`
}
