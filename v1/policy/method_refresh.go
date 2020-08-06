package policy

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/object"
)

// KindPolicyMethodRefresh is the identifier for a PolicyMethodRefresh config
const KindPolicyMethodRefresh object.Kind = "PolicyMethodRefresh"

// MethodRefreshPolicy are the properties used to update a payment account / payment method
type MethodRefreshPolicy struct {
	// ConnectorIDs is the ID of the account update connector
	ConnectorID string `json:"connectorID" yaml:"connectorIDs" validate:"min=1"`
}

// GetKind returns the MethodRefreshPolicy Kind
func (MethodRefreshPolicy) GetKind() object.Kind { return KindPolicyMethodRefresh }

// GetVersion returns the MethodRefreshPolicy version
func (MethodRefreshPolicy) GetVersion() string { return "v1" }

// MethodRefreshPolicyDefinition is the MethodRefreshPolicy config object definition
type MethodRefreshPolicyDefinition struct{ def *object.Definition }

// NewMethodRefreshPolicyDefinition creates a new MethodRefreshPolicyDefinition
func NewMethodRefreshPolicyDefinition(d *object.Definition) (*MethodRefreshPolicyDefinition, error) {
	if _, ok := d.Spec.(*CascadePolicy); ok {
		return &MethodRefreshPolicyDefinition{def: d}, nil
	}
	return nil, errors.New("invalid Cascade policy object")
}

// Definition returns the MethodRefreshPolicyDefinition structure
func (d *MethodRefreshPolicyDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the MethodRefreshPolicyDefinition
func (d *MethodRefreshPolicyDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the MethodRefreshPolicy contained within the MethodRefreshPolicyDefinition
func (d *MethodRefreshPolicyDefinition) Spec() *MethodRefreshPolicy {
	return d.def.Spec.(*MethodRefreshPolicy)
}
