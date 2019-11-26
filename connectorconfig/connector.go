package connectorconfig

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/v1/connector"
)

// CreateCredentials create credentials from json
func GetCredentials(c *connector.Connector) (Credentials, error) {
	switch Library(c.Library) {
	case LibraryAuthorize:
		creds := &AuthorizeCredentials{}
		err := json.Unmarshal(c.Configuration, creds)
		return creds, err
	case LibraryBraintree:
		creds := &BraintreeCredentials{}
		err := json.Unmarshal(c.Configuration, creds)
		return creds, err
	case LibraryQualPay:
		creds := &QualpayCredentials{}
		err := json.Unmarshal(c.Configuration, creds)
		return creds, err
	case LibraryStripe:
		creds := &StripeCredentials{}
		err := json.Unmarshal(c.Configuration, creds)
		return creds, err
	case LibraryPaySafe:
		creds := &PaySafeCredentials{}
		err := json.Unmarshal(c.Configuration, creds)
		return creds, err
	case LibraryPayPalWebsitePaymentsPro:
		creds := &PayPalWebsitePaymentsProCredentials{}
		err := json.Unmarshal(c.Configuration, creds)
		return creds, err
	case LibraryWorldpay:
		creds := &WorldpayCredentials{}
		err := json.Unmarshal(c.Configuration, creds)
		return creds, err
	case LibrarySandbox:
		creds := &SandboxCredentials{}
		err := json.Unmarshal(c.Configuration, creds)
		return creds, err
		// Fraud Libraries
	case LibraryMaxMind:
		creds := &MaxMindCredentials{}
		err := json.Unmarshal(c.Configuration, creds)
		return creds, err
	case LibraryCyberSource:
		creds := &CyberSourceCredentials{}
		err := json.Unmarshal(c.Configuration, creds)
		return creds, err
	case LibraryChargeHive:
		creds := &ChargeHiveCredentials{}
		err := json.Unmarshal(c.Configuration, creds)
		return creds, err
	default:
	}
	return nil, errors.New("invalid library specified")
}
