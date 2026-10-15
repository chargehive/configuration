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

// FraudCheckTime are the types of times that we can perform a fraud scan.
//
// Deprecated: set TransactionType, Position, Attempts, CIT and MIT instead.
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

// FraudTransactionType is the transaction a fraud policy runs at.
type FraudTransactionType string

const (
	// FraudTransactionTypeAuth is an authorisation.
	FraudTransactionTypeAuth FraudTransactionType = "auth"
	// FraudTransactionTypeCaptureAuth is the capture of an earlier authorisation.
	FraudTransactionTypeCaptureAuth FraudTransactionType = "capture-auth"
	// FraudTransactionTypeCapture is a payment captured without an earlier authorisation.
	FraudTransactionTypeCapture FraudTransactionType = "capture"
)

// FraudPosition is when a fraud policy runs relative to its transaction.
type FraudPosition string

const (
	// FraudPositionPre runs before the transaction, where a scan's result can decline it.
	FraudPositionPre FraudPosition = "pre"
	// FraudPositionPost runs after the transaction, approved or declined.
	FraudPositionPost FraudPosition = "post"
	// FraudPositionPostSuccess runs after the transaction is approved.
	FraudPositionPostSuccess FraudPosition = "post-success"
	// FraudPositionPostDecline runs after the transaction is declined.
	FraudPositionPostDecline FraudPosition = "post-decline"
	// FraudPositionOnDemand runs when a fraud scan is requested for the charge, outside any
	// transaction.
	FraudPositionOnDemand FraudPosition = "on-demand"
)

// FraudAttempts is which of a charge's attempts at the transaction a fraud policy runs on.
type FraudAttempts string

const (
	// FraudAttemptsFirst runs on the charge's first attempt at the transaction only.
	FraudAttemptsFirst FraudAttempts = "first"
	// FraudAttemptsEvery runs on every attempt.
	FraudAttemptsEvery FraudAttempts = "every"
)

// FraudMode is what a fraud policy does with a payment.
type FraudMode string

const (
	// FraudModeScan asks the connectors for a result. The merchant receives it in the fraud scan
	// webhook, and before the transaction, declineOn can decline the transaction on it.
	FraudModeScan FraudMode = "scan"
	// FraudModeInform sends the payment to the connectors for their records, without asking for a
	// result. A connector that has no way to receive a payment without scanning it is skipped.
	FraudModeInform FraudMode = "inform"
	// FraudModeNone leaves the payment out.
	FraudModeNone FraudMode = "none"
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

	// CheckTime is the time we should be running a fraud scan.
	//
	// Deprecated: set TransactionType, Position, Attempts, CIT and MIT instead.
	CheckTime FraudCheckTime `json:"checkTime,omitempty" yaml:"checkTime,omitempty" validate:"required_without=Position,excluded_with=Position,omitempty,oneof=preauth-first preauth-every auth-success ondemand"`

	// TransactionType is the transaction the policy runs at. An on-demand policy has none.
	TransactionType FraudTransactionType `json:"transactionType,omitempty" yaml:"transactionType,omitempty" validate:"omitempty,oneof=auth capture-auth capture"`

	// Position is when the policy runs relative to its transaction.
	Position FraudPosition `json:"position,omitempty" yaml:"position,omitempty" validate:"fraud-trigger,omitempty,oneof=pre post post-success post-decline on-demand"`

	// Attempts is which of the charge's attempts at the transaction the policy runs on: every one
	// when not set.
	Attempts FraudAttempts `json:"attempts,omitempty" yaml:"attempts,omitempty" validate:"omitempty,oneof=first every"`

	// CIT is what the policy does with a customer-initiated payment: none when not set.
	CIT FraudMode `json:"cit,omitempty" yaml:"cit,omitempty" validate:"omitempty,oneof=scan inform none"`

	// MIT is what the policy does with a merchant-initiated payment: none when not set.
	MIT FraudMode `json:"mit,omitempty" yaml:"mit,omitempty" validate:"omitempty,oneof=scan inform none"`

	// CheckType is the type of check that should be performed for this policy
	CheckType FraudCheckType `json:"checkType" yaml:"checkType" validate:"oneof=all failover"`

	// DeclineOn are the scan results that decline the transaction, when the policy scans before it.
	DeclineOn []string `json:"declineOn" yaml:"declineOn" validate:"dive,oneof=invalid no-suggestion review deny"`
}

// FraudTrigger is when a fraud policy runs and what it does there.
type FraudTrigger struct {
	TransactionType FraudTransactionType
	Position        FraudPosition
	Attempts        FraudAttempts
	CIT             FraudMode
	MIT             FraudMode
}

// Trigger returns when the policy runs and what it does there, reading a CheckTime as the trigger it
// stands for.
func (p FraudPolicy) Trigger() FraudTrigger {
	switch p.CheckTime {
	case FraudCheckTimePreFirstAuth:
		return FraudTrigger{FraudTransactionTypeAuth, FraudPositionPre, FraudAttemptsFirst, FraudModeScan, FraudModeScan}
	case FraudCheckTimePreEveryAuth:
		return FraudTrigger{FraudTransactionTypeAuth, FraudPositionPre, FraudAttemptsEvery, FraudModeScan, FraudModeScan}
	case FraudCheckTimeAuthSuccess:
		return FraudTrigger{FraudTransactionTypeAuth, FraudPositionPostSuccess, FraudAttemptsEvery, FraudModeScan, FraudModeScan}
	case FraudCheckTimeOnDemand:
		return FraudTrigger{"", FraudPositionOnDemand, FraudAttemptsEvery, FraudModeScan, FraudModeScan}
	}
	t := FraudTrigger{p.TransactionType, p.Position, p.Attempts, p.CIT, p.MIT}
	if t.Attempts == "" {
		t.Attempts = FraudAttemptsEvery
	}
	if t.CIT == "" {
		t.CIT = FraudModeNone
	}
	if t.MIT == "" {
		t.MIT = FraudModeNone
	}
	return t
}

// Mode returns what the trigger does with a payment initiated by the customer or the merchant.
func (t FraudTrigger) Mode(customerInitiated bool) FraudMode {
	if customerInitiated {
		return t.CIT
	}
	return t.MIT
}

// Runs reports whether the trigger runs at the transaction and position given, where position is
// pre, post-success or post-decline for the transaction's outcome, or on-demand. A post trigger
// runs after either outcome.
func (t FraudTrigger) Runs(transactionType FraudTransactionType, position FraudPosition) bool {
	if t.TransactionType != transactionType {
		return false
	}
	if t.Position == FraudPositionPost {
		return position == FraudPositionPostSuccess || position == FraudPositionPostDecline
	}
	return t.Position == position
}

// ValidFraudTrigger reports whether a policy's trigger fields fit together: an on-demand policy
// has no transaction type, every other policy has one, and a policy does something with at least
// one initiator's payments.
func ValidFraudTrigger(p FraudPolicy) bool {
	if p.Position == "" {
		return p.TransactionType == "" && p.Attempts == "" && p.CIT == "" && p.MIT == ""
	}
	if (p.Position == FraudPositionOnDemand) != (p.TransactionType == "") {
		return false
	}
	t := p.Trigger()
	return t.CIT != FraudModeNone || t.MIT != FraudModeNone
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
