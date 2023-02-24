package connectorconfig

type GooglePayCredential interface {
	GetGooglePay() *GooglePay
}

type GooglePay struct {
	// GoogleMerchantId REQUIRED TO ENABLE GOOGLE PAY (merchantInfo.merchantId) A Google merchant identifier issued after your website is approved by Google. Required when PaymentsClient is initialized with an environment property of PRODUCTION. See the Integration checklist for more information about the approval process and how to obtain a Google merchant identifier. (https://developers.google.com/pay/api/web/reference/request-objects#MerchantInfo)
	GoogleMerchantId string `json:"googleMerchantId,omitempty" yaml:"googleMerchantId,omitempty" validate:"-"`
	// GoogleEnvironment (environment) PRODUCTION: Used to return chargeable payment methods when a valid Google merchant ID is specified and configured for the domain.TEST: Dummy payment methods that are suitable for testing (default).
	GoogleEnvironment GoogleEnvironment `json:"googleEnvironment,omitempty" yaml:"googleEnvironment,omitempty" validate:"required_with=GoogleMerchantId,omitempty,oneof=TEST PRODUCTION"`
	// GoogleMerchantName (merchantInfo.merchantName) Merchant name encoded as UTF-8. Merchant name is rendered in the payment sheet. In TEST environment, or if a merchant isn't recognized, a “Pay Unverified Merchant” message is displayed in the payment sheet.
	GoogleMerchantName string `json:"googleMerchantName,omitempty" yaml:"googleMerchantName,omitempty" validate:"required_with=GoogleMerchantId"`
	// GoogleExistingMethodRequired (existingPaymentMethodRequired) If set to true then the IsReadyToPayResponse object includes an additional property that describes the visitor's readiness to pay with one or more payment methods specified in allowedPaymentMethods. (https://developers.google.com/pay/api/web/reference/request-objects#IsReadyToPayRequest)
	GoogleExistingMethodRequired bool `json:"googleExistingMethodRequired,omitempty" yaml:"googleExistingMethodRequired,omitempty" validate:"-"`
	// GoogleEmailReq (emailRequired) Set to true to request an email address. (https://developers.google.com/pay/api/web/reference/request-objects#PaymentDataRequest)
	GoogleEmailReq bool `json:"googleEmailReq,omitempty" yaml:"googleEmailReq,omitempty" validate:"-"`
	// GoogleAcceptCard (Card {type = "CARD"}) Enable this to allow card payments through GooglePay
	GoogleAcceptCard bool `json:"googleAcceptCard,omitempty" yaml:"googleAcceptCard,omitempty" validate:"-"`
	// GoogleCardAuthMethods (Card {parameters.allowedAuthMethods}) Fields supported to authenticate a card transaction.
	GoogleCardAuthMethods []GoogleCardAuthMethod `json:"googleCardAuthMethods,omitempty" yaml:"googleCardAuthMethods,omitempty" validate:"required_with=GoogleMerchantId,dive,oneof=PAN_ONLY CRYPTOGRAM_3DS"`
	// GoogleCardNetworks (Card {parameters.allowedCardNetworks}) One or more card networks that you support, also supported by the Google Pay API.
	GoogleCardNetworks []GoogleCardNetwork `json:"googleCardNetworks,omitempty" yaml:"googleCardNetworks,omitempty" validate:"required_with=GoogleMerchantId,dive,oneof=AMEX DISCOVER INTERAC JCB MASTERCARD VISA"`
	// GoogleCardAllowPrepaid (Card {parameters.allowPrepaidCards}) Allow customer to pay with prepaid card
	GoogleCardAllowPrepaid bool `json:"googleCardAllowPrepaid,omitempty" yaml:"googleCardAllowPrepaid,omitempty" validate:"-"`
	// GoogleCardAllowCredit (Card {parameters.allowCreditCards}) Allow customer to pay with credit card
	GoogleCardAllowCredit bool `json:"googleCardAllowCredit,omitempty" yaml:"googleCardAllowCredit,omitempty" validate:"-"`
	// GoogleCardBillingAddressReq (Card {parameters.billingAddressRequired}) Set to true if you require a billing address. A billing address should only be requested if it's required to process the transaction
	GoogleCardBillingAddressReq bool `json:"googleCardBillingAddressReq,omitempty" yaml:"googleCardBillingAddressReq,omitempty" validate:"-"`
	// GoogleCardBillingAddressFormat (Card {parameters.billingAddressParameters.format) Billing address format required to complete the transaction.
	GoogleCardBillingAddressFormat GoogleCardBillingAddressReq `json:"googleCardBillingAddressFormat,omitempty" yaml:"googleCardBillingAddressFormat,omitempty" validate:"required_with=GoogleMerchantId,omitempty,oneof=MIN FULL"`
	// GoogleCardBillingPhoneReq (Card {parameters.billingAddressParameters.phoneNumberRequired) Set to true if a phone number is required to process the transaction.
	GoogleCardBillingPhoneReq bool `json:"googleCardBillingPhoneReq,omitempty" yaml:"googleCardBillingPhoneReq,omitempty" validate:"-"`
	// GoogleCardTokenType (Card {tokenizationSpecification.type})
	GoogleCardTokenType GoogleTokenType `json:"googleCardTokenType,omitempty" yaml:"googleCardTokenType,omitempty" validate:"required_with=GoogleMerchantId,omitempty,oneof=DIRECT PAYMENT_GATEWAY"`
	// GoogleCardGateway (Card {tokenizationSpecification.parameters.gateway}) https://developers.google.com/pay/api/web/reference/request-objects#gateway
	GoogleCardGateway GoogleCardGateway `json:"googleCardGateway,omitempty" yaml:"googleCardGateway,omitempty" validate:"required_with=GoogleMerchantId,omitempty"`
	// GoogleCardMerchantId (Card {tokenizationSpecification.parameters.gatewayMerchantId}) https://developers.google.com/pay/api/web/reference/request-objects#gateway
	GoogleCardMerchantId string `json:"googleCardMerchantId,omitempty" yaml:"googleCardMerchantId,omitempty" validate:"required_with=GoogleMerchantId"`
}

func (g *GooglePay) IsValid() bool {
	return g.GetGoogleMerchantId() != "" &&
		g.GetGoogleCardGateway() != "" &&
		g.GetGoogleCardMerchantId() != ""
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

	// GoogleCardGatewayVANTIV
	// Deprecated
	GoogleCardGatewayVANTIV GoogleCardGateway = "vantiv"
	// GoogleCardGatewayPAYSAFE
	// Deprecated
	GoogleCardGatewayPAYSAFE GoogleCardGateway = "paysafe"

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

func (g *GooglePay) GetGoogleMerchantId() string {
	if g == nil {
		return ""
	}
	return g.GoogleMerchantId
}
func (g *GooglePay) GetGoogleEnvironment() GoogleEnvironment {
	if g == nil {
		return ""
	}

	return g.GoogleEnvironment
}
func (g *GooglePay) GetGoogleMerchantName() string {
	if g == nil {
		return ""
	}

	return g.GoogleMerchantName
}
func (g *GooglePay) GetGoogleExistingMethodRequired() bool {
	if g == nil {
		return false
	}

	return g.GoogleExistingMethodRequired
}
func (g *GooglePay) GetGoogleEmailReq() bool {
	if g == nil {
		return false
	}

	return g.GoogleEmailReq
}
func (g *GooglePay) GetGoogleAcceptCard() bool {
	if g == nil {
		return false
	}

	return g.GoogleAcceptCard
}
func (g *GooglePay) GetGoogleCardAuthMethods() []GoogleCardAuthMethod {
	if g == nil {
		return nil
	}
	return g.GoogleCardAuthMethods
}
func (g *GooglePay) GetGoogleCardNetworks() []GoogleCardNetwork {
	if g == nil {
		return nil
	}

	return g.GoogleCardNetworks
}
func (g *GooglePay) GetGoogleCardAllowPrepaid() bool {
	if g == nil {
		return false
	}

	return g.GoogleCardAllowPrepaid
}
func (g *GooglePay) GetGoogleCardAllowCredit() bool {
	if g == nil {
		return false
	}

	return g.GoogleCardAllowCredit
}
func (g *GooglePay) GetGoogleCardBillingAddressReq() bool {
	if g == nil {
		return false
	}
	return g.GoogleCardBillingAddressReq
}
func (g *GooglePay) GetGoogleCardBillingAddressFormat() GoogleCardBillingAddressReq {
	if g == nil {
		return ""
	}

	return g.GoogleCardBillingAddressFormat
}
func (g *GooglePay) GetGoogleCardBillingPhoneReq() bool {
	if g == nil {
		return false
	}
	return g.GoogleCardBillingAddressReq
}
func (g *GooglePay) GetGoogleCardTokenType() GoogleTokenType {
	if g == nil {
		return ""
	}
	return g.GoogleCardTokenType
}
func (g *GooglePay) GetGoogleCardGateway() GoogleCardGateway {
	if g == nil {
		return ""
	}
	return g.GoogleCardGateway
}
func (g *GooglePay) GetGoogleCardMerchantId() string {
	if g == nil {
		return ""
	}
	return g.GoogleCardMerchantId
}
