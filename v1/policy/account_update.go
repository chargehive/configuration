package policy

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/object"
)

// KindPolicyAccountUpdate is the identifier for a PolicyAccountUpdate config
const KindPolicyAccountUpdate object.Kind = "PolicyAccountUpdate"

// AccountUpdatePolicy are the properties used to update a payment account / payment method
type AccountUpdatePolicy struct {
	// ConnectorIDs is the ID of the account update connector
	ConnectorID string `json:"connectorID" yaml:"connectorIDs" validate:"min=1"`
}

// GetKind returns the AccountUpdatePolicy Kind
func (AccountUpdatePolicy) GetKind() object.Kind { return KindPolicyAccountUpdate }

// GetVersion returns the AccountUpdatePolicy version
func (AccountUpdatePolicy) GetVersion() string { return "v1" }

// AccountUpdatePolicyDefinition is the AccountUpdatePolicy config object definition
type AccountUpdatePolicyDefinition struct{ def *object.Definition }

// NewAccountUpdatePolicyDefinition creates a new AccountUpdatePolicyDefinition
func NewAccountUpdatePolicyDefinition(d *object.Definition) (*AccountUpdatePolicyDefinition, error) {
	if _, ok := d.Spec.(*CascadePolicy); ok {
		return &AccountUpdatePolicyDefinition{def: d}, nil
	}
	return nil, errors.New("invalid Cascade policy object")
}

// Definition returns the AccountUpdatePolicyDefinition structure
func (d *AccountUpdatePolicyDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the AccountUpdatePolicyDefinition
func (d *AccountUpdatePolicyDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the AccountUpdatePolicy contained within the AccountUpdatePolicyDefinition
func (d *AccountUpdatePolicyDefinition) Spec() *AccountUpdatePolicy {
	return d.def.Spec.(*AccountUpdatePolicy)
}
