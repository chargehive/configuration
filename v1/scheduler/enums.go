package scheduler

// ConnectorSelector indicates the method used to select a connector
type ConnectorSelector string

const (
	// ConnectorSelectorNone indicates to use no connectors
	// (this is the same as setting empty and is the default value)
	ConnectorSelectorNone ConnectorSelector = "none"

	// ConnectorSelectorStickyFirst indicates the connector should stick to the first connector
	// that was successful for that payment method
	ConnectorSelectorStickyFirst ConnectorSelector = "sticky-first"

	// ConnectorSelectorStickyLast indicates the connector should stick to the most recent connector
	// that was successful for that payment method
	ConnectorSelectorStickyLast ConnectorSelector = "sticky-last"

	// ConnectorSelectorStickyAny indicates to use any connector that has a past success for that payment method
	ConnectorSelectorStickyAny ConnectorSelector = "sticky-any"

	// ConnectorSelectorStickyVerified indicates to use any connector that has a past success and has been verified for that payment method
	ConnectorSelectorStickyVerified ConnectorSelector = "sticky-verified"

	// ConnectorSelectorConfig indicates to use only the connectors specified in the configuration
	ConnectorSelectorConfig ConnectorSelector = "config"
)

// InitiatorType indicates the usage type for an initiator configuration
type InitiatorType string

const (
	// InitiatorTypeAuth indicates a configuration for how an authorization request is handled
	InitiatorTypeAuth InitiatorType = "auth"

	// InitiatorTypeRenewal indicates a configuration for how a renewal request is handled
	InitiatorTypeRenewal InitiatorType = "renewal"

	// InitiatorTypeCapture indicates a configuration for how a secondary capture request is handled
	InitiatorTypeCapture InitiatorType = "capture"
)

type AttemptType string

const (
	// AttemptTypeAuth indicates an authorization attempt
	AttemptTypeAuth AttemptType = "auth"

	// AttemptTypeCapture indicates a capture attempt
	AttemptTypeCapture AttemptType = "capture"
)

// Type indicates the type of a scheduler
type Type string

const (
	// TypeAttempt indicates that this scheduler is to be used for a transaction attempt
	TypeAttempt Type = "attempt"

	// TypeOnDemand indicates this scheduler is to be run on demand (i.e golden day)
	TypeOnDemand Type = "ondemand"
)

// PoolType indicates the order that the items should be iterated within a pool
type PoolType string

const (
	// PoolTypeSingle provides a pool of a single connector
	PoolTypeSingle PoolType = "single"

	// PoolTypeFailover processes the pool items in order until retrieving a result
	PoolTypeFailover PoolType = "failover"

	// PoolTypeCascade iterate connectors according to cascade rules
	PoolTypeCascade PoolType = "cascade"
)

// AttemptOriginType indicates when a given time is based from
type AttemptOriginType string

const (
	// AttemptOriginTypeInitialisation indicates that the time is based from the initialisation of the charge
	AttemptOriginTypeInitialisation AttemptOriginType = "initialisation"

	// AttemptOriginTypeLastFailure indicates that the time is based from the last transaction failure
	AttemptOriginTypeLastFailure AttemptOriginType = "last-failure"
)

// MethodSelector is used to indicate the payment method that should be used
type MethodSelector string

const (
	// MethodSelectorPrimaryMethod indicates that the first available payment method should be used
	MethodSelectorPrimaryMethod MethodSelector = "primary"

	// MethodSelectorBackupMethod indicates that the second available payment method should be used
	MethodSelectorBackupMethod MethodSelector = "backup"

	// MethodSelectorAllMethods indicates that all methods can be used
	MethodSelectorAllMethods MethodSelector = "all"

	// MethodSelectorAllBackupMethods indicates that anything available other than the 1st (primary) should be used
	MethodSelectorAllBackupMethods MethodSelector = "all-backup"
)

// TimeDelaySync specifies when the transaction should be performed relative to the schedules TimeSync
type TimeDelaySync string

const (
	// TimeDelaySyncNone will ignore the TimeSyncHour value
	TimeDelaySyncNone TimeDelaySync = "None"

	// TimeDelaySyncEarliest will run the transaction at the earliest sync hour relative to TimeSync
	TimeDelaySyncEarliest TimeDelaySync = "Earliest"

	// TimeDelaySyncLatest will run the transaction at the latest sync hour relative to TimeSync
	TimeDelaySyncLatest TimeDelaySync = "Latest"

	// TimeDelaySyncClosest will run the transaction at the closest sync hour relative to TimeSync
	TimeDelaySyncClosest TimeDelaySync = "Closest"
)

// TimeZone represents a three character timezone
type TimeZone string

const (
	// TimeZoneULT Users Local Time
	TimeZoneULT TimeZone = "ULT"
	// TimeZoneUTC  Universal Time Coordinated
	TimeZoneUTC TimeZone = "UTC"
	// TimeZoneCIT Charge Initialisation Time (Midnight = Charge Time)
	TimeZoneCIT TimeZone = "CIT"
)

type DayType string

const (
	// DayTypeNone indicates no specific day type
	DayTypeNone DayType = ""
	// DayTypeWeekday indicates a weekday Monday to Friday
	DayTypeWeekday DayType = "weekday"
	// DayTypeWeekend indicates a weekend day Saturday or Sunday
	DayTypeWeekend DayType = "weekend"
)

type DayShift string

const (
	DayShiftNone DayShift = ""
	// DayShiftForward move the date forward to match the day
	DayShiftForward DayShift = "forward"
	// DayShiftBackward move the date backward to match the day
	DayShiftBackward DayShift = "backward"
	// DayShiftClosest move the date to the closest day
	DayShiftClosest DayShift = "closest"
)

type ConnectorRetryType string

const (
	// ConnectorRetryTypeNone No retry
	ConnectorRetryTypeNone ConnectorRetryType = ""
	// ConnectorRetryTypeTokenToPan Use the token first, cascade to Pan if available
	ConnectorRetryTypeTokenToPan ConnectorRetryType = "token-pan"
	// ConnectorRetryTypePanToToken Use the Pan first, cascade to Token if available
	ConnectorRetryTypePanToToken ConnectorRetryType = "pan-token"
)

type ChargeType string

const (
	ChargeTypeDefault     ChargeType = ""
	ChargeTypeUnscheduled            = "unscheduled"
)
