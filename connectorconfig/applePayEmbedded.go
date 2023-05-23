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
}

func (a *ApplePayEmbedded) GetSecureFields() []*string {
	if a == nil {
		return nil
	}
	return []*string{a.AppleMerchantPrivateKey, a.AppleMerchantCertificate}
}

func (a *ApplePayEmbedded) IsValid() bool {
	return a == nil ||
		a.ConnectorID != "" ||
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
