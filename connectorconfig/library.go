package connectorconfig

import (
	"errors"

	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type LibraryDef struct {
	DisplayName    string
	Credentials    Credentials
	SupportsMethod func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool
}

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
	LibraryNuvei                    Library = "nuvei"
	LibrarySandbanx                 Library = "sandbanx"

	// Fraud Libraries
	LibraryChargeHive  Library = "chargehive"
	LibraryMaxMind     Library = "maxmind"
	LibraryCyberSource Library = "cybersource"
	LibraryKount       Library = "kount"

	// Updater Libraries
	LibraryPaySafeAccountUpdater Library = "paysafe-accountupdater"
)

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
	lib, ok := LibraryRegister[l]
	if ok {
		return lib.DisplayName
	}
	return string(l)
}

func (l Library) GetCredential() (Credentials, error) {
	lib, ok := LibraryRegister[l]
	if ok {
		return lib.Credentials, nil
	}
	return nil, errors.New("invalid library specified")
}

func (l Library) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	lib, ok := LibraryRegister[l]
	if ok {
		return lib.SupportsMethod(methodType, methodProvider)
	}
	return false
}

var LibraryRegister = map[Library]LibraryDef{
	LibrarySandbox: {
		DisplayName: "ChargeHive Sandbox",
		Credentials: &SandboxCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryApplePay: {
		DisplayName: "ApplePay",
		Credentials: &ApplePayCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryAuthorize: {
		DisplayName: "Authorize",
		Credentials: &AuthorizeCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryBraintree: {
		DisplayName: "Braintree",
		Credentials: &BraintreeCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
		},
	},
	LibraryQualPay: {
		DisplayName: "QualPay",
		Credentials: &QualpayCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryStripe: {
		DisplayName: "Stripe",
		Credentials: &StripeCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryPaySafe: {
		DisplayName: "PaySafe",
		Credentials: &PaySafeCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return (methodType == chtype.PAYMENT_METHOD_TYPE_CARD) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
		},
	},
	LibraryPaySafeAccountUpdater: {
		DisplayName: "PaySafe Account Updater",
		Credentials: &PaySafeAccountUpdaterCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryWorldpay: {
		DisplayName: "Worldpay",
		Credentials: &WorldpayCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return (methodType == chtype.PAYMENT_METHOD_TYPE_CARD) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
		},
	},
	LibraryPayPalWebsitePaymentsPro: {
		DisplayName: "PayPal Website Payments Pro",
		Credentials: &PayPalWebsitePaymentsProCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryPayPalExpressCheckout: {
		DisplayName: "PayPal Express Checkout",
		Credentials: &PayPalExpressCheckoutCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_PAYPAL
		},
	},
	LibraryVindicia: {
		DisplayName: "Vindicia",
		Credentials: &VindiciaCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryBottomline: {
		DisplayName: "Bottomline",
		Credentials: &BottomlineCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_DIRECTDEBIT
		},
	},
	LibraryCheckout: {
		DisplayName: "Checkout.com",
		Credentials: &CheckoutCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
		},
	},
	LibraryChargeHive: {
		DisplayName: "ChargeHive",
		Credentials: &ChargeHiveCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},

	// Fraud Libraries
	LibraryMaxMind: {
		DisplayName: "MaxMind",
		Credentials: &MaxMindCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryCyberSource: {
		DisplayName: "CyberSource",
		Credentials: &CyberSourceCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryKount: {
		DisplayName: "Kount",
		Credentials: &KountCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryClearhaus: {
		DisplayName: "Clearhaus",
		Credentials: &ClearhausCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryTrustPayments: {
		DisplayName: "Trust Payments",
		Credentials: &TrustPaymentsCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryCWAMS: {
		DisplayName: "CWAMS",
		Credentials: &CWAMSCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return (methodType == chtype.PAYMENT_METHOD_TYPE_CARD) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
		},
	},
	LibraryYapstone: {
		DisplayName: "Yapstone",
		Credentials: &YapstoneCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibrarySandbanx: {
		DisplayName: "SandBanx",
		Credentials: &SandbanxCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryThreeDSecureIO: {
		DisplayName: "ThreeDSecureIO",
		Credentials: &ThreeDSecureIOCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryInovioPay: {
		DisplayName: "InovioPay",
		Credentials: &InovioPayCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryNuvei: {
		DisplayName: "Nuvei",
		Credentials: &NuveiCredentials{},
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
}
