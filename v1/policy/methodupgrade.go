package policy

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

// KindPolicyMethodUpgrade is the identifier for a PolicyFraud config
const KindPolicyMethodUpgrade object.Kind = "PolicyMethodUpgrade"

// MethodUpgradePolicy is used to temporarily alter the existing payment method information
type MethodUpgradePolicy struct {
	// ExtendExpiry date on payment methods to the next likely expiry date
	ExtendExpiry bool
}

// GetKind returns the MethodUpgradePolicy kind
func (MethodUpgradePolicy) GetKind() object.Kind { return KindPolicyMethodUpgrade }

// GetVersion returns the MethodUpgradePolicy version
func (MethodUpgradePolicy) GetVersion() string { return "v1" }

// NewMethodUpgradeDefinition creates a new MethodUpgradeDefinition
func NewMethodUpgradeDefinition(d *object.Definition) (*MethodUpgradeDefinition, error) {
	if _, ok := d.Spec.(*MethodUpgradePolicy); ok {
		return &MethodUpgradeDefinition{def: d}, nil
	}
	return nil, errors.New("invalid method upgrade policy object")
}

// MethodUpgradeDefinition is the Fraud config object definition
type MethodUpgradeDefinition struct{ def *object.Definition }

// Definition returns the ChargeExpiryDefinition structure
func (d *MethodUpgradeDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the FraudDefinition
func (d *MethodUpgradeDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the MethodUpgradePolicy contained within the MethodUpgradeDefinition
func (d *MethodUpgradeDefinition) Spec() *MethodUpgradePolicy {
	return d.def.Spec.(*MethodUpgradePolicy)
}
