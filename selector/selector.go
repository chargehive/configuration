package selector

type Selector struct {
	Priority    int32       `json:"priority" yaml:"priority"`
	Expressions []Predicate `json:"expressions" yaml:"expressions" validate:"dive"`
}
