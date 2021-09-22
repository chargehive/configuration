package connectorconfig

import (
	"errors"
	"github.com/chargehive/proto/golang/chargehive/chtype"
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
	LibraryKount       Library = "kount"

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
	LibraryKount:                    true,
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
		return "ChargeHive Sandbox"
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
	case LibraryKount:
		return "Kount"
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
	case LibraryKount:
		return &KountCredentials{}, nil

		// Updater libraries
	case LibraryPaySafeAccountUpdater:
		return &PaySafeAccountUpdaterCredentials{}, nil

	}

	return nil, errors.New("invalid library specified")
}

func (l Library) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	switch l {
	case LibrarySandbox:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryAuthorize:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryBraintree:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryQualPay:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryStripe:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryPaySafe:
		return (methodType == chtype.PAYMENT_METHOD_TYPE_CARD) ||
			(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
			(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
	case LibraryPaySafeAccountUpdater:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryWorldpay:
		return (methodType == chtype.PAYMENT_METHOD_TYPE_CARD) ||
			(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
			(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
	case LibraryPayPalWebsitePaymentsPro:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryPayPalExpressCheckout:
		return methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_PAYPAL
	case LibraryVindicia:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryBottomline:
		return methodType == chtype.PAYMENT_METHOD_TYPE_DIRECTDEBIT
	case LibraryCheckout:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryChargeHive:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryMaxMind:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryCyberSource:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryKount:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	}
	return false
}
