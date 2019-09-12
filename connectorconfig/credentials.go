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
	//Payment Libraries
	LibrarySandbox                  Library = "sandbox" //Connector for testing Charge hive
	LibraryAuthorize                Library = "authorize"
	LibraryBraintree                Library = "braintree"
	LibraryQualPay                  Library = "qualpay"
	LibraryStripe                   Library = "stripe"
	LibraryPaySafe                  Library = "paysafe"
	LibraryWorldpay                 Library = "worldpay"
	LibraryPayPalWebsitePaymentsPro Library = "paypal-websitepaymentspro"

	//Fraud Libraries
	LibraryChargeHive  Library = "chargehive"
	LibraryMaxMind     Library = "maxmind"
	LibraryCyberSource Library = "cybersource"
)

type LibraryType string

const (
	LibraryTypePayment LibraryType = "payment"
	LibraryTypeFraud   LibraryType = "fraud"
)
