package policy

import "testing"

func TestFraudPolicyRunsAt(t *testing.T) {
	single := FraudPolicy{CheckTime: FraudCheckTimeAuthSuccess}
	if !single.RunsAt(FraudCheckTimeAuthSuccess) || single.RunsAt(FraudCheckTimeCaptureSuccess) {
		t.Error("a single check time runs at that time only")
	}

	several := FraudPolicy{CheckTimes: []FraudCheckTime{FraudCheckTimeAuthSuccess, FraudCheckTimeCaptureDecline}}
	for _, at := range []FraudCheckTime{FraudCheckTimeAuthSuccess, FraudCheckTimeCaptureDecline} {
		if !several.RunsAt(at) {
			t.Errorf("expected the policy to run at %s", at)
		}
	}
	if several.RunsAt(FraudCheckTimeCaptureSuccess) {
		t.Error("a policy runs only at the times it lists")
	}
}
