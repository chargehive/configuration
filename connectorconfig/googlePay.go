package connectorconfig

type GooglePay struct {
	GoogleEnvironment              string   `json:"googleEnvironment,omitempty" yaml:"googleEnvironment,omitempty" validate:"required_with=GoogleMerchantId,omitempty,oneof=TEST PRODUCTION"`                        // (environment) PRODUCTION: Used to return chargeable payment methods when a valid Google merchant ID is specified and configured for the domain.TEST: Dummy payment methods that are suitable for testing (default).
	GoogleMerchantId               string   `json:"googleMerchantId,omitempty" yaml:"googleMerchantId,omitempty" validate:"-"`                                                                                       // (merchantInfo.merchantId) A Google merchant identifier issued after your website is approved by Google. Required when PaymentsClient is initialized with an environment property of PRODUCTION. See the Integration checklist for more information about the approval process and how to obtain a Google merchant identifier. (https://developers.google.com/pay/api/web/reference/request-objects#MerchantInfo)
	GoogleMerchantName             string   `json:"googleMerchantName,omitempty" yaml:"googleMerchantName,omitempty" validate:"required_with=GoogleMerchantId"`                                                      // (merchantInfo.merchantName) Merchant name encoded as UTF-8. Merchant name is rendered in the payment sheet. In TEST environment, or if a merchant isn't recognized, a “Pay Unverified Merchant” message is displayed in the payment sheet.
	GoogleExistingMethodRequired   bool     `json:"googleExistingMethodRequired,omitempty" yaml:"googleExistingMethodRequired,omitempty" validate:"-"`                                                               // (existingPaymentMethodRequired) If set to true then the IsReadyToPayResponse object includes an additional property that describes the visitor's readiness to pay with one or more payment methods specified in allowedPaymentMethods. (https://developers.google.com/pay/api/web/reference/request-objects#IsReadyToPayRequest)
	GoogleEmailReq                 bool     `json:"googleEmailReq,omitempty" yaml:"googleEmailReq,omitempty" validate:"-"`                                                                                           // (emailRequired) Set to true to request an email address. (https://developers.google.com/pay/api/web/reference/request-objects#PaymentDataRequest)
	GoogleAcceptCard               bool     `json:"googleAcceptCard,omitempty" yaml:"googleAcceptCard,omitempty" validate:"-"`                                                                                       // (Card {type = "CARD"}) Enable this to allow card payments through GooglePay
	GoogleCardAuthMethods          []string `json:"googleCardAuthMethods,omitempty" yaml:"googleCardAuthMethods,omitempty" validate:"required_with=GoogleMerchantId,dive,oneof=PAN_ONLY CRYPTOGRAM_3DS"`             // (Card {parameters.allowedAuthMethods}) Fields supported to authenticate a card transaction.
	GoogleCardNetworks             []string `json:"googleCardNetworks,omitempty" yaml:"googleCardNetworks,omitempty" validate:"required_with=GoogleMerchantId,dive,oneof=AMEX DISCOVER INTERAC JCB MASTERCARD VISA"` // (Card {parameters.allowedCardNetworks}) One or more card networks that you support, also supported by the Google Pay API.
	GoogleCardAllowPrepaid         bool     `json:"googleCardAllowPrepaid,omitempty" yaml:"googleCardAllowPrepaid,omitempty" validate:"-"`                                                                           // (Card {parameters.allowPrepaidCards}) Allow customer to pay with prepaid card
	GoogleCardAllowCredit          bool     `json:"googleCardAllowCredit,omitempty" yaml:"googleCardAllowCredit,omitempty" validate:"-"`                                                                             // (Card {parameters.allowCreditCards}) Allow customer to pay with credit card
	GoogleCardBillingAddressReq    bool     `json:"googleCardBillingAddressReq,omitempty" yaml:"googleCardBillingAddressReq,omitempty" validate:"-"`                                                                 // (Card {parameters.billingAddressRequired}) Set to true if you require a billing address. A billing address should only be requested if it's required to process the transaction
	GoogleCardBillingAddressFormat string   `json:"googleCardBillingAddressFormat,omitempty" yaml:"googleCardBillingAddressFormat,omitempty" validate:"required_with=GoogleMerchantId,omitempty,oneof=MIN FULL"`     // (Card {parameters.billingAddressParameters.format) Billing address format required to complete the transaction.
	GoogleCardBillingPhoneReq      bool     `json:"googleCardBillingPhoneReq,omitempty" yaml:"googleCardBillingPhoneReq,omitempty" validate:"-"`                                                                     // (Card {parameters.billingAddressParameters.phoneNumberRequired) Set to true if a phone number is required to process the transaction.
	GoogleCardTokenType            string   `json:"googleCardTokenType,omitempty" yaml:"googleCardTokenType,omitempty" validate:"required_with=GoogleMerchantId,omitempty,oneof=DIRECT PAYMENT_GATEWAY"`             // (Card {tokenizationSpecification.type})
	GoogleCardGateway              string   `json:"googleCardGateway,omitempty" yaml:"googleCardGateway,omitempty" validate:"required_with=GoogleMerchantId,omitempty,oneof=vantiv"`                                 // (Card {tokenizationSpecification.parameters.gateway}) https://developers.google.com/pay/api/web/reference/request-objects#gateway
	GoogleCardMerchantId           string   `json:"googleCardMerchantId,omitempty" yaml:"googleCardMerchantId,omitempty" validate:"required_with=GoogleMerchantId"`                                                  // (Card {tokenizationSpecification.parameters.gatewayMerchantId}) https://developers.google.com/pay/api/web/reference/request-objects#gateway
}

type (
	GoogleEnvironment           string
	GoogleCardGateway           string
	GoogleCardAuthMethod        string
	GoogleTokenType             string
	GoogleCardNetwork           string
	GoogleCardBillingAddressReq string
)

const (
	GoogleCardBillingAddressReqMIN  GoogleCardBillingAddressReq = "MIN"  // Name, country code, and postal code (default).
	GoogleCardBillingAddressReqFULL GoogleCardBillingAddressReq = "FULL" // Name, street address, locality, region, country code, and postal code.

	GoogleEnvironmentTEST GoogleEnvironment = "TEST"
	GoogleEnvironmentPROD GoogleEnvironment = "PRODUCTION"

	GoogleCardGatewayVANTIV GoogleCardGateway = "vantiv"

	GoogleCardTokenTypeDIRECT  GoogleTokenType = "DIRECT"
	GoogleCardTokenTypeGATEWAY GoogleTokenType = "PAYMENT_GATEWAY"

	GoogleCardAuthMethodPAN GoogleCardAuthMethod = "PAN_ONLY"       // This authentication method is associated with payment cards stored on file with the user's Google Account. Returned payment data includes personal account number (PAN) with the expiration month and the expiration year.
	GoogleCardAuthMethod3DS GoogleCardAuthMethod = "CRYPTOGRAM_3DS" // This authentication method is associated with cards stored as Android device tokens. Returned payment data includes a 3-D Secure (3DS) cryptogram generated on the device.

	GoogleCardNetworkAMEX       GoogleCardNetwork = "AMEX"
	GoogleCardNetworkDISCOVER   GoogleCardNetwork = "DISCOVER"
	GoogleCardNetworkINTERAC    GoogleCardNetwork = "INTERAC"
	GoogleCardNetworkJCB        GoogleCardNetwork = "JCB"
	GoogleCardNetworkMASTERCARD GoogleCardNetwork = "MASTERCARD"
	GoogleCardNetworkVISA       GoogleCardNetwork = "VISA"
)
