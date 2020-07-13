package connector

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

// KindConnectorPool kind indicates that the configuration is a collection of selectable connectors
const KindConnectorPool object.Kind = "ConnectorPool"

// Restriction is the method used to determine a connector selected from a pool
type Restriction string

const (
	// RestrictionUnrestricted no restrictions on selection, an empty string indicates this default value
	RestrictionUnrestricted Restriction = "unrestricted"

	// RestrictionNoRepeat indicates the pool should not select the last processed connector on this charge
	RestrictionNoRepeat Restriction = "noRepeat"

	// RestrictionLowestUsage the pool will prioritise the connector with the lowest amount of attempts across all charges
	RestrictionLowestUsage Restriction = "lowestUsage"
)

// Pool is used to select a group of connectors and the order that they should be used in
type Pool struct {
	Restriction Restriction `json:"restriction" yaml:"restriction" validate:"oneof=unrestricted noRepeat lowestUsage"`
	Connectors  []PoolItem  `json:"connectors" yaml:"connectors" validate:"gt=0,dive"`
}

// GetKind returns the pool kind
func (Pool) GetKind() object.Kind { return KindConnectorPool }

// GetVersion returns the pool version
func (Pool) GetVersion() string { return "v1" }

// NewPoolDefinition returns a new pool definition
func NewPoolDefinition(d *object.Definition) (*PoolDefinition, error) {
	if _, ok := d.Spec.(*Pool); ok {
		return &PoolDefinition{def: d}, nil
	}
	return nil, errors.New("invalid connector pool object")
}

// PoolDefinition defines the structure of a definition
type PoolDefinition struct{ def *object.Definition }

// Definition returns the defintion for a poolDefinition
func (d *PoolDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns JSON value for a poolDefinition
func (d *PoolDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the pool specification from a PoolDefinition
func (d *PoolDefinition) Spec() *Pool { return d.def.Spec.(*Pool) }

// PoolItem is a single entry into a pool used to determine the connector that should be used
type PoolItem struct {
	// ConnectorID is the identifier for a connector
	ConnectorID string `json:"connectorId" yaml:"connectorId" validate:"required,lowercase"`

	// Priority is a integer value where the item with the highest priority has lowest value
	// (zero is the highest priority number possible)
	Priority int32 `json:"priority" yaml:"priority" validate:"min=0"` // for specific ordering

	// Weighting is used to weigh items of the same priority, secondary to priority
	Weighting int32 `json:"weighting" yaml:"weighting" validate:"min=0,max=1000"` // 0 - 1000

	// Uses is the maximum times a connector can be used in a single charge
	Uses int32 `json:"uses" yaml:"uses" validate:"min=0"`
}
