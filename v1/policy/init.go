package policy

import "github.com/chargehive/configuration/object"

func AddHandlers() {
	scaPolicy()
	fraudPolicy()
}

func scaPolicy() {
	o := ScaPolicy{}
	object.AddKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &ScaPolicy{} })
	object.AddKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &ScaPolicy{} })
}

func fraudPolicy() {
	o := FraudPolicy{}
	object.AddKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &FraudPolicy{} })
	object.AddKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &FraudPolicy{} })
}
