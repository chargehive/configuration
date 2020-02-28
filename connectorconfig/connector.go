package connectorconfig

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/v1/connector"
)

// CreateCredentials create credentials from json
func GetCredentials(c *connector.Connector) (Credentials, error) {
	var credentials Credentials
	switch Library(c.Library) {

	// Payment Libraries
	case LibraryAuthorize:
		credentials = &AuthorizeCredentials{}
	case LibraryBraintree:
		credentials = &BraintreeCredentials{}
	case LibraryQualPay:
		credentials = &QualpayCredentials{}
	case LibraryStripe:
		credentials = &StripeCredentials{}
	case LibraryPaySafe:
		credentials = &PaySafeCredentials{}
	case LibraryPaySafeApplePay:
		credentials = &PaySafeApplePayCredentials{}
	case LibraryPaySafeGooglePay:
		credentials = &PaySafeGooglePayCredentials{}
	case LibraryPayPalExpressCheckout:
		credentials = &PayPalExpressCheckoutCredentials{}
	case LibraryPayPalWebsitePaymentsPro:
		credentials = &PayPalWebsitePaymentsProCredentials{}
	case LibraryWorldpay:
		credentials = &WorldpayCredentials{}
	case LibrarySandbox:
		credentials = &SandboxCredentials{}
	case LibraryVindicia:
		credentials = &VindiciaCredentials{}

		// Fraud Libraries
	case LibraryMaxMind:
		credentials = &MaxMindCredentials{}
	case LibraryCyberSource:
		credentials = &CyberSourceCredentials{}
	case LibraryChargeHive:
		credentials = &ChargeHiveCredentials{}
	default:
		return nil, errors.New("invalid library specified")
	}
	err := json.Unmarshal(c.Configuration, credentials)
	return credentials, err
}
