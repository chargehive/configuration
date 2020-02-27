package policy

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/object"
)

// KindPolicyMethodVerify is the identifier for a method verify policy config
const KindPolicyMethodVerify object.Kind = "PolicyMethodVerify"

// MethodVerifyPolicy are the settings used to perform a payment method verification
type MethodVerifyPolicy struct {
	// If true the payment method will be verified at the same time it is tokenized
	// [Optional. Defaults to false]
	VerifyMethodOnTokenization bool `json:"verifyMethodOnTokenization" yaml:"verifyMethodOnTokenization" validate:"-"`

	// Amount is a monetary value integer that will be authorized on a card to verify its ability to make payments
	// this should be an amount in the currencies smallest denomination i.e a value of 44 would equate to 0.44 GBP
	// [Required]
	Amount int64 `json:"amount" yaml:"amount" validate:"min=1"`

	// This is the currency code for the specified amount i.e GBP
	// [Required]
	AmountCurrency string `json:"amountCurrency" yaml:"amountCurrency"validate:"oneof=ARS AUD AZN BHD BYR BOB BAM BRL BGN CAD CLP CNY COP CRC HRK CZK DKK DOP XCD EGP ETB EUR FJD GEL GTQ HTG HNL HKD HUF ISK INR IDR IRR JMD JPY JOD KZT KES KRW KWD LVL LBP LYD LTL MWK MYR MUR MXN MDL MAD ILS NZD NGN NOK OMR PKR PAB PYG PEN PHP PLN GBP QAR RON RUB SAR RSD SGD ZAR LKR SEK CHF SYP TWD THB TTD TND TRY UAH AED UYU USD VEF VND"`

	// This is the ID of the connector that is used to verify payment methods
	// [Required]
	ConnectorID string `json:"connectorID" yaml:"connectorID" validate:"required"`
}

// GetKind returns the MethodVerifyPolicy kind
func (MethodVerifyPolicy) GetKind() object.Kind { return KindPolicyMethodVerify }

// GetVersion returns the MethodVerifyPolicy version
func (MethodVerifyPolicy) GetVersion() string { return "v1" }

// MethodVerifyDefinition is the verification config object definition
type MethodVerifyDefinition struct{ def *object.Definition }

// NewMethodVerifyDefinition creates a new MethodVerifyDefinition
func NewMethodVerifyDefinition(d *object.Definition) (*MethodVerifyDefinition, error) {
	if _, ok := d.Spec.(*MethodVerifyPolicy); ok {
		return &MethodVerifyDefinition{def: d}, nil
	}
	return nil, errors.New("invalid method verify policy object")
}

// Definition returns the MethodVerifyDefinition structure
func (d *MethodVerifyDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the MethodVerifyDefinition
func (d *MethodVerifyDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the MethodVerifyPolicy contained within the MethodVerifyDefinition
func (d *MethodVerifyDefinition) Spec() *MethodVerifyPolicy {
	return d.def.Spec.(*MethodVerifyPolicy)
}
