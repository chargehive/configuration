package policy

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/object"
)

// KindPolicyRateLimit is the identifier for a RateLimitPolicy config
const KindPolicyRateLimit object.Kind = "RateLimitPolicy"

type RateLimitKey string

const (
	RateLimitKeyChargePlacementID RateLimitKey = "PlacementID"
	RateLimitKeyMerchantReference RateLimitKey = "MerchantReference"
	RateLimitKeyCorrelationID     RateLimitKey = "CorrelationID"
	RateLimitKeyBillingProfileID  RateLimitKey = "BillingProfileID"
	RateLimitKeyIP                RateLimitKey = "IP"
	RateLimitKeyUserAgent         RateLimitKey = "UserAgent"
	RateLimitKeyPlatform          RateLimitKey = "Platform"
)

type RateLimitPolicy struct {
	// LimitProperty is the property to limit the rate by, global if empty
	LimitProperty RateLimitKey `json:"limitProperty" yaml:"limitProperty" validate:"oneof=PlacementID MerchantReference CorrelationID BillingProfileID IP UserAgent Platform"`
	// HardLimit is the maximum number of requests allowed in the window
	HardLimit int32 `json:"hardLimit" yaml:"hardLimit" validate:"required,min=1"`
	// Window is the time window in minutes that the limit is applied to
	Window int32 `json:"window" yaml:"window" validate:"required,min=1"`
}

func (r RateLimitPolicy) GetKind() object.Kind { return KindPolicyRateLimit }

func (r RateLimitPolicy) GetVersion() string { return "v1" }

type RateLimitDefinition struct{ def *object.Definition }

// NewRateLimitDefinition creates a new RateLimitDefinition
func NewRateLimitDefinition(d *object.Definition) (*RateLimitDefinition, error) {
	if _, ok := d.Spec.(*RateLimitPolicy); ok {
		return &RateLimitDefinition{def: d}, nil
	}
	return nil, errors.New("invalid Rate Limit Policy object")
}

// Definition returns the RateLimitDefinition structure
func (d *RateLimitDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the RateLimitDefinition
func (d *RateLimitDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the RateLimitPolicy contained within the RateLimitDefinition
func (d *RateLimitDefinition) Spec() *RateLimitPolicy { return d.def.Spec.(*RateLimitPolicy) }
