package object

var kindHandlers = map[Kind]map[string]func() Specification{}

const KindHandlerDefaultVersion = "_DEFAULT_"

func AddKindHandler(handler KindHandler) {
	if _, ok := kindHandlers[handler.Kind]; !ok {
		kindHandlers[handler.Kind] = make(map[string]func() Specification, 0)
	}
	kindHandlers[handler.Kind][handler.Version] = handler.handler
}

func getKindHandlerFunc(kind Kind, version string) (func() Specification, bool) {
	if handler, ok := kindHandlers[kind][version]; ok {
		return handler, true
	}

	// Look for the default Version
	if handler, ok := kindHandlers[kind][KindHandlerDefaultVersion]; ok {
		return handler, true
	}
	return func() Specification { return nil }, false
}

type KindHandler struct {
	Kind    Kind
	Version string
	handler func() Specification
}

func NewKindHandler(kind Kind, version string, handler func() Specification) KindHandler {
	return KindHandler{Kind: kind, Version: version, handler: handler}
}
