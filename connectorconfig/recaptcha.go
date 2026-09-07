package connectorconfig

import (
	"encoding/json"

	"github.com/chargehive/configuration/environment"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/configuration/v1/scheduler"
	"github.com/chargehive/proto/golang/chargehive/chtype"
)

// RecaptchaSuggestionSource names the reCAPTCHA score a suggestion range is matched
// against. The two scores answer different questions and run in opposite directions, so a
// range is only meaningful alongside the score it was written for.
//
// Kept in step with psp-configuration/recaptcha.SuggestionSource, which is where the
// connector reads it.
type RecaptchaSuggestionSource string

const (
	// RecaptchaSuggestionSourceRiskAnalysis matches riskAnalysis.score: how likely the
	// session is a human rather than automation. 1.0 is very likely legitimate, 0.0 very
	// likely a bot. This is the default, and what an unset source means.
	RecaptchaSuggestionSourceRiskAnalysis RecaptchaSuggestionSource = "recaptcha"

	// RecaptchaSuggestionSourceTransactionRisk matches
	// fraudPreventionAssessment.transactionRisk: how likely the payment is fraudulent.
	// 1.0 is very likely fraud, 0.0 very likely legitimate - the opposite direction to
	// riskAnalysis.
	RecaptchaSuggestionSourceTransactionRisk RecaptchaSuggestionSource = "transactionRisk"

	// RecaptchaSuggestionSourceCardTestingRisk matches
	// fraudPreventionAssessment.cardTestingVerdict.risk: how likely the attempt is part
	// of a card testing attack. 1.0 is highest risk.
	RecaptchaSuggestionSourceCardTestingRisk RecaptchaSuggestionSource = "cardTestingRisk"

	// RecaptchaSuggestionSourceStolenInstrumentRisk matches
	// fraudPreventionAssessment.stolenInstrumentVerdict.risk: how likely the instrument
	// is not the payer's. 1.0 is highest risk.
	RecaptchaSuggestionSourceStolenInstrumentRisk RecaptchaSuggestionSource = "stolenInstrumentRisk"

	// RecaptchaSuggestionSourceBehavioralTrust matches
	// fraudPreventionAssessment.behavioralTrustVerdict.trust: how trustworthily the
	// attempt behaved. 1.0 is MOST trustworthy - this runs with riskAnalysis, against the
	// three risk scores, so a deny range here sits at the bottom and not the top.
	RecaptchaSuggestionSourceBehavioralTrust RecaptchaSuggestionSource = "behavioralTrust"
)

type RecaptchaSuggestionRange struct {
	// Source selects which score Min and Max apply to. Empty means
	// RecaptchaSuggestionSourceRiskAnalysis, so ranges written before this field existed
	// keep working unchanged.
	Source RecaptchaSuggestionSource `json:"source,omitempty" yaml:"source,omitempty" validate:"omitempty,oneof=recaptcha transactionRisk cardTestingRisk stolenInstrumentRisk behavioralTrust"`
	Min    float32                   `json:"min" yaml:"min" validate:"required"`
	Max    float32                   `json:"max" yaml:"max" validate:"required"`
	Action string                    `json:"action" yaml:"action" validate:"required, oneof=review allow deny"`
}

// GetSource returns the source the range matches against, resolving the empty default.
func (r RecaptchaSuggestionRange) GetSource() RecaptchaSuggestionSource {
	if r.Source == "" {
		return RecaptchaSuggestionSourceRiskAnalysis
	}
	return r.Source
}

type RecaptchaCredentials struct {
	SiteKey     string                     `json:"siteKey" yaml:"siteKey" validate:"required"`
	ProjectID   string                     `json:"projectId" yaml:"projectId" validate:"required"`
	Suggestions []RecaptchaSuggestionRange `json:"suggestions" yaml:"suggestions"`
}

func (c *RecaptchaCredentials) GetLibrary() Library {
	return LibraryRecaptcha
}

func (c *RecaptchaCredentials) GetSupportedTypes() []LibraryType {
	return []LibraryType{LibraryTypeFraud}
}

func (c *RecaptchaCredentials) Validate() error {
	return nil
}

func (c *RecaptchaCredentials) GetSecureFields() []*string {
	return []*string{}
}

func (c *RecaptchaCredentials) ToConnector() connector.Connector {
	con := connector.Connector{Library: string(c.GetLibrary())}
	con.Configuration, _ = json.Marshal(c)
	return con
}

func (c *RecaptchaCredentials) FromJson(input []byte) error {
	return json.Unmarshal(input, c)
}

func (c *RecaptchaCredentials) SupportsSca() bool {
	return true
}

func (c *RecaptchaCredentials) SupportsMethod(methodType chtype.PaymentMethodType, methodProvider chtype.PaymentMethodProvider) bool {
	if !c.GetLibrary().SupportsMethod(methodType, methodProvider) {
		return false
	}
	return true
}

func (c *RecaptchaCredentials) SupportsCountry(country string) bool {
	return true
}

func (c *RecaptchaCredentials) CanPlanModeUse(mode environment.Mode) bool {
	return true
}

func (c *RecaptchaCredentials) IsRecoveryAgent() bool {
	return false
}

func (c *RecaptchaCredentials) Supports3RI() bool {
	return false
}

func (c *RecaptchaCredentials) IsAccountUpdater() bool {
	return false
}

func (c *RecaptchaCredentials) SupportedTokenTypes() []scheduler.TokenSource {
	return nil
}
