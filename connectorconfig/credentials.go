package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type Credentials interface {
	GetLibrary() Library
	GetSupportedTypes() []LibraryType
	GetSecureFields() []*string
	Validate() error
	ToConnector() connector.Connector
	FromJson(input []byte) error
	SupportsSca() bool
	SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool
	SupportsCountry(country string) bool
	CanPlanModeUse(mode environment.Mode) bool // Determine if this plan mode can use this configuration
	IsRecoveryAgent() bool
	Supports3RI() bool
}

type NoLibrary map[string]any

func (n NoLibrary) GetLibrary() Library {
	return LibraryNone
}

func (n NoLibrary) GetSupportedTypes() []LibraryType {
	return nil
}

func (n NoLibrary) GetSecureFields() []*string {
	return nil
}

func (n NoLibrary) Validate() error {
	return nil
}

func (n NoLibrary) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(n.GetLibrary())}
	con.Configuration, _ = json.Marshal(n)
	return con
}

func (n NoLibrary) FromJson(input []byte) error {
	return json.Unmarshal(input, &n)
}

func (n NoLibrary) SupportsSca() bool {
	return false
}

func (n NoLibrary) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	// returning false ensures that this library will not be used for any payment method
	// the intention is to use this to inject configuration for other libraries
	return false
}

func (n NoLibrary) SupportsCountry(country string) bool {
	return false
}

func (n NoLibrary) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (n NoLibrary) IsRecoveryAgent() bool {
	return false
}

func (n NoLibrary) Supports3RI() bool {
	return false
}
