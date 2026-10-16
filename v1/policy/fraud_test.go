package policy

import "testing"

func TestFraudPolicyTrigger(t *testing.T) {
	tests := []struct {
		name   string
		policy FraudPolicy
		want   FraudTrigger
	}{
		{"preauth-first", FraudPolicy{CheckTime: FraudCheckTimePreFirstAuth},
			FraudTrigger{FraudTransactionTypeAuth, FraudPositionPre, FraudAttemptsFirst, FraudModeScan, FraudModeScan}},
		{"preauth-every", FraudPolicy{CheckTime: FraudCheckTimePreEveryAuth},
			FraudTrigger{FraudTransactionTypeAuth, FraudPositionPre, FraudAttemptsEvery, FraudModeScan, FraudModeScan}},
		{"auth-success", FraudPolicy{CheckTime: FraudCheckTimeAuthSuccess},
			FraudTrigger{FraudTransactionTypeAuth, FraudPositionPostSuccess, FraudAttemptsEvery, FraudModeScan, FraudModeScan}},
		{"ondemand", FraudPolicy{CheckTime: FraudCheckTimeOnDemand},
			FraudTrigger{FraudTransactionTypeOnDemand, "", FraudAttemptsEvery, FraudModeScan, FraudModeScan}},
		{"set out in full", FraudPolicy{TransactionType: FraudTransactionTypeCapture, Position: FraudPositionPost, Attempts: FraudAttemptsFirst, CIT: FraudModeScan, MIT: FraudModeInform},
			FraudTrigger{FraudTransactionTypeCapture, FraudPositionPost, FraudAttemptsFirst, FraudModeScan, FraudModeInform}},
		{"defaults", FraudPolicy{TransactionType: FraudTransactionTypeAuth, Position: FraudPositionPre, CIT: FraudModeScan},
			FraudTrigger{FraudTransactionTypeAuth, FraudPositionPre, FraudAttemptsEvery, FraudModeScan, FraudModeNone}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.policy.Trigger(); got != tt.want {
				t.Errorf("expected %+v, got %+v", tt.want, got)
			}
		})
	}
}

func TestFraudTriggerRuns(t *testing.T) {
	post := FraudTrigger{TransactionType: FraudTransactionTypeCapture, Position: FraudPositionPost}
	if !post.Runs(FraudTransactionTypeCapture, FraudPositionPostSuccess) || !post.Runs(FraudTransactionTypeCapture, FraudPositionPostDecline) {
		t.Error("a post trigger runs on either outcome")
	}
	if post.Runs(FraudTransactionTypeCaptureAuth, FraudPositionPostSuccess) {
		t.Error("a trigger runs only at its own transaction type")
	}
	if post.Runs(FraudTransactionTypeCapture, FraudPositionPre) {
		t.Error("a post trigger does not run before the transaction")
	}
	declines := FraudTrigger{TransactionType: FraudTransactionTypeAuth, Position: FraudPositionPostDecline}
	if declines.Runs(FraudTransactionTypeAuth, FraudPositionPostSuccess) || !declines.Runs(FraudTransactionTypeAuth, FraudPositionPostDecline) {
		t.Error("a post-decline trigger runs on a decline only")
	}
}

func TestFraudTriggerRunsOnDemand(t *testing.T) {
	onDemand := FraudTrigger{TransactionType: FraudTransactionTypeOnDemand}
	if !onDemand.Runs(FraudTransactionTypeOnDemand, "") {
		t.Error("an on-demand trigger runs when a scan is requested")
	}
	if onDemand.Runs(FraudTransactionTypeAuth, FraudPositionPre) {
		t.Error("an on-demand trigger does not run at a payment")
	}
}

func TestFraudTriggerMode(t *testing.T) {
	tr := FraudTrigger{CIT: FraudModeScan, MIT: FraudModeInform}
	if tr.Mode(true) != FraudModeScan || tr.Mode(false) != FraudModeInform {
		t.Error("each initiator has its own mode")
	}
}
