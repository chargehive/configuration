package v1

import (
	"github.com/chargehive/configuration/object"
	"github.com/chargehive/configuration/v1/initiator"
)

func AddHandlers() {
	o := initiator.Initiator{}
	object.AddKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &initiator.Initiator{} })
	object.AddKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &initiator.Initiator{} })
}
