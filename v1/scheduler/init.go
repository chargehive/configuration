package scheduler

import (
	"github.com/chargehive/configuration/object"
)

// GetHandlers returns scheduler handlers
func GetHandlers() []object.KindHandler {
	funcs := make([]object.KindHandler, 0)
	funcs = append(funcs, initiator()...)
	funcs = append(funcs, sequentialScheduler()...)
	funcs = append(funcs, onDemandScheduler()...)
	return funcs
}

func initiator() []object.KindHandler {
	o := Initiator{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &Initiator{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &Initiator{} }),
	}
}

func sequentialScheduler() []object.KindHandler {
	o := Sequential{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &Sequential{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &Sequential{} }),
	}
}

func onDemandScheduler() []object.KindHandler {
	o := OnDemand{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &OnDemand{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &OnDemand{} }),
	}
}
