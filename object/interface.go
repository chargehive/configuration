package object

type Specification interface {
	GetKind() Kind
	GetVersion() string
}
