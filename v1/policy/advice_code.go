package policy

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/object"
)

// KindPolicyAdviceCode is the identifier for an advice codes policy config
const KindPolicyAdviceCode object.Kind = "PolicyAdviceCode"

// AdviceCodePolicy scales the delay values returned by a connector in response
// to a merchant advice code. Multiplier is applied to both the recommended and
// required delays surfaced on the connector TransactionResponse.
//
// 1.0 = original delay, 0.5 = half delay, 0.0 = no delay.
type AdviceCodePolicy struct {
	DelayMultiplier float64 `json:"delayMultiplier" yaml:"delayMultiplier" validate:"gte=0,lte=1"`
}

// GetKind returns the AdviceCodePolicy kind
func (AdviceCodePolicy) GetKind() object.Kind { return KindPolicyAdviceCode }

// GetVersion returns the AdviceCodePolicy version
func (AdviceCodePolicy) GetVersion() string { return "v1" }

// AdviceCodePolicyDefinition is the AdviceCodePolicy config object definition
type AdviceCodePolicyDefinition struct{ def *object.Definition }

// NewAdviceCodePolicyDefinition creates a new AdviceCodePolicyDefinition
func NewAdviceCodePolicyDefinition(d *object.Definition) (*AdviceCodePolicyDefinition, error) {
	if _, ok := d.Spec.(*AdviceCodePolicy); ok {
		return &AdviceCodePolicyDefinition{def: d}, nil
	}
	return nil, errors.New("invalid advice codes policy object")
}

// Definition returns the AdviceCodePolicyDefinition structure
func (d *AdviceCodePolicyDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the AdviceCodePolicyDefinition
func (d *AdviceCodePolicyDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the AdviceCodePolicy contained within the AdviceCodePolicyDefinition
func (d *AdviceCodePolicyDefinition) Spec() *AdviceCodePolicy {
	return d.def.Spec.(*AdviceCodePolicy)
}
