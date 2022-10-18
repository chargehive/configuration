package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type ClearhausEnvironment string

const (
	ClearhausEnvironmentTest ClearhausEnvironment = "test"
	ClearhausEnvironmentLive ClearhausEnvironment = "live"
)

type ClearhausCredentials struct {
	APIKey      string               `json:"apiKey" yaml:"apiKey" validate:"required"`
	Environment ClearhausEnvironment `json:"environment" yaml:"environment" validate:"required,oneof=test live"`
}

func (c *ClearhausCredentials) GetLibrary() Library {
	return LibraryClearhaus
}

func (c *ClearhausCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *ClearhausCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *ClearhausCredentials) Validate() error {
	return nil
}

func (c *ClearhausCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *ClearhausCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *ClearhausCredentials) SupportsSca() bool {
	return false
}

func (c *ClearhausCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}

	return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
}

func (c *ClearhausCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (c *ClearhausCredentials) IsRecoveryAgent() bool {
	return false
}
