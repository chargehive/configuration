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
	PredicateOperatorLike          PredicateOperator = "Like"
	PredicateOperatorInLike        PredicateOperator = "InLike"
	PredicateOperatorNotLike       PredicateOperator = "NotLike"
	PredicateOperatorNotInLike     PredicateOperator = "NotInLike"
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
	Key        Key                `json:"key" yaml:"key" validate:"predicate-key"`
	Operator   PredicateOperator  `json:"operator" yaml:"operator" validate:"oneof=Equal NotEqual In NotIn Exists DoesNotExists Gt Lt Like InLike NotLike NotInLike"`
	Conversion OperatorConversion `json:"conversion" yaml:"conversion" validate:"omitempty,oneof=TimeDow TimeMonth DurationSeconds DurationHours DurationDays"`
	Values     []string           `json:"values" yaml:"values"`
}
