package policy

import (
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
