package object

var kindHandlers = map[Kind]map[string]func() Specification{}

const KindHandlerDefaultVersion = "_DEFAULT_"

func AddKindHandler(kind Kind, version string, handler func() Specification) {
	if _, ok := kindHandlers[kind]; !ok {
		kindHandlers[kind] = make(map[string]func() Specification, 0)
	}
	kindHandlers[kind][version] = handler
}

func getKindHandler(kind Kind, version string) (func() Specification, bool) {
	if handler, ok := kindHandlers[kind][version]; ok {
		return handler, true
	}

	//Look for the default version
	if handler, ok := kindHandlers[kind][KindHandlerDefaultVersion]; ok {
		return handler, true
	}
	return func() Specification { return nil }, false
}
