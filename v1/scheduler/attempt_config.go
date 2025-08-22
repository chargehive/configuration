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

	// PreferNetworkToken indicates if the network token should be used instead of the PAN
	// Deprecated - use PreferredTokens
	PreferNetworkToken bool `json:"preferNetworkToken,omitempty" yaml:"preferNetworkToken,omitempty"`

	PreferredTokens []TokenSource `json:"preferredTokens,omitempty" yaml:"preferredTokens,omitempty" validate:"dive,uppercase,oneof=TS_PAN TS_CONNECTOR TS_NETWORK_TOKEN TS_GOOGLE_PAY TS_APPLE_PAY TS_SAMSUNG_PAY TS_AMAZON_PAY TS_REVOLUT_PAY TS_WECHAT_PAY TS_ALIPAY TS_PAYPAL TS_TEXT_TO_PAY TS_BACS TS_SEPA TS_ACH TS_BANK_TRANSFER TS_BANCONTACT TS_EPS TS_IDEAL TS_PRZELEWY24 TS_TWINT TS_SOFORT TS_MULTIBANCO TS_KLARNA TS_AFTER_PAY TS_COINBASE TS_GOOGLE_PLAY TS_APPLE_STORE"`

	// MethodRefresh indicates if the method should be refreshed before processing
	MethodRefresh bool `json:"methodRefresh,omitempty" yaml:"methodRefresh,omitempty"`

	ShouldTokenize bool `json:"shouldTokenize,omitempty" yaml:"shouldTokenize,omitempty"`

	RecoveryAgentConnectorID string `json:"recoveryAgentConnectorID" yaml:"recoveryAgentConnectorID"`

	ChargeType ChargeType `json:"chargeType" yaml:"chargeType" validate:"oneof='' unscheduled"`

	// AmountPercentage is the percentage of the full charge amount to attempt e.g. 20% of $100, would attempt $20
	AmountPercentage int32 `json:"amountPercentage" yaml:"amountPercentage" validate:"omitempty,min=0,max=100"`

	// MaxAmount is the maximum amount that can be handled in a single attempt, in minor units
	MaxAmount int64 `json:"maxAmount" yaml:"maxAmount" validate:"omitempty,min=0,max=100000"`

	// HandleUnderPayment indicates how to handle under payments
	HandleUnderPayment UnderPaymentHandler `json:"handleUnderPayment" yaml:"handleUnderPayment" validate:"omitempty,oneof=credit outstanding lock"`

	// UnderPaymentLockSeconds is the amount of seconds to lock the charge for if a lock underpayment handler is used
	UnderPaymentLockSeconds int64 `json:"underPaymentLockSeconds" yaml:"underPaymentLockSeconds" validate:"omitempty,min=0,max=100"`
}
