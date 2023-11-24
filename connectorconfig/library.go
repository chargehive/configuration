package connectorconfig

import (
	"errors"

	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type Library string

const (
	// Payment Libraries
	LibrarySandbox                  Library = "sandbox" // Connector for testing Charge hive
	LibraryApplePay                 Library = "applepay"
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
	LibraryClearhaus                Library = "clearhaus"
	LibraryTrustPayments            Library = "trust-payments"
	LibraryCWAMS                    Library = "cwams"
	LibraryYapstone                 Library = "yapstone"
	LibraryThreeDSecureIO           Library = "threedsecureio"
	LibraryInovioPay                Library = "inoviopay"
	LibrarySandbanx                 Library = "sandbanx"

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
	LibraryApplePay:                 true,
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
	LibraryClearhaus:                true,
	LibraryTrustPayments:            true,
	LibraryCWAMS:                    true,
	LibraryYapstone:                 true,
	LibrarySandbanx:                 true,
	LibraryThreeDSecureIO:           true,
	LibraryInovioPay:                true,
}

type LibraryType string

const (
	LibraryTypePayment        LibraryType = "payment"
	LibraryTypeFraud          LibraryType = "fraud"
	LibraryTypeMethodUpdater  LibraryType = "methodUpdater"
	LibraryTypeRecoveryAgent  LibraryType = "recoveryAgent"
	LibraryTypeAuthentication LibraryType = "authentication"
)

var LibraryTypeRegister = map[LibraryType]bool{
	LibraryTypePayment:        true,
	LibraryTypeFraud:          true,
	LibraryTypeRecoveryAgent:  true,
	LibraryTypeAuthentication: true,
}

func (l Library) GetDisplayName() string {
	switch l {
	case LibrarySandbox:
		return "ChargeHive Sandbox"
	case LibraryApplePay:
		return "ApplePay"
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
	case LibraryClearhaus:
		return "Clearhaus"
	case LibraryTrustPayments:
		return "Trust Payments"
	case LibraryCWAMS:
		return "CWAMS"
	case LibraryYapstone:
		return "Yapstone"
	case LibraryInovioPay:
		return "InovioPay"
	case LibrarySandbanx:
		return "SandBanx"
	}
	return string(l)
}

func (l Library) GetCredential() (Credentials, error) {
	switch l {
	// Payment Libraries
	case LibraryApplePay:
		return &ApplePayCredentials{}, nil
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
	case LibraryClearhaus:
		return &ClearhausCredentials{}, nil
	case LibraryTrustPayments:
		return &TrustPaymentsCredentials{}, nil
	case LibraryCWAMS:
		return &CWAMSCredentials{}, nil
	case LibraryYapstone:
		return &YapstoneCredentials{}, nil
	case LibrarySandbanx:
		return &SandbanxCredentials{}, nil

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

	case LibraryThreeDSecureIO:
		return &ThreeDSecureIOCredentials{}, nil
	case LibraryInovioPay:
		return &InovioPayCredentials{}, nil

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
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD ||
			(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
			(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
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
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD ||
			(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
			(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
	case LibraryChargeHive:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryMaxMind:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryCyberSource:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryKount:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryClearhaus:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryTrustPayments:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryCWAMS:
		return (methodType == chtype.PAYMENT_METHOD_TYPE_CARD) ||
			(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
			(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
	case LibraryYapstone:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryThreeDSecureIO:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibraryInovioPay:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	case LibrarySandbanx:
		return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
	}
	return false
}
