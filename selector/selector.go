package selector

type Selector struct {
	Priority    int32       `json:"priority" yaml:"priority" validate:"gte=0"`
	Expressions []Predicate `json:"expressions" yaml:"expressions" validate:"dive"`
}
