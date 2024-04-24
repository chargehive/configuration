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
	funcs = append(funcs, refundScheduler()...)
	funcs = append(funcs, onTriggerScheduler()...)
	funcs = append(funcs, connectorScheduler()...)
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

func onTriggerScheduler() []object.KindHandler {
	o := OnTrigger{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &OnTrigger{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &OnTrigger{} }),
	}
}

func onDemandScheduler() []object.KindHandler {
	o := OnDemand{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &OnDemand{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &OnDemand{} }),
	}
}

func connectorScheduler() []object.KindHandler {
	o := Connector{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &Connector{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &Connector{} }),
	}
}

func refundScheduler() []object.KindHandler {
	o := Refund{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &Refund{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &Refund{} }),
	}
}
