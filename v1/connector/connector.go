package connector

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/object"
)

// KindConnector is the identifier for a Connector config
const KindConnector object.Kind = "Connector"

// ProcessingState indicates how a connection should be used while processing
type ProcessingState string

const (
	// ProcessingStateLive process all transactions
	ProcessingStateLive ProcessingState = "live"
	// ProcessingStateCoolDown allow existing auths to be captured and refunds
	ProcessingStateCoolDown ProcessingState = "cool-down"
	// ProcessingStateRefundOnly  Only process refund transactions
	ProcessingStateRefundOnly ProcessingState = "refund-only"
)

// Connector is a configuration file for a single payment processing entity
type Connector struct {
	ProcessingState ProcessingState `json:"processingState,omitempty" yaml:"processingState,omitempty"`
	Library         string          `json:"library" yaml:"library" validate:"omitempty,oneof=epx recaptcha flexpay adyen bluesnap gpayments nuvei inoviopay threedsecureio sandbox sandbanx applepay authorizenet braintree qualpay stripe paysafe worldpay paypal-websitepaymentspro paypal-expresscheckout vindicia maxmind cybersource paysafe-accountupdater bottomline checkout kount clearhaus trust-payments cwams yapstone tokenex-accountupdater tokenex-api-accountupdater tokenex-networktokenization sticky-io googlepay paypal-completepayments"`
	ConfigurationID string          `json:"configurationID,omitempty" yaml:"configurationID,omitempty"`
	Configuration   []byte          `json:"configuration" yaml:"configuration" validate:"required"`
	ConfigID        string          `json:"configId,omitempty" yaml:"configId,omitempty"`
	ConfigAuth      string          `json:"configAuth,omitempty" yaml:"configAuth,omitempty"`
	EnablePCIB      bool            `json:"enablePCIB,omitempty" yaml:"enablePCIB,omitempty"`
	SCAConnectorID  string          `json:"scaConnectorID,omitempty" yaml:"scaConnectorID,omitempty"`
}

// GetKind returns the Connector kind
func (c Connector) GetKind() object.Kind { return KindConnector }

// GetVersion returns the Connector version
func (c Connector) GetVersion() string { return "v1" }

// NewDefinition returns a new connector definition
func NewDefinition(d *object.Definition) (*Definition, error) {
	if _, ok := d.Spec.(*Connector); ok {
		return &Definition{def: d}, nil
	}
	return nil, errors.New("invalid connector object")
}

// Definition is the connector definition structure
type Definition struct{ def *object.Definition }

// Definition returns the Definition for a connector
func (d *Definition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for a connector
func (d *Definition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the connector specification from within a definition
func (d *Definition) Spec() *Connector { return d.def.Spec.(*Connector) }
