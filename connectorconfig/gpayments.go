package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type GPaymentsEnvironment string

const (
	GPaymentsEnvironmentSandbox    GPaymentsEnvironment = "sandbox"
	GPaymentsEnvironmentProduction GPaymentsEnvironment = "production"
)

type GPaymentsCredentials struct {
	MerchantName                string               `json:"merchantName" yaml:"merchantName" validate:"required,gt=0"`
	MerchantID                  string               `json:"merchantID" yaml:"merchantID" validate:"required,gt=0"`
	MerchantCertificate         *string              `json:"merchantCertificate" yaml:"merchantCertificate" validate:"required,gt=0"`
	MerchantCertificatePassword *string              `json:"merchantCertificatePassword" yaml:"merchantCertificatePassword" validate:"required,gt=0"`
	CACertificates              *string              `json:"CACertificates" yaml:"CACertificates" validate:"required,gt=0"`
	Environment                 GPaymentsEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
}

func (c *GPaymentsCredentials) GetMID() string {
	return c.MerchantID
}

func (c *GPaymentsCredentials) GetLibrary() Library {
	return LibraryGPayments
}

func (c *GPaymentsCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeAuthentication}
}

func (c *GPaymentsCredentials) Validate() error {
	return nil
}

func (c *GPaymentsCredentials) GetSecureFields() []*string {
	return []*string{c.MerchantCertificate, c.MerchantCertificatePassword}
}

func (c *GPaymentsCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *GPaymentsCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *GPaymentsCredentials) SupportsSca() bool {
	return true
}

func (c *GPaymentsCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *GPaymentsCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *GPaymentsCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeSandbox && c.Environment == GPaymentsEnvironmentProduction {
		return false
	}
	return true
}

func (c *GPaymentsCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *GPaymentsCredentials) Supports3RI() bool {
	return true
}

func (c *GPaymentsCredentials) IsAccountUpdater() bool {
	return false
}
