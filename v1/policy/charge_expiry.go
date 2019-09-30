package policy

import (
	"encoding/json"
	"github.com/chargehive/configuration/object"
	"time"
)

const KindPolicyChargeExpiry object.Kind = "PolicyChargeExpiry"

// ChargeExpiryPolicy - The maximum amount of time or attempts to process a charge for
type ChargeExpiryPolicy struct {
	Timeout  time.Duration
	Attempts int64
}

func (ChargeExpiryPolicy) GetKind() object.Kind { return KindPolicyChargeExpiry }
func (ChargeExpiryPolicy) GetVersion() string   { return "v1" }

type ChargeExpiryDefinition struct{ def *object.Definition }

func (d *ChargeExpiryDefinition) Definition() *object.Definition { return d.def }
func (d *ChargeExpiryDefinition) MarshalJSON() ([]byte, error)   { return json.Marshal(d.def) }
func (d *ChargeExpiryDefinition) Spec() *ChargeExpiryPolicy      { return d.def.Spec.(*ChargeExpiryPolicy) }
