package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type ChargeHiveCredentials struct {
}

func (c *ChargeHiveCredentials) GetMID() string {
	return ""
}

func (c *ChargeHiveCredentials) GetLibrary() Library {
	return LibraryChargeHive
}

func (c *ChargeHiveCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeFraud}
}

func (c *ChargeHiveCredentials) Validate() error {
	return nil
}

func (c *ChargeHiveCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *ChargeHiveCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *ChargeHiveCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *ChargeHiveCredentials) SupportsSca() bool {
	return false
}

func (c *ChargeHiveCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *ChargeHiveCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *ChargeHiveCredentials) CanPlanModeUse(environment.Mode) bool {
	return true
}

func (c *ChargeHiveCredentials) IsRecoveryAgent() bool {
	return false
}
