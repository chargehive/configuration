package connector

import "github.com/chargehive/configuration/object"

// GetHandlers returns connector handlers
func GetHandlers() []object.KindHandler {
	funcs := make([]object.KindHandler, 0)
	funcs = append(funcs, connector()...)
	funcs = append(funcs, connectorPool()...)
	return funcs
}

func connectorPool() []object.KindHandler {
	o := Pool{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &Pool{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &Pool{} }),
	}
}

func connector() []object.KindHandler {
	o := Connector{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &Connector{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &Connector{} }),
	}
}
