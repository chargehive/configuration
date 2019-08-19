package scheduler

import (
	"github.com/chargehive/configuration/object"
)

func AddHandlers() {
	initiator()
	attemptScheduler()
	onDemandScheduler()
}

func initiator() {
	o := Initiator{}
	object.AddKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &Initiator{} })
	object.AddKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &Initiator{} })
}

func attemptScheduler() {
	o := AttemptScheduler{}
	object.AddKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &AttemptScheduler{} })
	object.AddKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &AttemptScheduler{} })
}

func onDemandScheduler() {
	o := OnDemandScheduler{}
	object.AddKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &OnDemandScheduler{} })
	object.AddKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &OnDemandScheduler{} })
}
