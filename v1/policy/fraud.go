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

	// FraudCheckTimeAuthSuccess indicates that a fraud scan should run after a successful auth.
	//
	// Deprecated: use FraudCheckTimePostAuth, which a CheckTime of auth-success runs as.
	FraudCheckTimeAuthSuccess FraudCheckTime = "auth-success"

	// FraudCheckTimePostAuth indicates that a fraud scan should run after an auth's outcome, approved
	// or declined, while an approved auth can still be voided before it is captured
	FraudCheckTimePostAuth FraudCheckTime = "postauth"

	// FraudCheckTimePostCapture indicates that a fraud scan should run after a capture's outcome,
	// collected or declined. A capture is a sale or the capture of an earlier auth.
	FraudCheckTimePostCapture FraudCheckTime = "postcapture"

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

	// CheckTime is the time we should be running a fraud scan, for every payment
	//
	// Deprecated: use CustomerCheckTime and MerchantCheckTime.
	CheckTime FraudCheckTime `json:"checkTime,omitempty" yaml:"checkTime,omitempty" validate:"required_without_all=CustomerCheckTime MerchantCheckTime,excluded_with=CustomerCheckTime MerchantCheckTime,omitempty,oneof=preauth-first preauth-every auth-success ondemand"`

	// CustomerCheckTime is when to run a fraud scan for a customer-initiated payment. Unset, the
	// policy does not check customer-initiated payments.
	CustomerCheckTime FraudCheckTime `json:"customerCheckTime,omitempty" yaml:"customerCheckTime,omitempty" validate:"omitempty,oneof=preauth-first preauth-every postauth postcapture ondemand"`

	// MerchantCheckTime is when to run a fraud scan for a merchant-initiated payment, which has no
	// customer present to collect device data from before the auth. Unset, the policy does not
	// check merchant-initiated payments.
	MerchantCheckTime FraudCheckTime `json:"merchantCheckTime,omitempty" yaml:"merchantCheckTime,omitempty" validate:"omitempty,oneof=postauth postcapture ondemand"`

	// CheckType is the type of check that should be performed for this policy
	CheckType FraudCheckType `json:"checkType" yaml:"checkType" validate:"oneof=all failover"`

	DeclineOn []string `json:"declineOn" yaml:"declineOn" validate:"dive,oneof=invalid no-suggestion review deny"`
}

// CheckTimeFor returns when the policy runs for a payment initiated by the customer or the merchant,
// or "" when it does not check that payment.
func (p FraudPolicy) CheckTimeFor(customerInitiated bool) FraudCheckTime {
	if p.CheckTime != "" {
		if p.CheckTime == FraudCheckTimeAuthSuccess {
			return FraudCheckTimePostAuth
		}
		return p.CheckTime
	}
	if customerInitiated {
		return p.CustomerCheckTime
	}
	return p.MerchantCheckTime
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
