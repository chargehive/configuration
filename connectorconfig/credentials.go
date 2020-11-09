package connectorconfig

import (
	"github.com/LucidCube/chargehive-transport-config/plans"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type Credentials interface {
	GetLibrary() Library
	GetSupportedTypes() []LibraryType
	GetSecureFields() []*string
	Validate() error
	ToConnector() connector.Connector
	FromJson(input []byte) error
	SupportsSca() bool
	SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool
	CanPlanModeUse(mode plans.Mode) bool // Determine if this plan mode can use this configuration
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
	LibraryWorldpay                 Library = "worldpay"
	LibraryPayPalWebsitePaymentsPro Library = "paypal-websitepaymentspro"
	LibraryPayPalExpressCheckout    Library = "paypal-expresscheckout"
	LibraryVindicia                 Library = "vindicia"
	LibraryBottomline               Library = "bottomline"

	// Fraud Libraries
	LibraryChargeHive  Library = "chargehive"
	LibraryMaxMind     Library = "maxmind"
	LibraryCyberSource Library = "cybersource"

	// Updater Libraries
	LibraryPaySafeAccountUpdater Library = "paysafe-accountupdater"
)

var LibraryRegister = map[Library]bool{
	LibrarySandbox:                  true,
	LibraryAuthorize:                true,
	LibraryBraintree:                true,
	LibraryQualPay:                  true,
	LibraryStripe:                   true,
	LibraryPaySafe:                  true,
	LibraryPaySafeAccountUpdater:    true,
	LibraryWorldpay:                 true,
	LibraryPayPalWebsitePaymentsPro: true,
	LibraryPayPalExpressCheckout:    true,
	LibraryVindicia:                 true,
	LibraryBottomline:               true,
	LibraryChargeHive:               true,
	LibraryMaxMind:                  true,
	LibraryCyberSource:              true,
}

type LibraryType string

const (
	LibraryTypePayment       LibraryType = "payment"
	LibraryTypeFraud         LibraryType = "fraud"
	LibraryTypeMethodUpdater LibraryType = "methodUpdater"
)

var LibraryTypeRegister = map[LibraryType]bool{
	LibraryTypePayment: true,
	LibraryTypeFraud:   true,
}
