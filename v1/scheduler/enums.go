package scheduler

type ConnectorSelector string

const (
	ConnectorSelectorStickyFirst    ConnectorSelector = "sticky-first"
	ConnectorSelectorStickyLast     ConnectorSelector = "sticky-last"
	ConnectorSelectorStickyAny      ConnectorSelector = "sticky-any"
	ConnectorSelectorStickyVerified ConnectorSelector = "sticky-verified"
	ConnectorSelectorConfig         ConnectorSelector = "config"
	ConnectorSelectorNone           ConnectorSelector = "none"
)

type InitiatorType string

const (
	InitiatorTypeAuth    InitiatorType = "auth"
	InitiatorTypeRenewal InitiatorType = "renewal"
)

type Type string

const (
	TypeOnDemand Type = "ondemand"
	TypeAttempt  Type = "attempt"
)

type PoolType string

const (
	PoolTypeSingle   PoolType = "single"
	PoolTypeFailover PoolType = "failover"
	PoolTypeCascade  PoolType = "cascade"
)

type AttemptOriginType string

const (
	AttemptOriginTypeInitialisation AttemptOriginType = "initialisation"
	AttemptOriginTypeLastFailure    AttemptOriginType = "last-failure"
)

type MethodSelector string

const (
	MethodSelectorPrimaryMethod MethodSelector = "primary"
	MethodSelectorBackupMethod  MethodSelector = "backup"

	MethodSelectorAllMethods       MethodSelector = "all"
	MethodSelectorAllBackupMethods MethodSelector = "all-backup"
)

type TimeDelaySync string

const (
	TimeDelaySyncEarliest TimeDelaySync = "Earliest"
	TimeDelaySyncLatest   TimeDelaySync = "Latest"
	TimeDelaySyncClosest  TimeDelaySync = "Closest"
)
