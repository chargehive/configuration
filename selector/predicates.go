package selector

type PredicateOperator string

const (
	PredicateOperatorEqual         PredicateOperator = "Equal"
	PredicateOperatorNotEqual      PredicateOperator = "NotEqual"
	PredicateOperatorIn            PredicateOperator = "In"
	PredicateOperatorNotIn         PredicateOperator = "NotIn"
	PredicateOperatorExists        PredicateOperator = "Exists"
	PredicateOperatorDoesNotExists PredicateOperator = "DoesNotExists"
	PredicateOperatorGreaterThan   PredicateOperator = "Gt"
	PredicateOperatorLessThan      PredicateOperator = "Lt"
)

var PredicateOperatorRegister = map[PredicateOperator]bool{
	PredicateOperatorEqual:         true,
	PredicateOperatorNotEqual:      true,
	PredicateOperatorIn:            true,
	PredicateOperatorNotIn:         true,
	PredicateOperatorExists:        true,
	PredicateOperatorDoesNotExists: true,
	PredicateOperatorGreaterThan:   true,
	PredicateOperatorLessThan:      true,
}

type OperatorConversion string

const (
	OperatorConversionDefault         OperatorConversion = ""
	OperatorConversionTimeDayOfWeek   OperatorConversion = "TimeDow"         // Day Of Week
	OperatorConversionTimeMonth       OperatorConversion = "TimeMonth"       // Month
	OperatorConversionDurationSeconds OperatorConversion = "DurationSeconds" // Duration in Seconds
	OperatorConversionDurationHours   OperatorConversion = "DurationHours"   // Duration in Hours
	OperatorConversionDurationDays    OperatorConversion = "DurationDays"    // Duration in Days
)

var OperatorConversionRegister = map[OperatorConversion]bool{
	OperatorConversionDefault:         true,
	OperatorConversionTimeDayOfWeek:   true,
	OperatorConversionTimeMonth:       true,
	OperatorConversionDurationSeconds: true,
	OperatorConversionDurationHours:   true,
	OperatorConversionDurationDays:    true,
}

type Predicate struct {
	Key        Key                `json:"key" yaml:"key" validate:"predicate-key"`
	Operator   PredicateOperator  `json:"operator" yaml:"operator" validate:"predicate-operator"`
	Conversion OperatorConversion `json:"conversion,omitempty" yaml:"conversion,omitempty" validate:"predicate-operator-conversion"`
	Values     []string           `json:"values,omitempty" yaml:"values,omitempty"`
}
