package connectorconfig

import (
	"encoding/json"
	"testing"
)

// Connector configuration is stored by unmarshalling it into these structs and
// re-marshalling the result, so a field this struct does not carry is silently dropped on
// save. That is what makes the round trip, rather than the field itself, worth pinning.
func TestRecaptchaSuggestionSourceSurvivesRoundTrip(t *testing.T) {
	const stored = `{"siteKey":"sk","projectId":"pid","suggestions":[` +
		`{"min":0,"max":0.3,"action":"deny"},` +
		`{"source":"transactionRisk","min":0.8,"max":1,"action":"review"}]}`

	creds := &RecaptchaCredentials{}
	if err := creds.FromJson([]byte(stored)); err != nil {
		t.Fatalf("FromJson: %v", err)
	}

	if len(creds.Suggestions) != 2 {
		t.Fatalf("expected 2 suggestions, got %d", len(creds.Suggestions))
	}

	// An unset source keeps meaning riskAnalysis, so ranges written before the field
	// existed carry over unchanged.
	if got := creds.Suggestions[0].Source; got != "" {
		t.Errorf("expected an empty source, got %q", got)
	}
	if got := creds.Suggestions[0].GetSource(); got != RecaptchaSuggestionSourceRiskAnalysis {
		t.Errorf("expected %q, got %q", RecaptchaSuggestionSourceRiskAnalysis, got)
	}
	if got := creds.Suggestions[1].GetSource(); got != RecaptchaSuggestionSourceTransactionRisk {
		t.Errorf("expected %q, got %q", RecaptchaSuggestionSourceTransactionRisk, got)
	}

	var out RecaptchaCredentials
	if err := json.Unmarshal(creds.ToConnector().Configuration, &out); err != nil {
		t.Fatalf("unmarshal ToConnector output: %v", err)
	}
	if got := out.Suggestions[1].GetSource(); got != RecaptchaSuggestionSourceTransactionRisk {
		t.Errorf("source dropped on save: expected %q, got %q", RecaptchaSuggestionSourceTransactionRisk, got)
	}

	// omitempty keeps an unsourced range byte-identical to what merchants have stored.
	marshalled, err := json.Marshal(creds.Suggestions[0])
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if string(marshalled) != `{"min":0,"max":0.3,"action":"deny"}` {
		t.Errorf("unsourced range changed shape: %s", marshalled)
	}
}
