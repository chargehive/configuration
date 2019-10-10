package policy

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

const KindPolicyFraud object.Kind = "PolicyFraud"

type FraudCheckType string

const (
	FraudCheckTypeAll      FraudCheckType = "all"      // Perform a fraud check on all provided connectors
	FraudCheckTypeFailover FraudCheckType = "failover" // Perform a fraud check on one connector ID at a time, stopping at the first success
)

type FraudCheckTime string

const (
	FraudCheckTimePreFirstAuth FraudCheckTime = "preauth-first" // Run fraud check before the first auth
	FraudCheckTimePreEveryAuth FraudCheckTime = "preauth-every" // Run fraud check before every auth
	FraudCheckTimeAuthSuccess  FraudCheckTime = "auth-success"  // Run after a successful auth
	FraudCheckTimeOnDemand     FraudCheckTime = "ondemand"      // Run on demand only
)

type FraudPolicy struct {
	ConnectorIDs []string
	CheckTime    FraudCheckTime
	CheckType    FraudCheckType
}

func (FraudPolicy) GetKind() object.Kind { return KindPolicyFraud }
func (FraudPolicy) GetVersion() string   { return "v1" }

func NewFraudDefinition(d *object.Definition) (*FraudDefinition, error) {
	if _, ok := d.Spec.(*FraudPolicy); ok {
		return &FraudDefinition{def: d}, nil
	}
	return nil, errors.New("invalid fraud policy object")
}

type FraudDefinition struct{ def *object.Definition }

func (d *FraudDefinition) Definition() *object.Definition { return d.def }
func (d *FraudDefinition) MarshalJSON() ([]byte, error)   { return json.Marshal(d.def) }
func (d *FraudDefinition) Spec() *FraudPolicy             { return d.def.Spec.(*FraudPolicy) }
