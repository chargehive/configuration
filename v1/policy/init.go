package policy

import "github.com/chargehive/configuration/object"

// GetHandlers returns policy handlers
func GetHandlers() []object.KindHandler {
	funcs := []object.KindHandler{}
	funcs = append(funcs, scaPolicy()...)
	funcs = append(funcs, fraudPolicy()...)
	funcs = append(funcs, chargeExpiryPolicy()...)
	funcs = append(funcs, methodUpgradePolicy()...)
	funcs = append(funcs, cascadePolicy()...)
	funcs = append(funcs, methodLockPolicy()...)
	funcs = append(funcs, methodVerifyPolicy()...)
	funcs = append(funcs, methodRefreshPolicy()...)
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

func methodLockPolicy() []object.KindHandler {
	o := MethodLockPolicy{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &MethodLockPolicy{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &MethodLockPolicy{} }),
	}
}

func methodVerifyPolicy() []object.KindHandler {
	o := MethodVerifyPolicy{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &MethodVerifyPolicy{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &MethodVerifyPolicy{} }),
	}
}

func methodRefreshPolicy() []object.KindHandler {
	o := MethodRefreshPolicy{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &MethodRefreshPolicy{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &MethodRefreshPolicy{} }),
	}
}
