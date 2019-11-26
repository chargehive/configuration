package v1

import (
	"github.com/chargehive/configuration/object"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/configuration/v1/integration"
	"github.com/chargehive/configuration/v1/policy"
	"github.com/chargehive/configuration/v1/scheduler"
)

// GetHandlers returns v1 handlers
func GetHandlers() []object.KindHandler {
	funcs := make([]object.KindHandler, 0)
	funcs = append(funcs, scheduler.GetHandlers()...)
	funcs = append(funcs, connector.GetHandlers()...)
	funcs = append(funcs, policy.GetHandlers()...)
	funcs = append(funcs, integration.GetHandlers()...)
	return funcs
}
