package connectorconfig

import "github.com/chargehive/configuration/v1/connector"

type Credentials interface {
	GetLibrary() Library
	GetSupportedTypes() []LibraryType
	GetSecureFields() []*string
	Validate() error
	ToConnector() connector.Connector
	FromJson(input []byte) error
}

type Library string

const (
	// Payment Libraries
	LibrarySandbox                  Library = "sandbox" // Connector for testing Charge hive
	LibraryAuthorize                Library = "authorize"
	LibraryBraintree                Library = "braintree"
	LibraryQualPay                  Library = "qualpay"
	LibraryStripe                   Library = "stripe"
	LibraryPaySafe                  Library = "paysafe"
	LibraryPaySafeApplePay          Library = "paysafe-applepay"
	LibraryPaySafeGooglePay         Library = "paysafe-googlepay"
	LibraryWorldpay                 Library = "worldpay"
	LibraryPayPalWebsitePaymentsPro Library = "paypal-websitepaymentspro"
	LibraryPayPalExpressCheckout    Library = "paypal-expresscheckout"
	LibraryVindicia                 Library = "vindicia"

	// Fraud Libraries
	LibraryChargeHive  Library = "chargehive"
	LibraryMaxMind     Library = "maxmind"
	LibraryCyberSource Library = "cybersource"
)

var LibraryRegister = map[Library]bool{
	LibrarySandbox:                  true,
	LibraryAuthorize:                true,
	LibraryBraintree:                true,
	LibraryQualPay:                  true,
	LibraryStripe:                   true,
	LibraryPaySafe:                  true,
	LibraryPaySafeApplePay:          true,
	LibraryPaySafeGooglePay:         true,
	LibraryWorldpay:                 true,
	LibraryPayPalWebsitePaymentsPro: true,
	LibraryPayPalExpressCheckout:    true,
	LibraryVindicia:                 true,
	LibraryChargeHive:               true,
	LibraryMaxMind:                  true,
	LibraryCyberSource:              true,
}

type LibraryType string

const (
	LibraryTypePayment LibraryType = "payment"
	LibraryTypeFraud   LibraryType = "fraud"
)

var LibraryTypeRegister = map[LibraryType]bool{
	LibraryTypePayment: true,
	LibraryTypeFraud:   true,
}
