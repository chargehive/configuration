package connectorconfig

import (
	"encoding/json"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type ThreeDSecureIoCredentials struct {
	APIKey *string `json:"apiKey" yaml:"apiKey" validate:"required,gt=0"`
}

func (c ThreeDSecureIoCredentials) GetLibrary() Library {
	return LibraryThreeDSecureIo
}

func (c *ThreeDSecureIoCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeSCA}
}

func (c *ThreeDSecureIoCredentials) Validate() error {
	return nil
}

func (c *ThreeDSecureIoCredentials) GetSecureFields() []*string {
	return []*string{c.APIKey}
}

func (c *ThreeDSecureIoCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *ThreeDSecureIoCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c ThreeDSecureIoCredentials) SupportsSca() bool {
	return true
}

func (c ThreeDSecureIoCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if methodType == chtype.PAYMENT_METHOD_TYPE_CARD {
		return true
	}
	return false
}
