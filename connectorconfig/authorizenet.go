package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type AuthorizeCredentials struct {
}

func (c *AuthorizeCredentials) GetMID() string {
	return ""
}

func (c *AuthorizeCredentials) GetLibrary() Library {
	return LibraryAuthorizeNet
}
func (c *AuthorizeCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *AuthorizeCredentials) Validate() error {
	return nil
}

func (c *AuthorizeCredentials) GetSecureFields() []*string {
	return nil
}

func (c *AuthorizeCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *AuthorizeCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *AuthorizeCredentials) SupportsSca() bool {
	return false
}

func (c *AuthorizeCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *AuthorizeCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *AuthorizeCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (c *AuthorizeCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *AuthorizeCredentials) Supports3RI() bool {
	return false
}
