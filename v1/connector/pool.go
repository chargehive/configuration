package connector

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

const KindConnectorPool object.Kind = "ConnectorPool"

type Restriction string

const (
	RestrictionNoRepeat     Restriction = "noRepeat"
	RestrictionLowestUsage  Restriction = "lowestUsage"
	RestrictionUnrestricted Restriction = "unrestricted"
)

type Pool struct {
	Restriction Restriction `json:"restriction,omitempty" yaml:"restriction,omitempty"`
	Connectors  []PoolItem  `json:"connectors,omitempty" yaml:"connectors,omitempty"`
}

func (Pool) GetKind() object.Kind { return KindConnectorPool }
func (Pool) GetVersion() string   { return "v1" }

func NewPoolDefinition(d *object.Definition) (*PoolDefinition, error) {
	if _, ok := d.Spec.(*Pool); ok {
		return &PoolDefinition{def: d}, nil
	}
	return nil, errors.New("invalid connector pool object")
}

type PoolDefinition struct{ def *object.Definition }

func (d *PoolDefinition) Definition() *object.Definition { return d.def }
func (d *PoolDefinition) MarshalJSON() ([]byte, error)   { return json.Marshal(d.def) }
func (d *PoolDefinition) Spec() *Pool                    { return d.def.Spec.(*Pool) }

type PoolItem struct {
	ConnectorID string `json:"connectorId,omitempty" yaml:"connectorId,omitempty"`
	Priority    int32  `json:"priority,omitempty" yaml:"priority,omitempty"`   // for specific ordering
	Weighting   int32  `json:"weighting,omitempty" yaml:"weighting,omitempty"` // 0 - 1000
	Uses        int32  `json:"uses,omitempty" yaml:"uses,omitempty"`
}
