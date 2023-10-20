package connectorconfig

type ApplePayEmbeddedCredential interface {
	GetApplePay() *ApplePayEmbedded
}

type ApplePayEmbedded struct {
	// ConnectorID The ID of the connector that provides the ApplePay service
	ConnectorID string `json:"applePayConnectorID,omitempty" yaml:"applePayConnectorID,omitempty" validate:"-"`

	// AppleMerchantIdentifier REQUIRED TO ENABLE APPLE PAY Merchant Identifier specified in the Apple Developer Merchant section
	AppleMerchantIdentifier string `json:"appleMerchantIdentifier,omitempty" yaml:"appleMerchantIdentifier,omitempty" validate:"-"`
	// AppleMerchantDisplayName Value to be displayed on the payment page
	AppleMerchantDisplayName string `json:"appleMerchantDisplayName,omitempty" yaml:"appleMerchantDisplayName,omitempty" validate:"required_with=AppleMerchantIdentifier"`
	// AppleMerchantCertificate Merchant certificate in the Apple Developer Merchant section (must be base64 encoded!)
	AppleMerchantCertificate *string `json:"appleMerchantCertificate" yaml:"appleMerchantCertificate" validate:"required_with=AppleMerchantIdentifier"`
	// AppleMerchantPrivateKey Merchant private key generated from the CSR in the Apple Developer Merchant section (must be base64 encoded!)
	AppleMerchantPrivateKey *string `json:"appleMerchantPrivateKey" yaml:"appleMerchantPrivateKey" validate:"required_with=AppleMerchantIdentifier"`
	// AppleSupportedNetworks Specifies which card networks will be accepted by ApplePay
	AppleSupportedNetworks []AppleSupportedNetwork `json:"appleSupportedNetworks,omitempty" yaml:"appleSupportedNetworks,omitempty" validate:"required_with=AppleMerchantIdentifier,dive,oneof=amex cartesBancaires chinaUnionPay discover eftpos electron elo interac jcb mada maestro masterCard privateLabel visa vPay"`
	// AppleMerchantCapabilities The payment capabilities supported by the merchant
	AppleMerchantCapabilities []AppleMerchantCapability `json:"appleMerchantCapabilities,omitempty" yaml:"appleMerchantCapabilities,omitempty" validate:"required_with=AppleMerchantIdentifier,dive,oneof=supports3DS supportsCredit supportsDebit supportsEMV"`

	// AppleExistingMethodRequired Chargehive will not use this connector if the customer does not have a apple payment method already saved
	AppleExistingMethodRequired bool `json:"appleExistingMethodRequired,omitempty" yaml:"appleExistingMethodRequired,omitempty" validate:"-"`
	// AppleExistingMethodReport Chargehive will request the existing payment method information from ApplePay
	AppleExistingMethodReport bool `json:"appleExistingMethodReport,omitempty" yaml:"appleExistingMethodReport,omitempty" validate:"-"`

	// AppleCardAllowDebit Allow customer to pay with debit card
	AppleCardAllowDebit bool `json:"appleCardAllowDebit,omitempty" yaml:"appleCardAllowDebit,omitempty" validate:"-"`
	// AppleCardAllowCredit Allow customer to pay with credit card
	AppleCardAllowCredit bool `json:"appleCardAllowCredit,omitempty" yaml:"appleCardAllowCredit,omitempty" validate:"-"`

	// AppleEmailRequired (emailRequired) Set to true to request an email address.
	AppleEmailRequired bool `json:"appleEmailRequired,omitempty" yaml:"appleEmailRequired,omitempty" validate:"-"`

	// AppleCardBillingAddressReq Set to true if you require a billing address. A billing address should only be requested if it's required to process the transaction
	AppleCardBillingAddressReq bool `json:"appleCardBillingAddressReq,omitempty" yaml:"appleCardBillingAddressReq,omitempty" validate:"-"`
	// AppleCardBillingPhoneReq  Set to true if a phone number is required to process the transaction.
	AppleCardBillingPhoneReq bool `json:"appleCardBillingPhoneReq,omitempty" yaml:"appleCardBillingPhoneReq,omitempty" validate:"-"`

	// AppleCardShippingAddressReq Set to true if you require a shipping address. A shipping address should only be requested if it's required to process the transaction
	AppleCardShippingAddressReq bool `json:"appleCardShippingAddressReq,omitempty" yaml:"appleCardShippingAddressReq,omitempty" validate:"-"`
	// AppleCardShippingPhoneReq  Set to true if a phone number is required to process the transaction.
	AppleCardShippingPhoneReq bool `json:"appleCardShippingPhoneReq,omitempty" yaml:"appleCardShippingPhoneReq,omitempty" validate:"-"`
}

func (a *ApplePayEmbedded) GetSecureFields() []*string {
	if a == nil {
		return nil
	}
	return []*string{a.AppleMerchantPrivateKey, a.AppleMerchantCertificate}
}

func (a *ApplePayEmbedded) IsValid() bool {
	if a == nil {
		return false
	}
	return a.ConnectorID != "" ||
		(a.GetAppleMerchantIdentifier() != "" &&
			a.GetAppleMerchantDisplayName() != "" &&
			a.GetAppleMerchantCertificate() != "" &&
			a.GetAppleMerchantPrivateKey() != "")
}

func (a *ApplePayEmbedded) GetAppleMerchantPrivateKey() string {
	if a == nil {
		return ""
	}
	if a.AppleMerchantPrivateKey == nil {
		return ""
	}
	return *a.AppleMerchantPrivateKey
}

func (a *ApplePayEmbedded) GetAppleMerchantCertificate() string {
	if a == nil {
		return ""
	}
	if a.AppleMerchantCertificate == nil {
		return ""
	}
	return *a.AppleMerchantCertificate
}

func (a *ApplePayEmbedded) GetAppleMerchantIdentifier() string {
	if a == nil {
		return ""
	}
	return a.AppleMerchantIdentifier
}

func (a *ApplePayEmbedded) GetAppleMerchantDisplayName() string {
	if a == nil {
		return ""
	}
	return a.AppleMerchantDisplayName
}

// GetAppleMerchantPublicKey
// Deprecated: use GetAppleMerchantCertificate instead
func (a *ApplePayEmbedded) GetAppleMerchantPublicKey() string {
	return a.GetAppleMerchantCertificate()
}

func (a *ApplePayEmbedded) GetAppleExistingMethodRequired() bool {
	if a == nil {
		return false
	}
	return a.AppleExistingMethodRequired
}
func (a *ApplePayEmbedded) GetAppleExistingMethodReport() bool {
	if a == nil {
		return false
	}
	return a.AppleExistingMethodReport
}
func (a *ApplePayEmbedded) GetAppleCardAllowDebit() bool {
	if a == nil {
		return false
	}
	return a.AppleCardAllowDebit
}
func (a *ApplePayEmbedded) GetAppleCardAllowCredit() bool {
	if a == nil {
		return false
	}
	return a.AppleCardAllowCredit
}
func (a *ApplePayEmbedded) GetAppleEmailRequired() bool {
	if a == nil {
		return false
	}
	return a.AppleEmailRequired
}
func (a *ApplePayEmbedded) GetAppleCardBillingAddressReq() bool {
	if a == nil {
		return false
	}
	return a.AppleCardBillingAddressReq
}
func (a *ApplePayEmbedded) GetAppleCardBillingPhoneReq() bool {
	if a == nil {
		return false
	}
	return a.AppleCardBillingPhoneReq
}
func (a *ApplePayEmbedded) GetAppleCardShippingAddressReq() bool {
	if a == nil {
		return false
	}
	return a.AppleCardShippingAddressReq
}
func (a *ApplePayEmbedded) GetAppleCardShippingPhoneReq() bool {
	if a == nil {
		return false
	}
	return a.AppleCardShippingPhoneReq
}
