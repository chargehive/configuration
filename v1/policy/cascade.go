package policy

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/connectorconfig"
	"github.com/chargehive/configuration/object"
)

// KindPolicyCascade is the identifier for a PolicyCascade config
const KindPolicyCascade object.Kind = "PolicyCascade"

// CascadePolicy is a collection of rules for connector cascading
type CascadePolicy struct {
	Rules []CascadeRule `json:"rules" yaml:"rules" validate:"min=1"`
}

// CascadeRule is a single cascade rule (all fields are required)
type CascadeRule struct {
	// Library designates the library that this cascade rule is applied to
	Library connectorconfig.Library `json:"library" yaml:"library" validate:"connector-library"`

	// OriginalResponseCode is the raw error code returned by the library to be matched
	OriginalResponseCode string `json:"originalResponseCode" yaml:"originalResponseCode" validate:"required"`

	// Cascade determines if this rule results in a cascade or not
	Cascade *bool `json:"cascade" yaml:"cascade" validate:"required"`
}

func (c CascadeRule) GetCascade() bool {
	if c.Cascade == nil {
		return true
	}
	return *c.Cascade
}

// GetKind returns the CascadePolicy Kind
func (CascadePolicy) GetKind() object.Kind { return KindPolicyCascade }

// GetVersion returns the CascadePolicy version
func (CascadePolicy) GetVersion() string { return "v1" }

// CascadePolicyDefinition is the CascadePolicy config object definition
type CascadePolicyDefinition struct{ def *object.Definition }

// NewCascadePolicyDefinition creates a new CascadePolicyDefinition
func NewCascadePolicyDefinition(d *object.Definition) (*CascadePolicyDefinition, error) {
	if _, ok := d.Spec.(*CascadePolicy); ok {
		return &CascadePolicyDefinition{def: d}, nil
	}
	return nil, errors.New("invalid Cascade policy object")
}

// Definition returns the CascadePolicyDefinition structure
func (d *CascadePolicyDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the CascadePolicyDefinition
func (d *CascadePolicyDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the CascadePolicy contained within the CascadePolicyDefinition
func (d *CascadePolicyDefinition) Spec() *CascadePolicy { return d.def.Spec.(*CascadePolicy) }
