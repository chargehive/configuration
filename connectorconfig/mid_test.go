package connectorconfig

import (
	"testing"
)

// A connector that names a configuration carries no credentials, so reading its
// credentials decodes nothing into the library's struct. Every library's GetMID has to
// answer that struct: assemble reads the MID of whichever connector transacted, and a
// method that dereferences an unset pointer takes the process down with it.
func TestGetMIDAnswersAnEmptyCredentialsStructForEveryLibrary(t *testing.T) {
	for library, def := range LibraryRegister {
		t.Run(string(library), func(t *testing.T) {
			credentials := def.Credentials()
			midder, ok := credentials.(MerchantIdentifier)
			if !ok {
				return
			}
			defer func() {
				if recovered := recover(); recovered != nil {
					t.Fatalf("GetMID panicked on empty credentials: %v", recovered)
				}
			}()
			if got := midder.GetMID(); got != "" {
				t.Errorf("expected no MID from empty credentials, got %q", got)
			}
		})
	}
}

// An unset merchant id is no merchant id. "0" would be taken as one by anything that
// only ignores an empty value.
func TestQualpayGetMIDIsEmptyWhenUnset(t *testing.T) {
	if got := (&QualpayCredentials{}).GetMID(); got != "" {
		t.Errorf("expected an unset merchant id to read as empty, got %q", got)
	}
	if got := (&QualpayCredentials{MerchantID: 212000001}).GetMID(); got != "212000001" {
		t.Errorf("expected the merchant id, got %q", got)
	}
}

func TestGetMIDReadsTheConfiguredValue(t *testing.T) {
	for name, midder := range map[string]MerchantIdentifier{
		"bottomline":               &BottomlineCredentials{ClientID: strPtr("MID-1")},
		"epx":                      &EpxCredentials{MerchNbr: strPtr("MID-1")},
		"nuvei":                    &NuveiCredentials{MerchantID: strPtr("MID-1")},
		"paypalexpresscheckout":    &PayPalExpressCheckoutCredentials{APIUsername: strPtr("MID-1")},
		"paypalwebsitepaymentspro": &PayPalWebsitePaymentsProCredentials{APIUsername: strPtr("MID-1")},
		"trustpayments":            &TrustPaymentsCredentials{SiteRef: "MID-1"},
	} {
		if got := midder.GetMID(); got != "MID-1" {
			t.Errorf("%s: expected the configured value, got %q", name, got)
		}
	}
}

// A credential is not a merchant id. FlexPay carries nothing else, and its recovered
// transactions are recorded against the gateway connector they were prepared on; Bottomline
// and Trust Payments name the account by client id and site reference, not by login.
func TestGetMIDIsEmptyWhenOnlyCredentialsAreSet(t *testing.T) {
	for name, midder := range map[string]MerchantIdentifier{
		"bottomline":    &BottomlineCredentials{Username: strPtr("user"), Password: strPtr("secret"), ClientSUN: strPtr("123456")},
		"flexpay":       &FlexPayCredentials{ApiKey: strPtr("FLEXPAY-KEY")},
		"trustpayments": &TrustPaymentsCredentials{Username: strPtr("user@example.com")},
	} {
		if got := midder.GetMID(); got != "" {
			t.Errorf("%s: expected no merchant id from credentials alone, got %q", name, got)
		}
	}
}
