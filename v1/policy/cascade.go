package policy

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/connectorconfig"
	"github.com/chargehive/configuration/object"
)

const KindPolicyCascade object.Kind = "PolicyCascade"

// CascadePolicy - Rules for connector casscading
type CascadePolicy struct {
	Rules []CascadeRule
}

type CascadeRule struct {
	Library              connectorconfig.Library
	OriginalResponseCode string
	Cascade              bool
}

func (CascadePolicy) GetKind() object.Kind { return KindPolicyCascade }
func (CascadePolicy) GetVersion() string   { return "v1" }

type CascadePolicyDefinition struct{ def *object.Definition }

func NewCascadePolicyDefinition(d *object.Definition) (*CascadePolicyDefinition, error) {
	if _, ok := d.Spec.(*CascadePolicy); ok {
		return &CascadePolicyDefinition{def: d}, nil
	}
	return nil, errors.New("invalid Cascade policy object")
}
func (d *CascadePolicyDefinition) Definition() *object.Definition { return d.def }
func (d *CascadePolicyDefinition) MarshalJSON() ([]byte, error)   { return json.Marshal(d.def) }
func (d *CascadePolicyDefinition) Spec() *CascadePolicy           { return d.def.Spec.(*CascadePolicy) }
