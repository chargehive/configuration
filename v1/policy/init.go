package policy

import "github.com/chargehive/configuration/object"

func GetHandlers() []object.KindHandler {
	funcs := make([]object.KindHandler, 0)
	funcs = append(funcs, scaPolicy()...)
	funcs = append(funcs, fraudPolicy()...)
	return funcs
}

func scaPolicy() []object.KindHandler {
	o := ScaPolicy{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &ScaPolicy{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &ScaPolicy{} }),
	}
}

func fraudPolicy() []object.KindHandler {
	o := FraudPolicy{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &FraudPolicy{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &FraudPolicy{} }),
	}
}
