package integration

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

const KindSlack object.Kind = "Integration.Slack"

type Slack struct {
	AccessToken string        `json:"accessToken"`
	Scopes      []string      `json:"scopes"`
	TeamName    string        `json:"teamName"`
	TeamID      string        `json:"teamID"`
	Webhook     *SlackWebhook `json:"webhook"`

	// What to do
	TransactionNotifications bool `json:"transactionNotifications"`
}

type SlackWebhook struct {
	Url              string `json:"url"`
	Channel          string `json:"channel"`
	ConfigurationUrl string `json:"configurationUrl"`
}

func (Slack) GetKind() object.Kind { return KindSlack }
func (Slack) GetVersion() string   { return "v1" }

func NewSlackDefinition(d *object.Definition) (*SlackDefinition, error) {
	if _, ok := d.Spec.(*Slack); ok {
		return &SlackDefinition{def: d}, nil
	}
	return nil, errors.New("invalid slack configuration")
}

type SlackDefinition struct{ def *object.Definition }

func (d *SlackDefinition) Definition() *object.Definition { return d.def }
func (d *SlackDefinition) MarshalJSON() ([]byte, error)   { return json.Marshal(d.def) }
func (d *SlackDefinition) Spec() *Slack                   { return d.def.Spec.(*Slack) }
