package policy

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

const KindPolicySCA object.Kind = "PolicySCA"

type SCABypassMode string

const (
	SCABypassModeCascade SCABypassMode = "cascade" //Attempt the auth on the next connector
	SCABypassModeCurrent SCABypassMode = "current" //Bypass, but stay on the current connector
)

type ScaPolicy struct {
	ShouldIdentify          bool
	ShouldChallengeOptional bool
	ShouldByPassChallenge   bool
	ChallengeByPassMode     SCABypassMode
}

func (ScaPolicy) GetKind() object.Kind { return KindPolicySCA }
func (ScaPolicy) GetVersion() string   { return "v1" }

func NewScaDefinition(d *object.Definition) (*Definition, error) {
	if _, ok := d.Spec.(*ScaPolicy); ok {
		return &Definition{def: d}, nil
	}
	return nil, errors.New("invalid sca policy object")
}

type Definition struct{ def *object.Definition }

func (d *Definition) Definition() *object.Definition { return d.def }
func (d *Definition) MarshalJSON() ([]byte, error)   { return json.Marshal(d.def) }
func (d *Definition) Spec() *ScaPolicy               { return d.def.Spec.(*ScaPolicy) }
