package integration

import "github.com/chargehive/configuration/object"

func GetHandlers() []object.KindHandler {
	funcs := make([]object.KindHandler, 0)
	funcs = append(funcs, slack()...)
	return funcs
}

func slack() []object.KindHandler {
	o := SlackWebhook{}
	return []object.KindHandler{
		object.NewKindHandler(o.GetKind(), object.KindHandlerDefaultVersion, func() object.Specification { return &SlackWebhook{} }),
		object.NewKindHandler(o.GetKind(), o.GetVersion(), func() object.Specification { return &SlackWebhook{} }),
	}
}
