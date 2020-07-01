package connectorconfig

type GooglePay struct {
	GooglePayPageId              string                      `json:"googlePayPageId" validate:"-"`
	GoogleEnvironment            GoogleEnvironment           `json:"googleEnvironment" validate:"-"`     // PRODUCTION: Used to return chargeable payment methods when a valid Google merchant ID is specified and configured for the domain.TEST: Dummy payment methods that are suitable for testing (default).
	GoogleMerchantId             string                      `json:"googleMerchantId" validate:"-"`      // A Google merchant identifier issued after your website is approved by Google. Required when PaymentsClient is initialized with an environment property of PRODUCTION. See the Integration checklist for more information about the approval process and how to obtain a Google merchant identifier.
	GoogleMerchantName           string                      `json:"googleMerchantName" validate:"-"`    // Merchant name encoded as UTF-8. Merchant name is rendered in the payment sheet. In TEST environment, or if a merchant isn't recognized, a “Pay Unverified Merchant” message is displayed in the payment sheet.
	GoogleApiVersion             int                         `json:"googleApiVersion" validate:"-"`      // Major API version. The value is 2 for this specification.
	GoogleApiVersionMinor        int                         `json:"googleApiVersionMinor" validate:"-"` // Minor API version. The value is 0 for this specification.
	GoogleExistingMethodRequired bool                        `json:"googleExistingMethodRequired"`       // If set to true then the IsReadyToPayResponse object includes an additional property that describes the visitor's readiness to pay with one or more payment methods specified in allowedPaymentMethods.
	GoogleAcceptCard             bool                        `json:"googleAcceptCard"`                   // enable this to allow card payments through google pay
	GoogleCardAuthMethods        []GoogleCardAuthMethod      `json:"googleCardAuthMethods"`
	GoogleCardNetworks           []GoogleCardNetwork         `json:"googleCardNetworks"`
	GoogleCardTokenType          GoogleCardTokenType         `json:"googleCardTokenType"`
	GoogleCardGateway            GoogleCardGateway           `json:"googleCardGateway"`           // https://developers.google.com/pay/api/web/reference/request-objects#gateway
	GoogleCardMerchantId         string                      `json:"googleCardMerchantId"`        // https://developers.google.com/pay/api/web/reference/request-objects#gateway
	GoogleCardAllowPrepaid       bool                        `json:"googleCardAllowPrepaid"`      // Allow customer to pay with prepaid card
	GoogleCardAllowCredit        bool                        `json:"googleCardAllowCredit"`       // Allow customer to pay with credit card
	GoogleCardBillingAddressReq  GoogleCardBillingAddressReq `json:"googleCardBillingAddressReq"` // Set level of billing requirement
}

type (
	GoogleEnvironment           string
	GoogleCardGateway           string
	GoogleCardAuthMethod        string
	GoogleCardTokenType         string
	GoogleCardNetwork           string
	GoogleCardBillingAddressReq string
)

const (
	GoogleCardBillingAddressReqNONE GoogleCardBillingAddressReq = ""
	GoogleCardBillingAddressReqMIN  GoogleCardBillingAddressReq = "MIN"  // Name, country code, and postal code (default).
	GoogleCardBillingAddressReqFULL GoogleCardBillingAddressReq = "FULL" // Name, street address, locality, region, country code, and postal code.

	GoogleEnvironmentTEST GoogleEnvironment = "TEST"
	GoogleEnvironmentPROD GoogleEnvironment = "PRODUCTION"

	GoogleCardGatewayWORLDPAY GoogleCardGateway = "worldpay"
	GoogleCardGatewayPAYSAFE  GoogleCardGateway = "paysafe"

	GoogleCardTokenTypeDIRECT  GoogleCardTokenType = "DIRECT"
	GoogleCardTokenTypeGATEWAY GoogleCardTokenType = "PAYMENT_GATEWAY"

	GoogleCardAuthMethodPAN GoogleCardAuthMethod = "PAN_ONLY"
	GoogleCardAuthMethod3DS GoogleCardAuthMethod = "CRYPTOGRAM_3DS"

	GoogleCardNetworkAMEX       GoogleCardNetwork = "AMEX"
	GoogleCardNetworkDISCOVER   GoogleCardNetwork = "DISCOVER"
	GoogleCardNetworkINTERAC    GoogleCardNetwork = "INTERAC"
	GoogleCardNetworkJCB        GoogleCardNetwork = "JCB"
	GoogleCardNetworkMASTERCARD GoogleCardNetwork = "MASTERCARD"
	GoogleCardNetworkVISA       GoogleCardNetwork = "VISA"
)
