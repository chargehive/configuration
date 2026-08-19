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

type ChallengePreference string

const (
	ChallengePreferenceNone        ChallengePreference = "no-preference"
	ChallengePreferenceNoChallenge ChallengePreference = "no-challenge"
	ChallengePreferenceRequest     ChallengePreference = "request"
	ChallengePreferenceMandate     ChallengePreference = "mandate"
)

// ScaPolicy options determine how to handle 3DS on connector requests
type ScaPolicy struct {
	// RequireSca indicates if a transaction will require SCA facilities. This is used to filter out connectors which cannot complete SCA
	RequireSca *bool `json:"requireSca" yaml:"requireSca" validate:"required"`

	// ShouldIdentify indicates if the identification stages should take place
	ShouldIdentify *bool `json:"shouldIdentify" yaml:"shouldIdentify" validate:"required"`

	// ShouldChallengeOptional challenge based on an optional response from the connector (setting this to false will not display the challenge)
	ShouldChallengeOptional *bool `json:"shouldChallengeOptional" yaml:"shouldChallengeOptional" validate:"required"`

	// ShouldByPassChallenge if the challenge is required, bypassing this will attempt an auth without displaying the challenge
	ShouldByPassChallenge SCABypassMode `json:"shouldByPassChallenge" yaml:"shouldByPassChallenge" validate:"omitempty,oneof=cascade current"`

	// ShouldChallenge3dSecureV1 determines if the connector can fallback to 3DS v1 when 3DS v2 is not available
	// Deprecated: ChargeHive does not currently support v1 SCA
	ShouldChallenge3dSecureV1 bool `json:"shouldChallenge3dSecureV1,omitempty" yaml:"shouldChallenge3dSecureV1,omitempty" validate:"-"`

	// ShouldAuthOnError if true and an error response is returned from the connector; proceed to auth anyway
	ShouldAuthOnError *bool `json:"shouldAuthOnError" yaml:"shouldAuthOnError" validate:"required"`

	// ShouldAuthOnN - Not authenticated — issuer/cardholder actively failed auth
	// if true and an "N" response is returned from the connector; proceed to auth anyway
	ShouldAuthOnN *bool `json:"shouldAuthOnN" yaml:"shouldAuthOnN" validate:"required"`

	// ShouldAuthOnR - Rejected by issuer
	// if true and an "R" response is returned from the connector; proceed to auth anyway.
	// Unset resolves to false, unlike the other shouldAuthOnX flags - see GetShouldAuthOnR.
	ShouldAuthOnR *bool `json:"shouldAuthOnR" yaml:"shouldAuthOnR" validate:"required"`

	// ShouldAuthOnU - Unable to authenticate (technical failure, sometimes also covers "no ACS / card not enrolled")
	// if true and an "U" response is returned from the connector; proceed to auth anyway
	ShouldAuthOnU *bool `json:"shouldAuthOnU" yaml:"shouldAuthOnU" validate:"required"`

	ChallengePreference ChallengePreference `json:"challengePreference" yaml:"challengePreference" validate:"omitempty,oneof=no-preference no-challenge request mandate"`
}

// scaFlag resolves an unset policy flag. Every flag on ScaPolicy is a pointer so
// that an omitted field can be told apart from an explicit false, which means each
// getter has to name the value it falls back to. Keeping the resolution here rather
// than in each body puts all of those fallbacks in one readable column.
func scaFlag(value *bool, whenUnset bool) bool {
	if value == nil {
		return whenUnset
	}
	return *value
}

// GetScaRequired reports whether SCA is required, defaulting to false so a policy
// that does not ask for authentication does not get it.
func (s ScaPolicy) GetScaRequired() bool { return scaFlag(s.RequireSca, false) }

// GetShouldIdentify reports whether the identification stages should run, defaulting
// to false for the same reason as GetScaRequired.
func (s ScaPolicy) GetShouldIdentify() bool { return scaFlag(s.ShouldIdentify, false) }

// GetShouldChallengeOptional reports whether an optional challenge should be shown,
// defaulting to false so a challenge is not displayed unless it was asked for.
func (s ScaPolicy) GetShouldChallengeOptional() bool {
	return scaFlag(s.ShouldChallengeOptional, false)
}

// GetShouldAuthOnError reports whether to authorize after a connector error,
// defaulting to true: an error says nothing about the cardholder, so an
// authentication outage should not decline every payment.
func (s ScaPolicy) GetShouldAuthOnError() bool { return scaFlag(s.ShouldAuthOnError, true) }

// GetShouldAuthOnN reports whether to authorize after a failed authentication,
// defaulting to true. N is an outcome the merchant may reasonably decide to accept.
func (s ScaPolicy) GetShouldAuthOnN() bool { return scaFlag(s.ShouldAuthOnN, true) }

// GetShouldAuthOnR reports whether to authorize after a rejected authentication.
//
// Unlike its siblings this defaults to false. N and U are outcomes the merchant can
// weigh up, but R is the issuer rejecting authentication and asking that
// authorization not be attempted at all - so proceeding has to be something a
// merchant chose, never something inherited from an unset field.
func (s ScaPolicy) GetShouldAuthOnR() bool { return scaFlag(s.ShouldAuthOnR, false) }

// GetShouldAuthOnU reports whether to authorize when authentication could not be
// performed, defaulting to true: U is a technical failure rather than a refusal.
func (s ScaPolicy) GetShouldAuthOnU() bool { return scaFlag(s.ShouldAuthOnU, true) }

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
