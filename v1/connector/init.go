package connector

import "github.com/chargehive/configuration/object"

func AddHandlers() {
	connector()
	connectorPool()
}

func connectorPool() {
	o := Pool{}
	object.AddKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &Pool{} })
	object.AddKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &Pool{} })
}

func connector() {
	o := Connector{}
	object.AddKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &Connector{} })
	object.AddKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &Connector{} })
}
