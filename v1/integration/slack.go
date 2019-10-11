package integration

import (
	"encoding/json"
	"errors"
	"github.com/chargehive/configuration/object"
)

const KindSlackWebhook object.Kind = "Integration.SlackWebhook"

type SlackWebhook struct {
	WebhookUrl string
}

func (SlackWebhook) GetKind() object.Kind { return KindSlackWebhook }
func (SlackWebhook) GetVersion() string   { return "v1" }

func NewSlackWebhookDefinition(d *object.Definition) (*SlackWebhookDefinition, error) {
	if _, ok := d.Spec.(*SlackWebhook); ok {
		return &SlackWebhookDefinition{def: d}, nil
	}
	return nil, errors.New("invalid slack webhook")
}

type SlackWebhookDefinition struct{ def *object.Definition }

func (d *SlackWebhookDefinition) Definition() *object.Definition { return d.def }
func (d *SlackWebhookDefinition) MarshalJSON() ([]byte, error)   { return json.Marshal(d.def) }
func (d *SlackWebhookDefinition) Spec() *SlackWebhook            { return d.def.Spec.(*SlackWebhook) }
