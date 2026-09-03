package connectorconfig

import (
	"testing"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

func strPtr(s string) *string { return &s }

func TestShift4Credentials(t *testing.T) {
	c := &Shift4Credentials{
		MerchantID:   strPtr("MID123"),
		SignatureKey: strPtr("key"),
		Environment:  Shift4EnvironmentIntegration,
	}

	if c.GetLibrary() != LibraryShift4 {
		t.Errorf("library: %s", c.GetLibrary())
	}
	if c.GetMID() != "MID123" {
		t.Errorf("mid: %s", c.GetMID())
	}
	if !c.CanPlanModeUse(environment.ModeSandbox) {
		t.Error("integration credentials must be usable in sandbox plan mode")
	}
	if c.CanPlanModeUse(environment.ModeProduction) {
		t.Error("integration credentials must not be usable in production plan mode")
	}
	if c.SupportsSca() {
		t.Error("no built-in 3DS; external SCA is passed through")
	}
	if !c.SupportsMethod(chtype.PAYMENT_METHOD_TYPE_CARD, chtype.PAYMENT_METHOD_PROVIDER_INVALID) {
		t.Error("must support card payments")
	}
	// wallet support requires embedded wallet credentials
	if c.SupportsMethod(chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET, chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) {
		t.Error("apple pay unsupported without embedded ApplePay credentials")
	}
}

func TestShift4Library_Registered(t *testing.T) {
	def, ok := LibraryRegister[LibraryShift4]
	if !ok {
		t.Fatal("LibraryShift4 must be in LibraryRegister")
	}
	if def.DisplayName != "Shift4" {
		t.Errorf("display name: %s", def.DisplayName)
	}
	if _, isShift4 := def.Credentials().(*Shift4Credentials); !isShift4 {
		t.Error("credentials constructor must return *Shift4Credentials")
	}
	if !def.SupportsMethod(chtype.PAYMENT_METHOD_TYPE_CARD, chtype.PAYMENT_METHOD_PROVIDER_INVALID) {
		t.Error("library must support card")
	}
	if !def.SupportsMethod(chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET, chtype.PAYMENT_METHOD_PROVIDER_APPLEPAY) {
		t.Error("library must support apple pay")
	}
	if !def.SupportsMethod(chtype.PAYMENT_METHOD_TYPE_DIGITALWALLET, chtype.PAYMENT_METHOD_PROVIDER_GOOGLEPAY) {
		t.Error("library must support google pay")
	}
}
