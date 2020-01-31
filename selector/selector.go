package selector

type Selector struct {
	Priority    int32       `json:"priority,omitempty" yaml:"priority,omitempty"`
	Expressions []Predicate `json:"expressions,omitempty" yaml:"expressions,omitempty" validate:"dive"`
}
