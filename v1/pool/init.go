package pool

import "github.com/chargehive/configuration/object"

func AddHandlers() {
	connectorPool()
}

func connectorPool() {
	o := ConnectorPool{}
	object.AddKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &ConnectorPool{} })
	object.AddKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &ConnectorPool{} })
}
