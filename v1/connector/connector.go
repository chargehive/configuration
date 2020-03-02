package connector

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/object"
)

// KindConnector is the identifier for a Connector config
const KindConnector object.Kind = "Connector"

// Connector is a configuration file for a single payment processing entity
type Connector struct {
	Library       string `json:"library" yaml:"library" validate:"required,oneof=sandbox authorize braintree qualpay stripe paysafe paysafe-applepay paysafe-googlepay worldpay paypal-websitepaymentspro paypal-expresscheckout vindicia chargehive maxmind cybersource"`
	Configuration []byte `json:"configuration" yaml:"configuration" validate:"required"`
}

// GetKind returns the Connector kind
func (Connector) GetKind() object.Kind { return KindConnector }

// GetVersion returns the Connector version
func (Connector) GetVersion() string { return "v1" }

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
