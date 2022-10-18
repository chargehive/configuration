package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type CWAMSCredentials struct {
	SecurityKey string `json:"securityKey" yaml:"securityKey" validate:"required"`
}

func (c *CWAMSCredentials) GetLibrary() Library {
	return LibraryCWAMS
}

func (c *CWAMSCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *CWAMSCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *CWAMSCredentials) Validate() error {
	return nil
}

func (c *CWAMSCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *CWAMSCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *CWAMSCredentials) SupportsSca() bool {
	return false
}

func (c *CWAMSCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}

	return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
}

func (c *CWAMSCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (c *CWAMSCredentials) IsRecoveryAgent() bool {
	return false
}
