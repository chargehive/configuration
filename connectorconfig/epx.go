package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/grpc/cht"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type EpxEnvironment string

const (
	EpxEnvironmentSandbox    EpxEnvironment = "sandbox"
	EpxEnvironmentProduction EpxEnvironment = "production"
)

type EpxCredentials struct {
	CustNbr                 *string        `json:"custNbr" yaml:"custNbr" validate:"required,gt=0"`
	MerchNbr                *string        `json:"merchNbr" yaml:"merchNbr" validate:"required,gt=0"`
	DbaNbr                  string         `json:"dbaNbr" yaml:"dbaNbr" validate:"required,gt=0"`
	TerminalNbr             *string        `json:"terminalNbr" yaml:"terminalNbr" validate:"required,gt=0"`
	MerchantDescriptor      string         `json:"merchantDescriptor" yaml:"merchantDescriptor" validate:"-"`
	MerchantDescriptorPhone string         `json:"merchantDescriptorPhone" yaml:"merchantDescriptorPhone" validate:"-"`
	Environment             EpxEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
}

func (c *EpxCredentials) GetMID() string {
	return c.DbaNbr
}

func (c *EpxCredentials) GetLibrary() Library {
	return LibraryEpx
}

func (c *EpxCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *EpxCredentials) Validate() error {
	return nil
}

func (c *EpxCredentials) GetSecureFields() []*string {
	return []*string{c.CustNbr, c.MerchNbr, c.TerminalNbr}
}

func (c *EpxCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *EpxCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *EpxCredentials) SupportsSca() bool {
	return true
}

func (c *EpxCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *EpxCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *EpxCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeSandbox && c.Environment == EpxEnvironmentProduction {
		return false
	}
	return true
}

func (c *EpxCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *EpxCredentials) Supports3RI() bool {
	return false
}

func (c *EpxCredentials) IsAccountUpdater() bool {
	return false
}

func (c *EpxCredentials) SupportedTokenSources() []cht.TokenSource {
	return []cht.TokenSource{cht.TS_PAN, cht.TS_NETWORK_TOKEN, cht.TS_APPLE_PAY, cht.TS_GOOGLE_PAY}
}
