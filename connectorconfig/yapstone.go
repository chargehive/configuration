package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type YapstoneCredentials struct {
}

func (c *YapstoneCredentials) GetLibrary() Library {
	return LibraryYapstone
}

func (c *YapstoneCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *YapstoneCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *YapstoneCredentials) Validate() error {
	return nil
}

func (c *YapstoneCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *YapstoneCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *YapstoneCredentials) SupportsSca() bool {
	return false
}

func (c *YapstoneCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}

	return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
}

func (c *YapstoneCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (c *YapstoneCredentials) IsRecoveryAgent() bool {
	return false
}
