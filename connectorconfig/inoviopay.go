package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/grpc/cht"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type InovioPayEnvironment string

const (
	InovioPayEnvironmentSandbox    InovioPayEnvironment = "sandbox"
	InovioPayEnvironmentProduction InovioPayEnvironment = "production"
)

type InovioPayCredentials struct {
	Username          *string              `json:"username" yaml:"username" validate:"required,gt=0"`
	Password          *string              `json:"password" yaml:"password" validate:"required,gt=0"`
	SiteID            string               `json:"siteId" yaml:"siteId" validate:"required,gt=0"`
	ProductID         string               `json:"productId" yaml:"productId" validate:"required,gt=0"`
	MerchantAccountID string               `json:"merchantAccountID" yaml:"MerchantAccountID" validate:"required,gt=0"`
	Environment       InovioPayEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
}

func (c *InovioPayCredentials) GetMID() string {
	return c.MerchantAccountID
}

func (c *InovioPayCredentials) GetLibrary() Library {
	return LibraryInovioPay
}

func (c *InovioPayCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *InovioPayCredentials) Validate() error {
	return nil
}

func (c *InovioPayCredentials) GetSecureFields() []*string {
	return []*string{c.Username, c.Password}
}

func (c *InovioPayCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *InovioPayCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *InovioPayCredentials) SupportsSca() bool {
	return true
}

func (c *InovioPayCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *InovioPayCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *InovioPayCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeSandbox && c.Environment == InovioPayEnvironmentProduction {
		return false
	}
	return true
}

func (c *InovioPayCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *InovioPayCredentials) Supports3RI() bool {
	return false
}

func (c *InovioPayCredentials) IsAccountUpdater() bool {
	return false
}

func (c *InovioPayCredentials) SupportedTokenSources() []cht.TokenSource {
	return []cht.TokenSource{cht.TS_PAN}
}
