package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/v1/connector"
)

type PaysafeApplePayLocale string

const (
	PaysafeApplePayLocaleENGB PaysafeApplePayLocale = "en_GB"
	PaysafeApplePayLocaleENUS PaysafeApplePayLocale = "en_US"
	PaysafeApplePayLocaleFRCA PaysafeApplePayLocale = "fr_CA"
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
	Acquirer               string
	AccountID              string
	APIUsername            *string
	APIPassword            *string
	Environment            PaySafeEnvironment
	Country                string
	Currency               string
	SingleUseTokenUsername *string
	SingleUseTokenPassword *string
	Locale                 PaysafeApplePayLocale

	ApplePayMerchantIdentityCert string
	ApplePayMerchantIdentityKey  string
	ApplePayMerchantIdentifier   string

	// On supported models of MacBook Pro, the Touch Bar displays the value you supply for the ApplePayDisplayName parameter.
	ApplePayDisplayName string
	ApplePayInitiative  string

	// Domain i.e cubernetes.io
	ApplePayInitiativeContext string
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
