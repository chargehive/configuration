package v1

import (
	"github.com/chargehive/configuration/object"
	"github.com/chargehive/configuration/v1/connector"
	"github.com/chargehive/configuration/v1/policy"
	"github.com/chargehive/configuration/v1/scheduler"
)

func GetHandlers() []object.KindHandler {
	funcs := make([]object.KindHandler, 0)
	funcs = append(funcs, scheduler.GetHandlers()...)
	funcs = append(funcs, connector.GetHandlers()...)
	funcs = append(funcs, policy.GetHandlers()...)
	return funcs
}
