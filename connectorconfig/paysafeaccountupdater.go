package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/proto/golang/chargehive/chtype"

	"github.com/chargehive/configuration/v1/connector"
)

type PaySafeAccountUpdaterCredentials struct {
	MerchantAccountNumber string  `json:"merchantAccountNumber" yaml:"merchantAccountNumber" validate:"required,gt=0"`
	StoreID               string  `json:"storeID" yaml:"storeID" validate:"required,gt=0"`
	StorePassword         *string `json:"storePassword" yaml:"storePassword" validate:"required,gt=0"`

	// Encryption PGP key provided TO you BY paysafe (note this is a seperate pair from DecryptionPGPPrivateKey.)
	EncryptionPGPPublicKey string `json:"encPGPPublic" yaml:"encPGPPublic" validate:"required,gt=0"`

	// The decryption PGP private key is generated BY you and the public key component you have provided TO paysafe
	DecryptionPGPPrivateKey           *string `json:"decPGPPrivate" yaml:"decPGPPrivate" validate:"required,gt=0"`
	DecryptionPGPPrivateKeyPassphrase *string `json:"decPGPPrivatePassphrase" yaml:"decPGPPrivatePassphrase" validate:"required"`

	// This is your localy generated SFTP private key generated BY you and you will have given the public component TO paysafe
	SFTPPrivateKey *string `json:"sftpPrivate" yaml:"sftpPrivate" validate:"required,gt=0"`

	SFTPUser string `json:"sftpUser" yaml:"sftpUser" validate:"required,gt=0"`
	SFTPHost string `json:"sftpHost" yaml:"sftpHost" validate:"required,gt=0"`
}

func (c PaySafeAccountUpdaterCredentials) GetLibrary() Library {
	return LibraryPaySafeAccountUpdater
}

func (c *PaySafeAccountUpdaterCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeMethodUpdater}
}

func (c *PaySafeAccountUpdaterCredentials) Validate() error {
	return nil
}

func (c *PaySafeAccountUpdaterCredentials) GetSecureFields() []*string {
	return []*string{c.StorePassword, c.DecryptionPGPPrivateKey, c.SFTPPrivateKey, c.DecryptionPGPPrivateKeyPassphrase}
}

func (c *PaySafeAccountUpdaterCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *PaySafeAccountUpdaterCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *PaySafeAccountUpdaterCredentials) SupportsSca() bool {
	return false
}

func (c PaySafeAccountUpdaterCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if methodType == chtype.PAYMENT_METHOD_TYPE_CARD {
		return true
	}
	return false
}
