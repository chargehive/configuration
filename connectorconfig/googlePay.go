package connectorconfig

type GooglePayCredential interface {
	GetGooglePay() *GooglePay
	GetGooglePayParams() map[string]string
}

type GooglePay struct {
	// GoogleMerchantId REQUIRED TO ENABLE GOOGLE PAY (merchantInfo.merchantId) A Google merchant identifier issued after your website is approved by Google. Required when PaymentsClient is initialized with an environment property of PRODUCTION. See the Integration checklist for more information about the approval process and how to obtain a Google merchant identifier. (https://developers.google.com/pay/api/web/reference/request-objects#MerchantInfo)
	GoogleMerchantId string `json:"googleMerchantId,omitempty" yaml:"googleMerchantId,omitempty" validate:"-"`
	// GoogleEnvironment (environment) PRODUCTION: Used to return chargeable payment methods when a valid Google merchant ID is specified and configured for the domain.TEST: Dummy payment methods that are suitable for testing (default).
	GoogleEnvironment GoogleEnvironment `json:"googleEnvironment,omitempty" yaml:"googleEnvironment,omitempty" validate:"required_with=GoogleMerchantId,omitempty,oneof=TEST PRODUCTION"`
	// GoogleMerchantName (merchantInfo.merchantName) Merchant name encoded as UTF-8. Merchant name is rendered in the payment sheet. In TEST environment, or if a merchant isn't recognized, a “Pay Unverified Merchant” message is displayed in the payment sheet.
	GoogleMerchantName string `json:"googleMerchantName,omitempty" yaml:"googleMerchantName,omitempty" validate:"required_with=GoogleMerchantId"`
	// GoogleExistingMethodRequired Chargehive will not use this connector if the customer does not have a google payment method already saved
	GoogleExistingMethodRequired bool `json:"googleExistingMethodRequired,omitempty" yaml:"googleExistingMethodRequired,omitempty" validate:"-"`
	// GoogleExistingMethodReport Chargehive will request the existing payment method information from GooglePay
	GoogleExistingMethodReport bool `json:"googleExistingMethodReport,omitempty" yaml:"googleExistingMethodReport,omitempty" validate:"-"`
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
	// GoogleCardBillingPhoneReq (Card {parameters.billingAddressParameters.phoneNumberRequired) Set to true if a phone number is required to process the transaction.
	GoogleCardBillingPhoneReq bool `json:"googleCardBillingPhoneReq,omitempty" yaml:"googleCardBillingPhoneReq,omitempty" validate:"-"`
	// GoogleCardBillingAddressFormat (Card {parameters.billingAddressParameters.format) Billing address format required to complete the transaction.
	GoogleCardBillingAddressFormat GoogleCardBillingAddressReq `json:"googleCardBillingAddressFormat,omitempty" yaml:"googleCardBillingAddressFormat,omitempty" validate:"required_with=GoogleMerchantId,omitempty,oneof=MIN FULL"`
	// GoogleCardTokenType (Card {tokenizationSpecification.type})
	GoogleCardTokenType GoogleTokenType `json:"googleCardTokenType,omitempty" yaml:"googleCardTokenType,omitempty" validate:"required_with=GoogleMerchantId,omitempty,oneof=DIRECT PAYMENT_GATEWAY"`
	// GoogleCardGateway (Card {tokenizationSpecification.parameters.gateway}) https://developers.google.com/pay/api/web/reference/request-objects#gateway
	GoogleCardGateway GoogleCardGateway `json:"googleCardGateway,omitempty" yaml:"googleCardGateway,omitempty" validate:"required_with=GoogleMerchantId,omitempty"`
	// GoogleCardMerchantId (Card {tokenizationSpecification.parameters.gatewayMerchantId}) https://developers.google.com/pay/api/web/reference/request-objects#gateway
	GoogleCardMerchantId string `json:"googleCardMerchantId,omitempty" yaml:"googleCardMerchantId,omitempty" validate:"required_with=GoogleMerchantId"`

	// GoogleCardShippingAddressReq Set to true if you require a shipping address. A shipping address should only be requested if it's required to process the transaction
	GoogleCardShippingAddressReq bool `json:"googleCardShippingAddressReq,omitempty" yaml:"googleCardShippingAddressReq,omitempty" validate:"-"`
	// GoogleCardShippingPhoneReq  Set to true if a phone number is required to process the transaction.
	GoogleCardShippingPhoneReq bool `json:"googleCardShippingPhoneReq,omitempty" yaml:"googleCardShippingPhoneReq,omitempty" validate:"-"`
	// GoogleCardShippingAddressFormat (Card {parameters.shippingAddressParameters.format) Shipping address format required to complete the transaction.
	GoogleCardShippingAddressFormat GoogleCardBillingAddressReq `json:"googleCardShippingAddressFormat,omitempty" yaml:"googleCardShippingAddressFormat,omitempty" validate:"omitempty,oneof=MIN FULL"`

	// GoogleAcquirerCountry The ISO 3166-1 alpha-2 country code where the transaction is processed. This property is required for merchants who process transactions in European Economic Area (EEA) countries and any other countries that are subject to Strong Customer Authentication (SCA). Merchants must specify the acquirer bank country code.
	GoogleAcquirerCountry string `json:"googleAcquirerCountry,omitempty" yaml:"googleAcquirerCountry,omitempty" validate:"omitempty,oneof=AF AX AL DZ AS AD AO AI AQ AG AR AM AW AU AT AZ BS BH BD BB BY BE BZ BJ BM BT BO BQ BA BW BV BR IO BN BG BF BI KH CM CA CV KY CF TD CL CN CX CC CO KM CG CD CK CR CI HR CU CW CY CZ DK DJ DM DO EC EG SV GQ ER EE ET FK FO FJ FI FR GF PF TF GA GM GE DE GH GI GR GL GD GP GU GT GG GN GW GY HT HM HN HK HU IS IN ID IR IQ IE IM IL IT JM JP JE JO KZ KE KI KP KR KW KG LA LV LB LS LR LY LI LT LU MO MK MG MW MY MV ML MT MH MQ MR MU YT MX FM MD MC MN ME MS MA MZ MM NA NR NP NC NZ NI NE NG NU NF MP NO OM PK PW PS PA PG PY PE PH PN PL PT PR QA RE RO RU RW BL SH KN LC MF VC WS SM ST SA SN RS SC SL SG SX SK SI SB SO ZA GS SS ES LK PM SD SR SJ SZ SE CH SY TW TJ TZ TH NL TL TG TK TO TT TN TR TM TC TV UG UA AE GB US UM UY UZ VU VA VE VN VG VI WF EH YE ZM ZW"`
}

func (g *GooglePay) GetSecureFields() []*string {
	return nil
}

func (g *GooglePay) IsValid() bool {
	if g == nil {
		return false
	}
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
func (g *GooglePay) GetGoogleExistingMethodReport() bool {
	if g == nil {
		return false
	}

	return g.GoogleExistingMethodReport
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
	return g.GoogleCardBillingPhoneReq
}
func (g *GooglePay) GetGoogleCardShippingAddressReq() bool {
	if g == nil {
		return false
	}
	return g.GoogleCardShippingAddressReq
}
func (g *GooglePay) GetGoogleCardShippingAddressFormat() GoogleCardBillingAddressReq {
	if g == nil {
		return ""
	}

	return g.GoogleCardShippingAddressFormat
}
func (g *GooglePay) GetGoogleCardShippingPhoneReq() bool {
	if g == nil {
		return false
	}
	return g.GoogleCardShippingPhoneReq
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
func (g *GooglePay) GetGoogleAcquirerCountry() string {
	if g == nil {
		return ""
	}
	return g.GoogleAcquirerCountry
}
