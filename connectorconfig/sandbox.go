package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type SandboxMode string

const (
	SandboxModeDynamic       SandboxMode = "dynamic"
	SandboxModeOffline       SandboxMode = "offline"
	SandboxModeDelayed       SandboxMode = "delayed"
	SandboxModeRandomTimeout SandboxMode = "random-timeout"
	SandboxModeChaos         SandboxMode = "chaos"
)

type SandboxCredentials struct {
	MerchantID          string      `json:"merchantId" yaml:"merchantID" validate:"-"`
	Mode                SandboxMode `json:"mode" yaml:"mode" validate:"oneof=dynamic offline delayed random-timeout chaos"`
	TransactionIDPrefix string      `json:"transactionIDPrefix" yaml:"transactionIDPrefix" validate:"-"`
}

func (c *SandboxCredentials) GetMID() string {
	return c.MerchantID
}

func (c *SandboxCredentials) GetLibrary() Library {
	return LibrarySandbox
}

func (c *SandboxCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *SandboxCredentials) Validate() error {
	return nil
}

func (c *SandboxCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *SandboxCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *SandboxCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *SandboxCredentials) SupportsSca() bool {
	return true
}

func (c *SandboxCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *SandboxCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *SandboxCredentials) CanPlanModeUse(environment.Mode) bool {
	return true
}

func (c *SandboxCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *SandboxCredentials) Supports3RI() bool {
	return false
}

func (c *SandboxCredentials) IsAccountUpdater() bool {
	return false
}
