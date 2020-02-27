package policy

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

// KindPolicySCA is the identifier for a ScaPolicy config
const KindPolicySCA object.Kind = "PolicySCA"

// SCABypassMode indicates the action to be performed when a verification result is returned by a connector
type SCABypassMode string

const (
	// SCABypassModeNone indicates to not bypass a required challenge
	SCABypassModeNone SCABypassMode = ""

	// SCABypassModeCascade indicates to auth on the next connector
	SCABypassModeCascade SCABypassMode = "cascade"

	// SCABypassModeCurrent indicate to bypass, but stay on the current connector (attempt auth anyway)
	SCABypassModeCurrent SCABypassMode = "current"
)

// ScaPolicy options determine how to handle 3DS on connector requests
type ScaPolicy struct {
	// ShouldIdentify indicates if the identification stages should take place
	ShouldIdentify bool `json:"shouldIdentify" yaml:"shouldIdentify" validate:"-"`

	// ShouldChallengeOptional challenge based on an optional response from the connector (setting this to false will not display the challenge)
	ShouldChallengeOptional bool `json:"shouldChallengeOptional" yaml:"shouldChallengeOptional" validate:"-"`

	// ShouldByPassChallenge if the challenge is required, bypassing this will attempt an auth without displaying the challenge
	ShouldByPassChallenge SCABypassMode `json:"shouldByPassChallenge" yaml:"shouldByPassChallenge" validate:"omitempty,oneof=cascade current"`

	// ShouldChallenge3dSecureV1 determines if the connector can fallback to 3DS v1 when 3DS v2 is not available
	// Deprecated: ChargeHive does not currently support v1 SCA
	ShouldChallenge3dSecureV1 bool `json:"shouldChallenge3dSecureV1,omitempty" yaml:"shouldChallenge3dSecureV1,omitempty" validate:"-"`

	// ShouldAuthOnError if true and an error response is returned from the connector; proceed to auth anyway
	ShouldAuthOnError bool `json:"shouldAuthOnError" yaml:"shouldAuthOnError" validate:"-"`
}

// GetKind returns the ScaPolicy kind
func (ScaPolicy) GetKind() object.Kind { return KindPolicySCA }

// GetVersion returns the ScaPolicy version
func (ScaPolicy) GetVersion() string { return "v1" }

// NewScaDefinition creates a new ScaDefinition
func NewScaDefinition(d *object.Definition) (*ScaDefinition, error) {
	if _, ok := d.Spec.(*ScaPolicy); ok {
		return &ScaDefinition{def: d}, nil
	}
	return nil, errors.New("invalid sca policy object")
}

// ScaDefinition is the SCA config object definition
type ScaDefinition struct{ def *object.Definition }

// Definition returns the ScaDefinition structure
func (d *ScaDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the ScaDefinition
func (d *ScaDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the ScaPolicy contained within the ScaDefinition
func (d *ScaDefinition) Spec() *ScaPolicy { return d.def.Spec.(*ScaPolicy) }
