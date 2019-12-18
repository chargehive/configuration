package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type VindiciaEnvironment string

const (
	VindiciaEnvironmentDevelopment VindiciaEnvironment = "development"
	VindiciaEnvironmentStage       VindiciaEnvironment = "stage"
	VindiciaEnvironmentProduction  VindiciaEnvironment = "production"
)

type VindiciaCredentials struct {
	MerchantID     string
	TransactionKey *string
	Environment    VindiciaEnvironment
}

func (c VindiciaCredentials) GetLibrary() Library {
	return LibraryVindicia
}

func (c *VindiciaCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *VindiciaCredentials) Validate() error {
	return nil
}

func (c *VindiciaCredentials) GetSecureFields() []*string {
	return []*string{c.TransactionKey}
}

func (c *VindiciaCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *VindiciaCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
