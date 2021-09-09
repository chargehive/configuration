package connectorconfig

import (
	"errors"
)

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
	LibraryCheckout                 Library = "checkout"

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
	LibraryCheckout:                 true,
	LibraryChargeHive:               true,
	LibraryMaxMind:                  true,
	LibraryCyberSource:              true,
}

type LibraryType string

const (
	LibraryTypePayment       LibraryType = "payment"
	LibraryTypeFraud         LibraryType = "fraud"
	LibraryTypeMethodUpdater LibraryType = "methodUpdater"
	LibraryRecoveryAgent     LibraryType = "recoveryAgent"
)

var LibraryTypeRegister = map[LibraryType]bool{
	LibraryTypePayment:   true,
	LibraryTypeFraud:     true,
	LibraryRecoveryAgent: true,
}

func (l Library) GetDisplayName() string {
	switch l {
	case LibrarySandbox:
		return "Sandbox"
	case LibraryAuthorize:
		return "Authorize"
	case LibraryBraintree:
		return "Braintree"
	case LibraryQualPay:
		return "QualPay"
	case LibraryStripe:
		return "Stripe"
	case LibraryPaySafe:
		return "PaySafe"
	case LibraryPaySafeAccountUpdater:
		return "PaySafe Account Updater"
	case LibraryWorldpay:
		return "Worldpay"
	case LibraryPayPalWebsitePaymentsPro:
		return "PayPal Website Payments Pro"
	case LibraryPayPalExpressCheckout:
		return "PayPal Express Checkout"
	case LibraryVindicia:
		return "Vindicia"
	case LibraryBottomline:
		return "Bottomline"
	case LibraryCheckout:
		return "Checkout.com"
	case LibraryChargeHive:
		return "ChargeHive"
	case LibraryMaxMind:
		return "MaxMind"
	case LibraryCyberSource:
		return "CyberSource"
	}
	return string(l)
}

func (l Library) GetCredential() (Credentials, error) {
	switch l {
	// Payment Libraries
	case LibraryAuthorize:
		return &AuthorizeCredentials{}, nil
	case LibraryBraintree:
		return &BraintreeCredentials{}, nil
	case LibraryQualPay:
		return &QualpayCredentials{}, nil
	case LibraryStripe:
		return &StripeCredentials{}, nil
	case LibraryPaySafe:
		return &PaySafeCredentials{}, nil
	case LibraryPayPalExpressCheckout:
		return &PayPalExpressCheckoutCredentials{}, nil
	case LibraryPayPalWebsitePaymentsPro:
		return &PayPalWebsitePaymentsProCredentials{}, nil
	case LibraryWorldpay:
		return &WorldpayCredentials{}, nil
	case LibrarySandbox:
		return &SandboxCredentials{}, nil
	case LibraryVindicia:
		return &VindiciaCredentials{}, nil
	case LibraryBottomline:
		return &BottomlineCredentials{}, nil
	case LibraryCheckout:
		return &CheckoutCredentials{}, nil

		// Fraud Libraries
	case LibraryMaxMind:
		return &MaxMindCredentials{}, nil
	case LibraryCyberSource:
		return &CyberSourceCredentials{}, nil
	case LibraryChargeHive:
		return &ChargeHiveCredentials{}, nil

		// Updater libraries
	case LibraryPaySafeAccountUpdater:
		return &PaySafeAccountUpdaterCredentials{}, nil

	}

	return nil, errors.New("invalid library specified")
}
