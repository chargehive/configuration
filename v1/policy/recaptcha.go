package policy

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/object"
)

// KindPolicyRecaptcha is the identifier for a PolicyRecaptcha config
const KindPolicyRecaptcha object.Kind = "PolicyRecaptcha"

// RecaptchaPolicy is the policy ran against a charge to determine its Recaptcha status
type RecaptchaPolicy struct {
	ServerKey      string `json:"serverKey" yaml:"serverKey" validate:"required"`
	BlockThreshold int    `json:"blockThreshold" yaml:"blockThreshold" validate:"required,min=0,max=100"`
}

// GetKind returns the RecaptchaPolicy kind
func (RecaptchaPolicy) GetKind() object.Kind { return KindPolicyRecaptcha }

// GetVersion returns the RecaptchaPolicy version
func (RecaptchaPolicy) GetVersion() string { return "v3" }

// NewRecaptchaDefinition creates a new RecaptchaDefinition
func NewRecaptchaDefinition(d *object.Definition) (*RecaptchaDefinition, error) {
	if _, ok := d.Spec.(*RecaptchaPolicy); ok {
		return &RecaptchaDefinition{def: d}, nil
	}
	return nil, errors.New("invalid Recaptcha policy object")
}

// RecaptchaDefinition is the Recaptcha config object definition
type RecaptchaDefinition struct{ def *object.Definition }

// Definition returns the ChargeExpiryDefinition structure
func (d *RecaptchaDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the RecaptchaDefinition
func (d *RecaptchaDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the RecaptchaPolicy contained within the RecaptchaDefinition
func (d *RecaptchaDefinition) Spec() *RecaptchaPolicy { return d.def.Spec.(*RecaptchaPolicy) }
