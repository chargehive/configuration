package scheduler

import (
	"github.com/chargehive/configuration/object"
)

func AddHandlers() {
	initiator()
	sequentialScheduler()
	onDemandScheduler()
}

func initiator() {
	o := Initiator{}
	object.AddKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &Initiator{} })
	object.AddKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &Initiator{} })
}

func sequentialScheduler() {
	o := Sequential{}
	object.AddKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &Sequential{} })
	object.AddKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &Sequential{} })
}

func onDemandScheduler() {
	o := OnDemand{}
	object.AddKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &OnDemand{} })
	object.AddKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &OnDemand{} })
}
