package policy

import "github.com/chargehive/configuration/object"

func GetHandlers() []object.KindHandler {
	funcs := make([]object.KindHandler, 0)
	funcs = append(funcs, scaPolicy()...)
	funcs = append(funcs, fraudPolicy()...)
	funcs = append(funcs, chargeExpiryPolicy()...)
	funcs = append(funcs, methodUpgradePolicy()...)
	funcs = append(funcs, cascadePolicy()...)
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

func chargeExpiryPolicy() []object.KindHandler {
	o := ChargeExpiryPolicy{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &ChargeExpiryPolicy{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &ChargeExpiryPolicy{} }),
	}
}

func methodUpgradePolicy() []object.KindHandler {
	o := MethodUpgradePolicy{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &MethodUpgradePolicy{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &MethodUpgradePolicy{} }),
	}
}

func cascadePolicy() []object.KindHandler {
	o := CascadePolicy{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &CascadePolicy{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &CascadePolicy{} }),
	}
}
