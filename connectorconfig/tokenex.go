package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type TokenExEnvironment string

const (
	TokenExEnvironmentSandbox    TokenExEnvironment = "sandbox"
	TokenExEnvironmentProduction TokenExEnvironment = "production"
)

type TokenExRegion string

const (
	TokenExRegionUS TokenExRegion = "us"
	TokenExRegionEU TokenExRegion = "eu"
)

type TokenExAccountUpdaterCredentials struct {
	// Encryption PGP key provided TO you BY TokenEx (note this is a separate pair from DecryptionPGPPrivateKey.)
	EncryptionPGPPublicKey string `json:"encPGPPublic" yaml:"encPGPPublic" validate:"required,gt=0"`

	// The decryption PGP private key is generated BY you and the public key component you have provided TO TokenEx
	DecryptionPGPPrivateKey           *string `json:"decPGPPrivate" yaml:"decPGPPrivate" validate:"required,gt=0"`
	DecryptionPGPPrivateKeyPassphrase *string `json:"decPGPPrivatePassphrase" yaml:"decPGPPrivatePassphrase" validate:"required"`

	SFTPUsername string  `json:"SFTPUsername" yaml:"SFTPUsername" validate:"required,gt=0"`
	SFTPPassword *string `json:"SFTPPassword" yaml:"SFTPPassword" validate:"required,gt=0"`

	Environment TokenExEnvironment `json:"environment" yaml:"environment" validate:"oneof=sandbox production"`
	Region      TokenExRegion      `json:"region" yaml:"region" validate:"oneof=us eu"`
}

func (c *TokenExAccountUpdaterCredentials) GetMID() string {
	return c.SFTPUsername
}

func (c *TokenExAccountUpdaterCredentials) GetLibrary() Library {
	return LibraryTokenExAccountUpdater
}

func (c *TokenExAccountUpdaterCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeMethodUpdater}
}

func (c *TokenExAccountUpdaterCredentials) Validate() error {
	return nil
}

func (c *TokenExAccountUpdaterCredentials) GetSecureFields() []*string {
	return []*string{c.DecryptionPGPPrivateKey, c.SFTPPassword, c.DecryptionPGPPrivateKeyPassphrase}
}

func (c *TokenExAccountUpdaterCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *TokenExAccountUpdaterCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *TokenExAccountUpdaterCredentials) SupportsSca() bool {
	return false
}

func (c *TokenExAccountUpdaterCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *TokenExAccountUpdaterCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *TokenExAccountUpdaterCredentials) CanPlanModeUse(environment.Mode) bool {
	return true
}

func (c *TokenExAccountUpdaterCredentials) IsRecoveryAgent() bool {
	return false
}
