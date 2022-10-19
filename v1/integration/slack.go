package integration

import (
	"encoding/json"
	"errors"

	"github.com/chargehive/configuration/object"
)

// KindSlack indicates that the configuration is a slack integration
const KindSlack object.Kind = "Integration.Slack"

// Slack configuration structure
type Slack struct {
	// AccessToken slack access token
	AccessToken string `json:"accessToken" yaml:"accessToken" validate:"required"`

	// Scopes for OAuth authentication
	Scopes []string `json:"scopes" yaml:"scopes" validate:"-"`

	// TeamName for posting
	TeamName string `json:"teamName" yaml:"teamName" validate:"required"`

	// TeamID for posting
	TeamID string `json:"teamID" yaml:"teamID" validate:"required"`

	// Webhook endpoint
	Webhook *SlackWebhook `json:"webhook" yaml:"webhook" validate:"required"`

	// TransactionNotifications indicates the action to perform
	TransactionNotifications *bool `json:"transactionNotifications" yaml:"transactionNotifications" validate:"required"`
}

func (s Slack) GetTransactionNotifications() bool {
	if s.TransactionNotifications == nil {
		return true
	}
	return *s.TransactionNotifications
}

// SlackWebhook structure
type SlackWebhook struct {
	// Url is the slack webhook URL
	Url string `json:"url" yaml:"url" validate:"required"`

	// Channel is the slack channel to post in
	Channel string `json:"channel" yaml:"channel" validate:"required"`

	// ConfigurationUrl is the slack endpoint for configuration
	ConfigurationUrl string `json:"configurationUrl" yaml:"configurationUrl" validate:"required"`
}

// GetKind returns the Slack kind
func (Slack) GetKind() object.Kind { return KindSlack }

// GetVersion returns the Slack version
func (Slack) GetVersion() string { return "v1" }

// NewSlackDefinition returns a new slack config definition
func NewSlackDefinition(d *object.Definition) (*SlackDefinition, error) {
	if _, ok := d.Spec.(*Slack); ok {
		return &SlackDefinition{def: d}, nil
	}
	return nil, errors.New("invalid slack configuration")
}

// SlackDefinition is the full slack definition config structure
type SlackDefinition struct{ def *object.Definition }

// Definition returns the slack config definition
func (d *SlackDefinition) Definition() *object.Definition { return d.def }

// MarshalJSON returns the JSON value for the given slack definition
func (d *SlackDefinition) MarshalJSON() ([]byte, error) { return json.Marshal(d.def) }

// Spec returns the slack specification from within the slack definition
func (d *SlackDefinition) Spec() *Slack { return d.def.Spec.(*Slack) }
