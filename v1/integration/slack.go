package integration

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

const KindSlack object.Kind = "Integration.Slack"

type Slack struct {
	AccessToken string
	Scopes      []string
	TeamName    string
	TeamID      string
	Webhook     *SlackWebhook

	// What to do
	TransactionNotifications bool
}

type SlackWebhook struct {
	Url              string
	Channel          string
	ConfigurationUrl string
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
