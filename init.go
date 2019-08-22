package configuration

import (
	"github.com/chargehive/configuration/object"
	v1 "github.com/chargehive/configuration/v1"
)

var buildInKinds = make(map[object.Kind]map[string]bool, 0)

func Initialise() {
	handlers := v1.GetHandlers()

	for _, h := range handlers {
		if _, ok := buildInKinds[h.Kind]; !ok {
			buildInKinds[h.Kind] = make(map[string]bool)
		}
		buildInKinds[h.Kind][h.Version] = true
		object.AddKindHandler(h)
	}
}

func GetBuiltInKinds() map[object.Kind]map[string]bool {
	return buildInKinds
}

func IsBuiltInKind(kind object.Kind, version string) bool {
	if has, ok := buildInKinds[kind][version]; ok {
		return has
	}
	return false
}
