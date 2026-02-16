package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type VindiciaEnvironment string

const (
	VindiciaEnvironmentDevelopment VindiciaEnvironment = "development"
	VindiciaEnvironmentStage       VindiciaEnvironment = "stage"
	VindiciaEnvironmentProduction  VindiciaEnvironment = "production"
)

type ConnectorAttempt struct {
	ConnectorID    string `json:"connectorID"`
	DivisionNumber string `json:"divisionNumber"`
	Weight         int    `json:"weight"`
}
type VindiciaCredentials struct {
	Login         string              `json:"login" yaml:"login" validate:"required"`
	Password      *string             `json:"password" yaml:"password" validate:"required,gt=0"`
	HMACKey       *string             `json:"hmacKey" yaml:"hmacKey" validate:"required,gt=0"`
	PGPPrivateKey *string             `json:"pgpPrivateKey" yaml:"pgpPrivateKey" validate:"required,gt=0"`
	DirectRefund  bool                `json:"directRefund" yaml:"directRefund" validate:"-"`
	ConnectorPool []ConnectorAttempt  `json:"connectorPool" yaml:"connectorPool" validate:"required"`
	Environment   VindiciaEnvironment `json:"environment" yaml:"environment" validate:"oneof=development stage production"`
}

func (c *VindiciaCredentials) GetMID() string {
	return c.Login
}

func (c *VindiciaCredentials) GetLibrary() Library {
	return LibraryVindicia
}

func (c *VindiciaCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeRecoveryAgent}
}

func (c *VindiciaCredentials) Validate() error {
	return nil
}

func (c *VindiciaCredentials) GetSecureFields() []*string {
	return []*string{c.Password, c.HMACKey, c.PGPPrivateKey}
}

func (c *VindiciaCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *VindiciaCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *VindiciaCredentials) SupportsSca() bool {
	return false
}

func (c *VindiciaCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *VindiciaCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *VindiciaCredentials) CanPlanModeUse(mode environment.Mode) bool {
	if mode == environment.ModeSandbox && c.Environment == VindiciaEnvironmentProduction {
		return false
	}
	return true
}

func (c *VindiciaCredentials) IsRecoveryAgent() bool {
	return true
}

func (c *VindiciaCredentials) Supports3RI() bool {
	return false
}

func (c *VindiciaCredentials) IsAccountUpdater() bool {
	return false
}
