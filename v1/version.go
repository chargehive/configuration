package v1

import (
	"github.com/chargehive/configuration/object"
	"github.com/chargehive/configuration/v1/scheduler"
)

func AddHandlers() {
	o := scheduler.Initiator{}
	object.AddKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &scheduler.Initiator{} })
	object.AddKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &scheduler.Initiator{} })
}
