package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type ChargeHiveCredentials struct {
}

func (c ChargeHiveCredentials) GetLibrary() Library {
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
