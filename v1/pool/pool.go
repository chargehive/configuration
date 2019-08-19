package pool

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

type ConnectorPool struct {
	Restriction Restriction         `json:"restriction,omitempty" yaml:"restriction,omitempty"`
	Connectors  []ConnectorPoolItem `json:"connectors,omitempty" yaml:"connectors,omitempty"`
}

func (ConnectorPool) GetKind() object.Kind { return KindConnectorPool }
func (ConnectorPool) GetVersion() string   { return "v1" }

func NewConnectorPoolInstance(i *object.Instance) (*ConnectorPoolInstance, error) {
	if _, ok := i.Spec.(*ConnectorPool); ok {
		return &ConnectorPoolInstance{i: i}, nil
	}
	return nil, errors.New("invalid connector pool object")
}

type ConnectorPoolInstance struct{ i *object.Instance }

func (i *ConnectorPoolInstance) MarshalJSON() ([]byte, error) { return json.Marshal(i.i) }
func (i *ConnectorPoolInstance) ConnectorPool() *ConnectorPool {
	return i.i.Spec.(*ConnectorPool)
}

type ConnectorPoolItem struct {
	ConnectorID string `json:"connectorId,omitempty" yaml:"connectorId,omitempty"`
	Priority    int32  `json:"priority,omitempty" yaml:"priority,omitempty"`   // for specific ordering
	Weighting   int32  `json:"weighting,omitempty" yaml:"weighting,omitempty"` // 0 - 1000
	Uses        int32  `json:"uses,omitempty" yaml:"uses,omitempty"`
}
