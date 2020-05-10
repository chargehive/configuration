package connectorconfig

type GooglePaySubCredentials struct {
	Environment                    GoogleEnvironment             `json:"environment" validate:"omitempty,oneof=TEST PRODUCTION"` // PRODUCTION: Used to return chargeable payment methods when a valid Google merchant ID is specified and configured for the domain.TEST: Dummy payment methods that are suitable for testing (default).
	MerchantId                     string                        `json:"merchantId" validate:"required"`                         // A Google merchant identifier issued after your website is approved by Google. Required when PaymentsClient is initialized with an environment property of PRODUCTION. See the Integration checklist for more information about the approval process and how to obtain a Google merchant identifier.
	MerchantName                   string                        `json:"merchantName" validate:"-"`                              // Merchant name encoded as UTF-8. Merchant name is rendered in the payment sheet. In TEST environment, or if a merchant isn't recognized, a “Pay Unverified Merchant” message is displayed in the payment sheet.
	ApiVersion                     int                           `json:"apiVersion" validate:"gte=0"`                            // Major API version. The value is 2 for this specification.
	ApiVersionMinor                int                           `json:"apiVersionMinor" validate:"gte=0"`                       // Minor API version. The value is 0 for this specification.
	ExistingPaymentMethodRequired  bool                          `json:"existingPaymentMethodRequired"`                          // If set to true then the IsReadyToPayResponse object includes an additional property that describes the visitor's readiness to pay with one or more payment methods specified in allowedPaymentMethods.
	AcceptCard                     bool                          `json:"acceptCard"`                                             // enable this to allow card payments through google pay
	CardAuthMethods                []CardAuthMethod              `json:"cardAuthMethods"`
	CardNetworks                   []CardNetwork                 `json:"cardNetworks"`
	CardTokenizationType           CardTokenizationType          `json:"cardTokenizationType"`
	CardGateway                    CardGateway                   `json:"cardGateway"`                    // https://developers.google.com/pay/api/web/reference/request-objects#gateway
	CardMerchantId                 string                        `json:"cardMerchantId"`                 // https://developers.google.com/pay/api/web/reference/request-objects#gateway
	CardAllowPrepaid               bool                          `json:"cardAllowPrepaid"`               // Allow customer to pay with prepaid card
	CardAllowCredit                bool                          `json:"cardAllowCredit"`                // Allow customer to pay with credit card
	CardBillingAddressRequirements CardBillingAddressRequirement `json:"cardBillingAddressRequirements"` // Set level of billing requirement
}

type (
	GoogleEnvironment             string
	CardGateway                   string
	CardAuthMethod                string
	CardTokenizationType          string
	CardNetwork                   string
	CardBillingAddressRequirement string
)

const (
	BillingAddressRequirementNone CardBillingAddressRequirement = ""
	BillingAddressRequirementMin  CardBillingAddressRequirement = "MIN"  // Name, country code, and postal code (default).
	BillingAddressRequirementFull CardBillingAddressRequirement = "FULL" // Name, street address, locality, region, country code, and postal code.

	GoogleEnvironmentTest GoogleEnvironment = "TEST"
	GoogleEnvironmentProd GoogleEnvironment = "PRODUCTION"

	CardTokenizationSpecGatewayWorldpay CardGateway = "worldpay"
	CardTokenizationSpecGatewayPaysafe  CardGateway = "paysafe"

	CardTokenizationSpecTypeDirect         CardTokenizationType = "DIRECT"
	CardTokenizationSpecTypePaymentGateway CardTokenizationType = "PAYMENT_GATEWAY"

	AuthMethodsPan CardAuthMethod = "PAN_ONLY"
	AuthMethods3ds CardAuthMethod = "CRYPTOGRAM_3DS"

	AllowedCardNetworksAMEX       CardNetwork = "AMEX"
	AllowedCardNetworksDISCOVER   CardNetwork = "DISCOVER"
	AllowedCardNetworksINTERAC    CardNetwork = "INTERAC"
	AllowedCardNetworksJCB        CardNetwork = "JCB"
	AllowedCardNetworksMASTERCARD CardNetwork = "MASTERCARD"
	AllowedCardNetworksVISA       CardNetwork = "VISA"
)
