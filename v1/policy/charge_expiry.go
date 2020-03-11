package policy

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/chargehive/configuration/object"
)

// KindPolicyChargeExpiry is the identifier for a PolicyChargeExpiry config
const KindPolicyChargeExpiry object.Kind = "PolicyChargeExpiry"

// ChargeExpiryPolicy defines the constraints that when exceeded a charge will not expire
type ChargeExpiryPolicy struct {
	// Timeout is the duration from when the charge was first initialized
	Timeout time.Duration `json:"timeout" yaml:"timeout" validate:"gte=0"`

	// Attempts is the maximum number of attempts that can be performed on a charge
	Attempts int64 `json:"attempts" yaml:"attempts" validate:"min=0"`
}

// GetKind returns the ChargeExpiryPolicy kind
func (ChargeExpiryPolicy) GetKind() object.Kind { return KindPolicyChargeExpiry }

// GetVersion returns the ChargeExpiryPolicy version
func (ChargeExpiryPolicy) GetVersion() string { return "v1" }

// ChargeExpiryDefinition is the ChargeExpiry config object definition
type ChargeExpiryDefinition struct{ def *object.Definition }

// NewChargeExpiryDefinition creates a new ChargeExpiryDefinition
func NewChargeExpiryDefinition(d *object.Definition) (*ChargeExpiryDefinition, error) {
	if _, ok := d.Spec.(*ChargeExpiryPolicy); ok {
		return &ChargeExpiryDefinition{def: d}, nil
	}
	return nil, errors.New("invalid Charge Expiry policy object")
}

// Definition returns the ChargeExpiryDefinition structure
func (d *ChargeExpiryDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the ChargeExpiryDefinition
func (d *ChargeExpiryDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// UnmarshalJSON unmarshals a JSON encoded ChargeExpiryDefinition
func (d *ChargeExpiryDefinition) UnmarshalJSON(data []byte) error {
	type alias ChargeExpiryDefinition
	a := &struct{ *alias }{alias: (*alias)(d)}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	// convert seconds entry to time.Duration (nanoseconds)
	if a.alias.def != nil {
		dspec := a.alias.def.Spec.(*ChargeExpiryPolicy)
		intVal := int64(dspec.Timeout)
		if intVal != 0 && intVal < int64(time.Second) {
			dspec.Timeout = time.Second * time.Duration(intVal)
			a.alias.def.Spec = dspec
		}
	}

	d = (*ChargeExpiryDefinition)(a.alias)
	return nil
}

// Spec returns the CascadePolicy contained within the ChargeExpiryDefinition
func (d *ChargeExpiryDefinition) Spec() *ChargeExpiryPolicy { return d.def.Spec.(*ChargeExpiryPolicy) }
