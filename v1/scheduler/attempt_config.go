package scheduler

import "time"

type AttemptConfig struct {
	PoolType                 PoolType
	MethodSelector           MethodSelector
	ConnectorLimit           int32
	MethodLimit              int
	CascadeDelay             *time.Duration
	OverridePoolConnectorIDs []string // If defined, will override the pool
}
