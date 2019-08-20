package v1

import (
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/configuration/v1/policy"
	"github.com/chargehive/configuration/v1/scheduler"
)

func AddHandlers() {
	connector.AddHandlers()
	policy.AddHandlers()
	scheduler.AddHandlers()
}
