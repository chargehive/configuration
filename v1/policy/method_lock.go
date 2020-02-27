package policy

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/object"
)

// KindPolicyMethodLock is the identifier for a method lock policy config
const KindPolicyMethodLock object.Kind = "PolicyMethodLock"

// MethodLockPolicy is used to lock a payment method preventing it being used for payment
type MethodLockPolicy struct {
	// Duration is the duration of time (in seconds) that a payment method should be locked for on application of this policy
	Duration int64 `json:"duration" yaml:"duration" validate:"min=0"`

	// Reason is a description of why a card locked due to this policy was locked.
	Reason string `json:"reason" yaml:"reason" validate:"-"`
}

// GetKind returns the MethodLockPolicy kind
func (MethodLockPolicy) GetKind() object.Kind { return KindPolicyMethodLock }

// GetVersion returns the MethodLockPolicy version
func (MethodLockPolicy) GetVersion() string { return "v1" }

// MethodLockDefinition is the Lock config object definition
type MethodLockDefinition struct{ def *object.Definition }

// NewMethodLockDefinition creates a new MethodLockDefinition
func NewMethodLockDefinition(d *object.Definition) (*MethodLockDefinition, error) {
	if _, ok := d.Spec.(*MethodLockPolicy); ok {
		return &MethodLockDefinition{def: d}, nil
	}
	return nil, errors.New("invalid method lock policy object")
}

// Definition returns the MethodLockDefinition structure
func (d *MethodLockDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the MethodLockDefinition
func (d *MethodLockDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the MethodLockPolicy contained within the MethodLockDefinition
func (d *MethodLockDefinition) Spec() *MethodLockPolicy {
	return d.def.Spec.(*MethodLockPolicy)
}
