package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type PaysafeApplePayInitiative string

const (
	// PaysafeApplePayInitiativeWeb - For Apple Pay on the web, use "web" for the initiative parameter.
	// For the initiativeContext parameter, provide your fully qualified domain name associated with your Apple Pay Merchant Identity Certificate.
	PaysafeApplePayInitiativeWeb PaysafeApplePayInitiative = "web"

	// PaysafeApplePayInitiativeMessaging - For Business Chat, use "messaging" for the initiative parameter.
	// For the initiativeContext parameter, pass your payment gateway URL. See Sending an Apple Pay Payment Request for more information.
	PaysafeApplePayInitiativeMessaging PaysafeApplePayInitiative = "messaging"
)

type PaySafeApplePayCredentials struct {
	Acquirer               string             `json:"acquirer" yaml:"acquirer"`
	AccountID              string             `json:"accountID" yaml:"accountID"`
	APIUsername            *string            `json:"apiUsername" yaml:"apiUsername"`
	APIPassword            *string            `json:"apiPassword" yaml:"apiPassword"`
	Environment            PaySafeEnvironment `json:"environment" yaml:"environment"`
	Country                string             `json:"country" yaml:"country"`
	Currency               string             `json:"currency" yaml:"currency"`
	SingleUseTokenUsername *string            `json:"singleUseTokenUsername" yaml:"singleUseTokenUsername"`
	SingleUseTokenPassword *string            `json:"singleUseTokenPassword" yaml:"singleUseTokenPassword"`
	Locale                 PaysafeLocale      `json:"locale" yaml:"locale"`

	ApplePayMerchantIdentityCert string `json:"applePayMerchantIdentityCert" yaml:"applePayMerchantIdentityCert"`
	ApplePayMerchantIdentityKey  string `json:"applePayMerchantIdentityKey" yaml:"applePayMerchantIdentityKey"`
	ApplePayMerchantIdentifier   string `json:"applePayMerchantIdentifier" yaml:"applePayMerchantIdentifier"`

	// On supported models of MacBook Pro, the Touch Bar displays the value you supply for the ApplePayDisplayName parameter.
	ApplePayDisplayName string `json:"applePayDisplayName" yaml:"applePayDisplayName"`
	ApplePayInitiative  string `json:"applePayInitiative" yaml:"applePayInitiative"`
	// Domain i.e cubernetes.io
	ApplePayInitiativeContext string `json:"applePayInitiativeContext" yaml:"applePayInitiativeContext"`
}

func (c PaySafeApplePayCredentials) GetLibrary() Library {
	return LibraryPaySafeApplePay
}

func (c *PaySafeApplePayCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypePayment}
}

func (c *PaySafeApplePayCredentials) Validate() error {
	return nil
}

func (c *PaySafeApplePayCredentials) GetSecureFields() []*string {
	return []*string{c.APIUsername, c.APIPassword, c.SingleUseTokenUsername, c.SingleUseTokenPassword}
}

func (c *PaySafeApplePayCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *PaySafeApplePayCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}
