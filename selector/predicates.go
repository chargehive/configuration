package selector

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

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
	Key        string             `json:"key" yaml:"key"`
	Operator   PredicateOperator  `json:"operator" yaml:"operator"`
	Conversion OperatorConversion `json:"conversion,omitempty" yaml:"conversion,omitempty"`
	Values     []string           `json:"values,omitempty" yaml:"values,omitempty"`
}

func (p *Predicate) MatchSystem() (matched, checked bool) {
	// System Selectors
	switch Key(strings.ToLower(p.Key)) {
	case KeyNow:
		return p.matchTime(time.Now())
	case KeyRandomPercent:
		return p.matchInt(int64(rand.Intn(100)))
	}
	return false, false
}

func (p *Predicate) matchTime(propertyValue time.Time) (matched, checked bool) {

	switch p.Conversion {
	case OperatorConversionDurationDays:
	case OperatorConversionDurationHours:
	case OperatorConversionDurationSeconds:

		var duration int64
		compare, _ := strconv.ParseInt(p.Values[0], 10, 64)
		switch p.Conversion {
		case OperatorConversionDurationDays:
			duration = int64(math.Floor(time.Since(propertyValue).Hours() / 24))
		case OperatorConversionDurationHours:
			duration = int64(math.Floor(time.Since(propertyValue).Hours()))
		case OperatorConversionDurationSeconds:
			duration = int64(math.Floor(time.Since(propertyValue).Seconds()))
		}

		switch p.Operator {
		case PredicateOperatorEqual:
			return duration == compare, true
		case PredicateOperatorNotEqual:
			return duration != compare, true
		case PredicateOperatorGreaterThan:
			return duration > compare, true
		case PredicateOperatorLessThan:
			return duration < compare, true
		}

	case OperatorConversionTimeDayOfWeek:
		t1, err := time.Parse("Mon", p.Values[0])
		if err == nil {
			switch p.Operator {
			case PredicateOperatorEqual:
				return propertyValue.Weekday() == t1.Weekday(), true
			case PredicateOperatorNotEqual:
				return propertyValue.Weekday() != t1.Weekday(), true
			case PredicateOperatorGreaterThan:
				return propertyValue.Weekday() > t1.Weekday(), true
			case PredicateOperatorLessThan:
				return propertyValue.Weekday() < t1.Weekday(), true
			}
		}
	case OperatorConversionTimeMonth:
		t1, err := time.Parse("Jan", p.Values[0])
		if err == nil {
			switch p.Operator {
			case PredicateOperatorEqual:
				return propertyValue.Month() == t1.Month(), true
			case PredicateOperatorNotEqual:
				return propertyValue.Month() != t1.Month(), true
			case PredicateOperatorGreaterThan:
				return propertyValue.Month() > t1.Month(), true
			case PredicateOperatorLessThan:
				return propertyValue.Month() < t1.Month(), true
			}
		}

	case OperatorConversionDefault:
	default:
		t1, err := time.Parse(time.RFC3339, p.Values[0])
		if err == nil {
			switch p.Operator {
			case PredicateOperatorEqual:
				return propertyValue == t1, true
			case PredicateOperatorNotEqual:
				return propertyValue != t1, true
			case PredicateOperatorGreaterThan:
				return propertyValue.Before(t1), true
			case PredicateOperatorLessThan:
				return propertyValue.After(t1), true
			}
		}
	}

	return false, false
}

func (p *Predicate) matchString(propertyValue string) (matched, checked bool) {
	switch p.Operator {
	case PredicateOperatorEqual:
		return propertyValue == p.Values[0], true
	case PredicateOperatorNotEqual:
		return propertyValue != p.Values[0], true
	case PredicateOperatorGreaterThan:
		return strings.Compare(propertyValue, p.Values[0]) > 0, true
	case PredicateOperatorLessThan:
		return strings.Compare(propertyValue, p.Values[0]) < 0, true
	case PredicateOperatorNotIn:
		located := false
		for _, val := range p.Values {
			if val == propertyValue {
				located = true
			}
		}
		return !located, true
	case PredicateOperatorIn:
		for _, val := range p.Values {
			if val == propertyValue {
				return true, true
			}
		}
	}
	return false, false
}
func (p *Predicate) matchInt(propertyValue int64) (matched, checked bool) {
	var values = make([]int64, 0)
	for _, val := range p.Values {
		intval, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			values = append(values, intval)
		}
	}
	switch p.Operator {
	case PredicateOperatorEqual:
		return propertyValue == values[0], true
	case PredicateOperatorNotEqual:
		return propertyValue != values[0], true
	case PredicateOperatorGreaterThan:
		return propertyValue > values[0], true
	case PredicateOperatorLessThan:
		return propertyValue < values[0], true
	case PredicateOperatorNotIn:
		located := false
		for _, val := range values {
			if val == propertyValue {
				located = true
			}
		}
		return !located, true
	case PredicateOperatorIn:
		for _, val := range values {
			if val == propertyValue {
				return true, true
			}
		}
	}
	return false, false
}
