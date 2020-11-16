package policy

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
	"time"
)

// KindPolicySCA is the identifier for a ScaPolicy config
const KindPolicySCA object.Kind = "PolicySCA"

// SCABypassMode indicates the action to be performed when a verification result is returned by a connector
type SCABypassMode string

const (
	// SCABypassModeNone indicates to not bypass a required challenge
	// Deprecated: OLD CONFIG
	SCABypassModeNone SCABypassMode = ""

	// SCABypassModeCascade indicates to auth on the next connector
	// Deprecated: OLD CONFIG
	SCABypassModeCascade SCABypassMode = "cascade"

	// SCABypassModeCurrent indicate to bypass, but stay on the current connector (attempt auth anyway)
	// Deprecated: OLD CONFIG
	SCABypassModeCurrent SCABypassMode = "current"
)

// ScaPolicy options determine how to handle 3DS on connector requests
type ScaPolicy struct {
	// RequireSca indicates if a transaction will require SCA facilities. This is used to filter out connectors which cannot complete SCA
	// Deprecated: OLD CONFIG
	RequireSca *bool `json:"requireSca" yaml:"requireSca" validate:"required"`

	// ShouldIdentify indicates if the identification stages should take place
	// Deprecated: OLD CONFIG
	ShouldIdentify *bool `json:"shouldIdentify" yaml:"shouldIdentify" validate:"required"`

	// ShouldChallengeOptional challenge based on an optional response from the connector (setting this to false will not display the challenge)
	// Deprecated: OLD CONFIG
	ShouldChallengeOptional *bool `json:"shouldChallengeOptional" yaml:"shouldChallengeOptional" validate:"required"`

	// ShouldByPassChallenge if the challenge is required, bypassing this will attempt an auth without displaying the challenge
	// Deprecated: ShouldByPassChallenge OLD CONFIG
	ShouldByPassChallenge SCABypassMode `json:"shouldByPassChallenge" yaml:"shouldByPassChallenge" validate:"omitempty,oneof=cascade current"`

	// ShouldChallenge3dSecureV1 determines if the connector can fallback to 3DS v1 when 3DS v2 is not available
	// Deprecated: ChargeHive does not currently support v1 SCA
	ShouldChallenge3dSecureV1 bool `json:"shouldChallenge3dSecureV1,omitempty" yaml:"shouldChallenge3dSecureV1,omitempty" validate:"-"`

	// ShouldAuthOnError if true and an error response is returned from the connector; proceed to auth anyway
	// Deprecated: OLD CONFIG
	ShouldAuthOnError *bool `json:"shouldAuthOnError" yaml:"shouldAuthOnError" validate:"required"`

	// New policy components

	// SCAEnabled Enables or disables any SCA actions
	SCAEnabled bool
	// AuthenticateOnSCAError Enable to retry a transaction with SCA if an SCA based error message is returned from an Authorization or Capture
	AuthenticateOnSCAError bool
	// ChallengeAttemptLimit Limits the number of SCA Challenge attempts a customer can try before failing the transaction
	ChallengeAttemptLimit int
	// OfflineChallengeAction Set action to take if an offline SCA frictionless transaction fails to authenticate
	OfflineChallengeAction OfflineChallengeAction
	// OfflineChallengeTimeout Amount of time to wait for a customer to perform a challenge before timing out and continuing
	OfflineChallengeTimeout time.Time
	// OnlineChallengeOverride Enable to bypass requested challenge to perform a normal authorization
	OnlineChallengeOverride bool
	// OnlineChallengeRequested Enable to challenge customer if the challenge is optional and not mandated
	OnlineChallengeRequested bool
}

// Deprecated: Should use new SCA policy methods
func (s ScaPolicy) GetScaRequired() bool {
	if s.RequireSca == nil {
		return false
	}
	return *s.RequireSca
}

// Deprecated: Should use new SCA policy methods
func (s ScaPolicy) GetShouldIdentify() bool {
	if s.ShouldIdentify == nil {
		return false
	}
	return *s.ShouldIdentify
}

// Deprecated: Should use new SCA policy methods
func (s ScaPolicy) GetShouldChallengeOptional() bool {
	if s.ShouldChallengeOptional == nil {
		return false
	}
	return *s.ShouldChallengeOptional
}

// Deprecated: Should use new SCA policy methods
func (s ScaPolicy) GetShouldAuthOnError() bool {
	if s.ShouldAuthOnError == nil {
		return true
	}
	return *s.ShouldAuthOnError
}

type OfflineChallengeAction string

const (
	// will challenge if sca frictionless failed
	OfflineChallengeActionChallenge OfflineChallengeAction = "challenge"
	// will decline the transaction
	OfflineChallengeActionDecline OfflineChallengeAction = "decline"
	// will attempt a transaction through one-leg-out connector if available
	OfflineChallengeActionOneLegOut OfflineChallengeAction = "one-leg-out"
)

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
