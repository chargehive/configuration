package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type AppleStoreCredentials struct {
}

func (c *AppleStoreCredentials) GetMID() string {
	return ""
}

func (c *AppleStoreCredentials) GetLibrary() Library {
	return LibraryAppleStore
}
func (c *AppleStoreCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeThirdPartyStore}
}

func (c *AppleStoreCredentials) Validate() error {
	return nil
}

func (c *AppleStoreCredentials) GetSecureFields() []*string {
	return nil
}

func (c *AppleStoreCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *AppleStoreCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *AppleStoreCredentials) SupportsSca() bool {
	return false
}

func (c *AppleStoreCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *AppleStoreCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *AppleStoreCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (c *AppleStoreCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *AppleStoreCredentials) Supports3RI() bool {
	return false
}

func (c *AppleStoreCredentials) IsAccountUpdater() bool {
	return false
}
