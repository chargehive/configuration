package policy

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

const KindPolicySCA object.Kind = "PolicySCA"

type SCABypassMode string

const (
	SCABypassModeNone    SCABypassMode = ""        //Do not bypass a required challenge
	SCABypassModeCascade SCABypassMode = "cascade" //Attempt the auth on the next connector
	SCABypassModeCurrent SCABypassMode = "current" //Bypass, but stay on the current connector
)

type ScaPolicy struct {
	ShouldIdentify            bool
	ShouldChallengeOptional   bool          //If the challenge is optional, setting this to false will not display the challenge
	ShouldByPassChallenge     SCABypassMode //If the challenge is required, bypassing this will attempt an auth without displaying the challenge
	ShouldChallenge3dSecureV1 bool
}

func (ScaPolicy) GetKind() object.Kind { return KindPolicySCA }
func (ScaPolicy) GetVersion() string   { return "v1" }

func NewScaDefinition(d *object.Definition) (*ScaDefinition, error) {
	if _, ok := d.Spec.(*ScaPolicy); ok {
		return &ScaDefinition{def: d}, nil
	}
	return nil, errors.New("invalid sca policy object")
}

type ScaDefinition struct{ def *object.Definition }

func (d *ScaDefinition) Definition() *object.Definition { return d.def }
func (d *ScaDefinition) MarshalJSON() ([]byte, error)   { return json.Marshal(d.def) }
func (d *ScaDefinition) Spec() *ScaPolicy               { return d.def.Spec.(*ScaPolicy) }
