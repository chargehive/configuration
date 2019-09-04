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

type OperatorConversion string

const (
	OperatorConversionDefault         OperatorConversion = ""
	OperatorConversionTimeDayOfWeek   OperatorConversion = "TimeDow"         // Day Of Week
	OperatorConversionTimeMonth       OperatorConversion = "TimeMonth"       // Month
	OperatorConversionDurationSeconds OperatorConversion = "DurationSeconds" // Duration in Seconds
	OperatorConversionDurationHours   OperatorConversion = "DurationHours"   // Duration in Hours
	OperatorConversionDurationDays    OperatorConversion = "DurationDays"    // Duration in Days
)

type Predicate struct {
	Key        Key                `json:"key" yaml:"key"`
	Operator   PredicateOperator  `json:"operator" yaml:"operator"`
	Conversion OperatorConversion `json:"conversion,omitempty" yaml:"conversion,omitempty"`
	Values     []string           `json:"values,omitempty" yaml:"values,omitempty"`
}
