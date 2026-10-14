package policy

import (
	"encoding/json"
	"errors"
	"slices"

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

	// FraudCheckTimeAuthDecline indicates that a fraud scan should run after a declined auth
	FraudCheckTimeAuthDecline FraudCheckTime = "auth-decline"

	// FraudCheckTimeCaptureSuccess indicates that a fraud scan should run after a successful
	// capture, which is a sale or the capture of an earlier auth
	FraudCheckTimeCaptureSuccess FraudCheckTime = "capture-success"

	// FraudCheckTimeCaptureDecline indicates that a fraud scan should run after a declined capture
	FraudCheckTimeCaptureDecline FraudCheckTime = "capture-decline"

	// FraudCheckTimeOnDemand indicates that a fraud scan should run on demand only
	FraudCheckTimeOnDemand FraudCheckTime = "ondemand"
)

// ValidFraudCheckTimes reports whether a policy may run at all of the given times together. A
// pre-auth policy runs before either the first auth or every auth, and an on-demand policy runs
// at no other time.
func ValidFraudCheckTimes(times []FraudCheckTime) bool {
	has := map[FraudCheckTime]bool{}
	for _, t := range times {
		has[t] = true
	}
	if has[FraudCheckTimePreFirstAuth] && has[FraudCheckTimePreEveryAuth] {
		return false
	}
	return !has[FraudCheckTimeOnDemand] || len(has) == 1
}

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
	//
	// Deprecated: use CheckTimes.
	CheckTime FraudCheckTime `json:"checkTime,omitempty" yaml:"checkTime,omitempty" validate:"required_without=CheckTimes,excluded_with=CheckTimes,omitempty,oneof=preauth-first preauth-every auth-success ondemand"`

	// CheckTimes are the times we should be running a fraud scan
	CheckTimes []FraudCheckTime `json:"checkTimes,omitempty" yaml:"checkTimes,omitempty" validate:"required_without=CheckTime,omitempty,unique,fraud-check-times,dive,oneof=preauth-first preauth-every auth-success auth-decline capture-success capture-decline ondemand"`

	// CheckType is the type of check that should be performed for this policy
	CheckType FraudCheckType `json:"checkType" yaml:"checkType" validate:"oneof=all failover"`

	DeclineOn []string `json:"declineOn" yaml:"declineOn" validate:"dive,oneof=invalid no-suggestion review deny"`
}

// RunsAt reports whether the policy runs at the given time.
func (p FraudPolicy) RunsAt(t FraudCheckTime) bool {
	return p.CheckTime == t || slices.Contains(p.CheckTimes, t)
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
