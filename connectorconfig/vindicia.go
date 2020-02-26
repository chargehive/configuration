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
	Login         string              `json:"login" yaml:"login" validate:"required"`
	Password      *string             `json:"password" yaml:"password" validate:"required"`
	HMACKey       *string             `json:"hmacKey" yaml:"hmacKey" validate:"required"`
	PGPPrivateKey *string             `json:"pgpPrivateKey" yaml:"pgpPrivateKey" validate:"required"`
	Environment   VindiciaEnvironment `json:"environment" yaml:"environment" validate:"required"`
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
	return []*string{c.Password, c.HMACKey, c.PGPPrivateKey}
}

func (c *VindiciaCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *VindiciaCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
