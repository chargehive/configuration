package policy

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

const KindPolicyMethodUpgrade object.Kind = "PolicyMethodUpgrade"

type MethodUpgradePolicy struct {
	ExtendExpiry bool // Extend expiry date on payment methods to the next likely expiry date
}

func (MethodUpgradePolicy) GetKind() object.Kind { return KindPolicyMethodUpgrade }
func (MethodUpgradePolicy) GetVersion() string   { return "v1" }

func NewMethodUpgradeDefinition(d *object.Definition) (*MethodUpgradeDefinition, error) {
	if _, ok := d.Spec.(*MethodUpgradePolicy); ok {
		return &MethodUpgradeDefinition{def: d}, nil
	}
	return nil, errors.New("invalid method upgrade policy object")
}

type MethodUpgradeDefinition struct{ def *object.Definition }

func (d *MethodUpgradeDefinition) Definition() *object.Definition { return d.def }
func (d *MethodUpgradeDefinition) MarshalJSON() ([]byte, error)   { return json.Marshal(d.def) }
func (d *MethodUpgradeDefinition) Spec() *MethodUpgradePolicy {
	return d.def.Spec.(*MethodUpgradePolicy)
}
