package policy

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/object"
)

// KindPolicyFraud is the identifier for a PolicyFraud config
const KindPolicyFraud object.Kind = "PolicyFraud"

// FraudCheckType indicates the type of check that should be performed
type FraudCheckType string

const (
	// FraudCheckTypeAll perform a fraud check on all provided connectors
	FraudCheckTypeAll FraudCheckType = "all"

	// FraudCheckTypeFailover will perform a fraud check on one connector ID at a time, stopping at the first success
	FraudCheckTypeFailover FraudCheckType = "failover"
)

// FraudCheckTime are the types of times that we can perform a fraud scan
type FraudCheckTime string

const (
	// FraudCheckTimePreFirstAuth indicates that a fraud scan should check before the first auth
	FraudCheckTimePreFirstAuth FraudCheckTime = "preauth-first"

	// FraudCheckTimePreEveryAuth indicates that a fraud scan should check before every auth
	FraudCheckTimePreEveryAuth FraudCheckTime = "preauth-every"

	// FraudCheckTimeAuthSuccess indicates that a fraud scan should run after a successful auth
	FraudCheckTimeAuthSuccess FraudCheckTime = "auth-success"

	// FraudCheckTimeOnDemand indicates that a fraud scan should run on demand only
	FraudCheckTimeOnDemand FraudCheckTime = "ondemand"
)

// FraudDeclineReason are the suggestions for declining a charge
type FraudDeclineReason string

const (
	FraudDeclineOnInvalid      FraudDeclineReason = "invalid"
	FraudDeclineOnNoSuggestion FraudDeclineReason = "no-suggestion"
	FraudDeclineOnReview       FraudDeclineReason = "review"
	FraudDeclineOnDeny         FraudDeclineReason = "deny"
)

// FraudPolicy is the policy ran against a charge to determine its fraud status
type FraudPolicy struct {
	// ConnectorIDs is the IDs of the fraud connectors
	ConnectorIDs []string `json:"connectorIDs" yaml:"connectorIDs" validate:"min=1,dive,lowercase"`

	// CheckTime is the time we should be running a fraud scan
	CheckTime FraudCheckTime `json:"checkTime" yaml:"checkTime" validate:"oneof=preauth-first preauth-every auth-success ondemand"`

	// CheckType is the type of check that should be performed for this policy
	CheckType FraudCheckType `json:"checkType" yaml:"checkType" validate:"oneof=all failover"`

	DeclineOn []string `json:"declineOn" yaml:"declineOn" validate:"dive,oneof=invalid no-suggestion review deny"`
}

// GetKind returns the FraudPolicy kind
func (FraudPolicy) GetKind() object.Kind { return KindPolicyFraud }

// GetVersion returns the FraudPolicy version
func (FraudPolicy) GetVersion() string { return "v1" }

// NewFraudDefinition creates a new FraudDefinition
func NewFraudDefinition(d *object.Definition) (*FraudDefinition, error) {
	if _, ok := d.Spec.(*FraudPolicy); ok {
		return &FraudDefinition{def: d}, nil
	}
	return nil, errors.New("invalid fraud policy object")
}

// FraudDefinition is the Fraud config object definition
type FraudDefinition struct{ def *object.Definition }

// Definition returns the FraudDefinition structure
func (d *FraudDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the FraudDefinition
func (d *FraudDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the FraudPolicy contained within the FraudDefinition
func (d *FraudDefinition) Spec() *FraudPolicy { return d.def.Spec.(*FraudPolicy) }
