package policy

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

// KindPolicyCharge is the identifier for a PolicyCharge config
const KindPolicyCharge object.Kind = "PolicyCharge"

// Restriction is the type of API involvement
type Restriction string

const (
	// RestrictionNone requires no api involvement
	RestrictionNone Restriction = ""

	// RestrictionAPIInitiate requires the api to initiate the request
	RestrictionAPIInitiate Restriction = "api-initiate"

	// RestrictionAPIVerify requires the api to be involved before proceeding
	RestrictionAPIVerify Restriction = "api-verify"

	// RestrictionBlock should block the request
	RestrictionBlock Restriction = "block"
)

// ChargePolicy defines the constraints that when exceeded a charge will not expire
type ChargePolicy struct {
	// MaxAuthAttempts is the number of auths that can be processed, 0 for unlimited
	MaxAuthAttempts int64 `json:"authAttempts" yaml:"authAttempts" validate:"min=0"`

	// AmountUpperLimit is the maximum amount that can be processed
	AmountUpperLimit int64 `json:"amountUpperLimit" yaml:"amountUpperLimit" validate:"min=0"`

	// AmountLowerLimit is the minimum amount that can be processed
	AmountLowerLimit int64 `json:"amountLowerLimit" yaml:"amountLowerLimit" validate:"min=0"`

	// AllowAmounts is a list of amounts that are allowed to be processed
	AllowAmounts []int64 `json:"allowAmounts" yaml:"allowAmounts"`

	// OnCreation is the restriction on charge creation
	OnCreation Restriction `json:"onCreation" yaml:"onCreation" validate:"omitempty,oneof='' api-initiate api-verify block"`

	// OnAuth is the restriction on charge auth
	OnAuth Restriction `json:"onAuth" yaml:"onAuth" validate:"omitempty,oneof='' api-initiate api-verify block"`

	// OnCapture is the restriction on charge capture
	OnCapture Restriction `json:"onCapture" yaml:"onCapture" validate:"omitempty,oneof='' api-initiate api-verify block"`

	// OnRefund is the restriction on charge refund
	OnRefund Restriction `json:"onRefund" yaml:"onRefund" validate:"omitempty,oneof='' api-initiate api-verify block"`
}

// GetKind returns the ChargePolicy kind
func (ChargePolicy) GetKind() object.Kind { return KindPolicyCharge }

// GetVersion returns the ChargePolicy version
func (ChargePolicy) GetVersion() string { return "v1" }

// ChargePolicyDefinition is the Charge config object definition
type ChargePolicyDefinition struct{ def *object.Definition }

// NewChargePolicyDefinition creates a new ChargePolicyDefinition
func NewChargePolicyDefinition(d *object.Definition) (*ChargePolicyDefinition, error) {
	if _, ok := d.Spec.(*ChargePolicy); ok {
		return &ChargePolicyDefinition{def: d}, nil
	}
	return nil, errors.New("invalid Charge policy object")
}

// Definition returns the ChargePolicyDefinition structure
func (d *ChargePolicyDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the ChargePolicyDefinition
func (d *ChargePolicyDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the CascadePolicy contained within the ChargePolicyDefinition
func (d *ChargePolicyDefinition) Spec() *ChargePolicy {
	return d.def.Spec.(*ChargePolicy)
}
