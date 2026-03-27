package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/configuration/v1/scheduler"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type GoogleStoreCredentials struct {
}

func (c *GoogleStoreCredentials) GetMID() string {
	return ""
}

func (c *GoogleStoreCredentials) GetLibrary() Library {
	return LibraryGoogleStore
}
func (c *GoogleStoreCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeThirdPartyStore}
}

func (c *GoogleStoreCredentials) Validate() error {
	return nil
}

func (c *GoogleStoreCredentials) GetSecureFields() []*string {
	return nil
}

func (c *GoogleStoreCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *GoogleStoreCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *GoogleStoreCredentials) SupportsSca() bool {
	return false
}

func (c *GoogleStoreCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *GoogleStoreCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *GoogleStoreCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (c *GoogleStoreCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *GoogleStoreCredentials) Supports3RI() bool {
	return false
}

func (c *GoogleStoreCredentials) IsAccountUpdater() bool {
	return false
}

func (c *GoogleStoreCredentials) SupportedTokenTypes() []scheduler.TokenSource {
	return nil
}
