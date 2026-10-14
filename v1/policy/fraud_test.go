package policy

import "testing"

func TestFraudPolicyCheckTimeFor(t *testing.T) {
	tests := []struct {
		name               string
		policy             FraudPolicy
		customer, merchant FraudCheckTime
	}{
		{"every payment", FraudPolicy{CheckTime: FraudCheckTimePreEveryAuth}, FraudCheckTimePreEveryAuth, FraudCheckTimePreEveryAuth},
		{"merchant check", FraudPolicy{CheckTime: FraudCheckTimePreEveryAuth, MerchantCheck: true}, FraudCheckTimePreEveryAuth, FraudCheckTimePostPayment},
		{"auth-success runs after the auth", FraudPolicy{CheckTime: FraudCheckTimeAuthSuccess}, FraudCheckTimePostAuth, FraudCheckTimePostAuth},
		{"after the capture", FraudPolicy{CheckTime: FraudCheckTimePostCapture}, FraudCheckTimePostCapture, FraudCheckTimePostCapture},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.policy.CheckTimeFor(true); got != tt.customer {
				t.Errorf("customer-initiated: expected %q, got %q", tt.customer, got)
			}
			if got := tt.policy.CheckTimeFor(false); got != tt.merchant {
				t.Errorf("merchant-initiated: expected %q, got %q", tt.merchant, got)
			}
		})
	}
}
