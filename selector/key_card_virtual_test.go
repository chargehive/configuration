package selector

import "testing"

func TestCardVirtualKeyIsRegistered(t *testing.T) {
	if KeyMethodInfoCardVirtual != "method.info.card.virtual" {
		t.Errorf("KeyMethodInfoCardVirtual = %q", KeyMethodInfoCardVirtual)
	}
	if !KeyRegister[KeyMethodInfoCardVirtual] {
		t.Error("method.info.card.virtual is not in KeyRegister, so configs using it fail validation")
	}
}
