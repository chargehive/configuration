package policy

import "testing"

func TestFraudPolicyCheckTimeFor(t *testing.T) {
	tests := []struct {
		name               string
		policy             FraudPolicy
		customer, merchant FraudCheckTime
	}{
		{"per initiator", FraudPolicy{CustomerCheckTime: FraudCheckTimePreEveryAuth, MerchantCheckTime: FraudCheckTimePostCapture}, FraudCheckTimePreEveryAuth, FraudCheckTimePostCapture},
		{"customers only", FraudPolicy{CustomerCheckTime: FraudCheckTimePostAuth}, FraudCheckTimePostAuth, ""},
		{"merchants only", FraudPolicy{MerchantCheckTime: FraudCheckTimePostCapture}, "", FraudCheckTimePostCapture},
		{"check time for every payment", FraudPolicy{CheckTime: FraudCheckTimePreFirstAuth}, FraudCheckTimePreFirstAuth, FraudCheckTimePreFirstAuth},
		{"auth-success runs after the auth", FraudPolicy{CheckTime: FraudCheckTimeAuthSuccess}, FraudCheckTimePostAuth, FraudCheckTimePostAuth},
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
