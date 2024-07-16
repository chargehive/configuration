package connectorconfig

import (
	"errors"

	"github.com/chargehive/proto/golang/chargehive/chtype"
)

type LibraryDef struct {
	DisplayName    string
	Credentials    func() Credentials
	SupportsMethod func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool
}

type Library string

const (
	LibraryNone Library = ""
	// Payment Libraries
	LibrarySandbox                  Library = "sandbox" // Connector for testing Charge hive
	LibraryAdyen                    Library = "adyen"
	LibraryApplePay                 Library = "applepay"
	LibraryGooglePay                Library = "googlepay"
	LibraryAuthorize                Library = "authorize"
	LibraryBraintree                Library = "braintree"
	LibraryBlueSnap                 Library = "bluesnap"
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
	LibraryGPayments                Library = "gpayments"
	LibrarySandbanx                 Library = "sandbanx"

	// Fraud Libraries
	LibraryChargeHive  Library = "chargehive"
	LibraryMaxMind     Library = "maxmind"
	LibraryCyberSource Library = "cybersource"
	LibraryKount       Library = "kount"

	// Updater Libraries
	LibraryPaySafeAccountUpdater Library = "paysafe-accountupdater"
	LibraryTokenExAccountUpdater Library = "tokenex-accountupdater"

	// Scheduler Libraries
	LibraryStickyIO Library = "sticky-io"

	// Tokenization Libraries
	LibraryTokenExNetworkTokenization Library = "tokenex-networktokenization"
)

type LibraryType string

const (
	LibraryTypePayment             LibraryType = "payment"
	LibraryTypeFraud               LibraryType = "fraud"
	LibraryTypeMethodUpdater       LibraryType = "methodUpdater"
	LibraryTypeRecoveryAgent       LibraryType = "recoveryAgent"
	LibraryTypeAuthentication      LibraryType = "authentication"
	LibraryTypeScheduler           LibraryType = "scheduler"
	LibraryTypeNetworkTokenization LibraryType = "networkTokenization"
)

var LibraryTypeRegister = map[LibraryType]bool{
	LibraryTypePayment:             true,
	LibraryTypeFraud:               true,
	LibraryTypeRecoveryAgent:       true,
	LibraryTypeAuthentication:      true,
	LibraryTypeScheduler:           true,
	LibraryTypeNetworkTokenization: true,
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
		return lib.Credentials(), nil
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
	LibraryNone: {
		DisplayName: "None",
		Credentials: func() Credentials { return &NoLibrary{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return false
		},
	},
	LibrarySandbox: {
		DisplayName: "ChargeHive Sandbox",
		Credentials: func() Credentials { return &SandboxCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryAdyen: {
		DisplayName: "Adyen",
		Credentials: func() Credentials { return &AdyenCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
		},
	},
	LibraryApplePay: {
		DisplayName: "ApplePay",
		Credentials: func() Credentials { return &ApplePayCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryGooglePay: {
		DisplayName: "GooglePayEmbedded",
		Credentials: func() Credentials { return &GooglePayCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryAuthorize: {
		DisplayName: "Authorize",
		Credentials: func() Credentials { return &AuthorizeCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryBlueSnap: {
		DisplayName: "BlueSnap",
		Credentials: func() Credentials { return &BlueSnapCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryBraintree: {
		DisplayName: "Braintree",
		Credentials: func() Credentials { return &BraintreeCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
		},
	},
	LibraryQualPay: {
		DisplayName: "QualPay",
		Credentials: func() Credentials { return &QualpayCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryStripe: {
		DisplayName: "Stripe",
		Credentials: func() Credentials { return &StripeCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryPaySafe: {
		DisplayName: "PaySafe",
		Credentials: func() Credentials { return &PaySafeCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return (methodType == chtype.PAYMENT_METHOD_TYPE_CARD) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
		},
	},
	LibraryPaySafeAccountUpdater: {
		DisplayName: "PaySafe Account Updater",
		Credentials: func() Credentials { return &PaySafeAccountUpdaterCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryWorldpay: {
		DisplayName: "Worldpay",
		Credentials: func() Credentials { return &WorldpayCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return (methodType == chtype.PAYMENT_METHOD_TYPE_CARD) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
		},
	},
	LibraryPayPalWebsitePaymentsPro: {
		DisplayName: "PayPal Website Payments Pro",
		Credentials: func() Credentials { return &PayPalWebsitePaymentsProCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryPayPalExpressCheckout: {
		DisplayName: "PayPal Express Checkout",
		Credentials: func() Credentials { return &PayPalExpressCheckoutCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_PAYPAL
		},
	},
	LibraryVindicia: {
		DisplayName: "Vindicia",
		Credentials: func() Credentials { return &VindiciaCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryBottomline: {
		DisplayName: "Bottomline",
		Credentials: func() Credentials { return &BottomlineCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_DIRECTDEBIT
		},
	},
	LibraryCheckout: {
		DisplayName: "Checkout.com",
		Credentials: func() Credentials { return &CheckoutCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
		},
	},
	// Fraud Libraries
	LibraryMaxMind: {
		DisplayName: "MaxMind",
		Credentials: func() Credentials { return &MaxMindCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryCyberSource: {
		DisplayName: "CyberSource",
		Credentials: func() Credentials { return &CyberSourceCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryKount: {
		DisplayName: "Kount",
		Credentials: func() Credentials { return &KountCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryClearhaus: {
		DisplayName: "Clearhaus",
		Credentials: func() Credentials { return &ClearhausCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryTrustPayments: {
		DisplayName: "Trust Payments",
		Credentials: func() Credentials { return &TrustPaymentsCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryCWAMS: {
		DisplayName: "CWAMS",
		Credentials: func() Credentials { return &CWAMSCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return (methodType == chtype.PAYMENT_METHOD_TYPE_CARD) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) ||
				(methodType == chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET && methodProvider == chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY)
		},
	},
	LibraryYapstone: {
		DisplayName: "Yapstone",
		Credentials: func() Credentials { return &YapstoneCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibrarySandbanx: {
		DisplayName: "SandBanx",
		Credentials: func() Credentials { return &SandbanxCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryThreeDSecureIO: {
		DisplayName: "ThreeDSecureIO",
		Credentials: func() Credentials { return &ThreeDSecureIOCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryInovioPay: {
		DisplayName: "InovioPay",
		Credentials: func() Credentials { return &InovioPayCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryNuvei: {
		DisplayName: "Nuvei",
		Credentials: func() Credentials { return &NuveiCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryGPayments: {
		DisplayName: "GPayments",
		Credentials: func() Credentials { return &GPaymentsCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryTokenExAccountUpdater: {
		DisplayName: "TokenEx Account Updater",
		Credentials: func() Credentials { return &TokenExAccountUpdaterCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
	LibraryTokenExNetworkTokenization: {
		DisplayName: "TokenEx Network Tokenization",
		Credentials: func() Credentials { return &TokenExNetworkTokenizationCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},

	LibraryStickyIO: {
		DisplayName: "Sticky IO",
		Credentials: func() Credentials { return &StickyIOCredentials{} },
		SupportsMethod: func(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
			return methodType == chtype.PAYMENT_METHOD_TYPE_CARD
		},
	},
}
