package utils

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/chargehive/configuration"
	"github.com/chargehive/configuration/selector"
	"github.com/go-playground/assert/v2"
)

// test for additional unknown fields
func TestChargeLabelVariables(t *testing.T) {
	rawJson := []byte(`{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.label.unsubscribed","operator":"Equal","values":["True"]}]},"spec":{"library":"paysafe","configuration":"eyJtZXJjaGFudFVybCI6ImNoYW5nZS1tZSIsImFjcXVpcmVyIjoiY2hhbmdlLW1lIiwiYWNjb3VudElEIjoiY2hhbmdlLW1lIiwiYXBpVXNlcm5hbWUiOiJjaGFuZ2UtbWUiLCJhcGlQYXNzd29yZCI6ImNoYW5nZS1tZSIsImVudmlyb25tZW50IjoiTU9DSyIsImNvdW50cnkiOiIiLCJjdXJyZW5jeSI6IlVTRCIsInVzZVZhdWx0IjpmYWxzZSwic2luZ2xlVXNlVG9rZW5QYXNzd29yZCI6IiIsInNpbmdsZVVzZVRva2VuVXNlcm5hbWUiOiIifQ=="}}`)
	configuration.Initialise()
	errs := Validate(rawJson, "v1")
	// _ = PrettyPrint(errs)
	assert.Equal(t, errs, map[string]string{})
}

// test for additional unknown fields
func TestAdditionalUnknownFields(t *testing.T) {
	rawJson := []byte(`{"Kind":"Connector","metadata":{"projectId":"change-me","bob":"cat","Name":"change-me","uuid":"","displayName":"","description":"","annotations":null,"labels":null,"disabled":true},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"paypal-websitepaymentspro","configuration":"eyJhcGlVc2VybmFtZSI6IkNIQU5HRS1NRSIsImFwaVBhc3N3b3JkIjoiQ0hBTkdFLU1FIiwiYXBpU2lnbmF0dXJlIjoiQ0hBTkdFLU1FIiwic3VwcG9ydGVkQ3VycmVuY2llcyI6WyJVU0QiXSwiY2FyZGluYWxQcm9jZXNzb3JJRCI6bnVsbCwiY2FyZGluYWxNZXJjaGFudElEIjpudWxsLCJjYXJkaW5hbFRyYW5zYWN0aW9uUHciOm51bGwsImNhcmRpbmFsVHJhbnNhY3Rpb25VUkwiOm51bGwsImNhcmRpbmFsQVBJSWRlbnRpZmllciI6bnVsbCwiY2FyZGluYWxBUElLZXkiOm51bGwsImNhcmRpbmFsT3JnVW5pdElEIjpudWxsLCJlbnZpcm9ubWVudCI6InNhbmRib3gifQ=="}}`)
	configuration.Initialise()
	errs := Validate(rawJson, "v1")
	// _ = PrettyPrint(errs)
	assert.Equal(t, errs["json"], "json: unknown field \"bob\"")
}

// test for missing field
func TestMissingFields(t *testing.T) {
	rawJson := []byte(`{"kind":"Connector","metadata":{"displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"paypal-websitepaymentspro","configuration":"eyJhcGlQYXNzd29yZCI6bnVsbCwiYXBpU2lnbmF0dXJlIjoiQ0hBTkdFLU1FIiwic3VwcG9ydGVkQ3VycmVuY2llcyI6WyJVU0QiXSwiY2FyZGluYWxQcm9jZXNzb3JJRCI6IkNIQU5HRS1NRSIsImNhcmRpbmFsTWVyY2hhbnRJRCI6IkNIQU5HRS1NRSIsImNhcmRpbmFsVHJhbnNhY3Rpb25QdyI6IkNIQU5HRS1NRSIsImNhcmRpbmFsVHJhbnNhY3Rpb25VUkwiOiJDSEFOR0UtTUUiLCJjYXJkaW5hbEFQSUlkZW50aWZpZXIiOiJDSEFOR0UtTUUiLCJjYXJkaW5hbEFQSUtleSI6IkNIQU5HRS1NRSIsImNhcmRpbmFsT3JnVW5pdElEIjoiQ0hBTkdFLU1FIiwiZW52aXJvbm1lbnQiOiJzYW5kYm94In0="}}`)
	configuration.Initialise()
	errs := Validate(rawJson, "v1")
	// _ = PrettyPrint(errs)
	assert.Equal(t, 3, len(errs))
	assert.Equal(t, errs["PayPalWebsitePaymentsProCredentials.APIUsername"], "APIUsername is a required field") // missing field
	assert.Equal(t, errs["PayPalWebsitePaymentsProCredentials.APIPassword"], "APIPassword is a required field") // null field
	assert.Equal(t, errs["Definition.MetaData.Name"], "Name is a required field")
}

// test for invalid values in correct fields
func TestValidation(t *testing.T) {
	rawJson := []byte(`{"Kind":"Connector","metadata":{"projectId":"change-me","Name":"change-me","uuid":"","displayName":"","description":"","annotations":null,"labels":null,"disabled":true},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equals","conversion":"","values":["GBP"]}]},"spec":{"library":"paypal-websitepaymentspro","configuration":"eyJhcGlVc2VybmFtZSI6IkNIQU5HRS1NRSIsImFwaVBhc3N3b3JkIjoiQ0hBTkdFLU1FIiwiYXBpU2lnbmF0dXJlIjoiQ0hBTkdFLU1FIiwic3VwcG9ydGVkQ3VycmVuY2llcyI6WyJHQVJZIl0sImNhcmRpbmFsUHJvY2Vzc29ySUQiOiJDSEFOR0UtTUUiLCJjYXJkaW5hbE1lcmNoYW50SUQiOiJDSEFOR0UtTUUiLCJjYXJkaW5hbFRyYW5zYWN0aW9uUHciOiJDSEFOR0UtTUUiLCJjYXJkaW5hbFRyYW5zYWN0aW9uVVJMIjoiQ0hBTkdFLU1FIiwiY2FyZGluYWxBUElJZGVudGlmaWVyIjoiQ0hBTkdFLU1FIiwiY2FyZGluYWxBUElLZXkiOiJDSEFOR0UtTUUiLCJjYXJkaW5hbE9yZ1VuaXRJRCI6IkNIQU5HRS1NRSIsImVudmlyb25tZW50Ijoic2FuZGJveCJ9"}}`)
	configuration.Initialise()
	errs := Validate(rawJson, "v1")
	// _ = PrettyPrint(errs)
	assert.Equal(t, 2, len(errs))
	assert.Equal(t, errs["PayPalWebsitePaymentsProCredentials.SupportedCurrencies[0]"], "SupportedCurrencies[0] must be one of [AUD BRL CAD CZK DKK EUR HKD HUF INR ILS JPY MYR MXN TWD NZD NOK PHP PLN GBP RUB SGD SEK CHF THB USD]")
	assert.Equal(t, errs["Definition.Selector.Expressions[0].Operator"], "Operator must be one of [Equal NotEqual In NotIn Exists DoesNotExists Gt Lt Like InLike NotLike NotInLike]")
}

func TestLowerCaseName(t *testing.T) {
	rawJson := []byte(`{"kind":"Connector","metadata":{"projectId":"change-me","name":"UPPERCASE","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"authorize","configuration":"eyJhcGlMb2dpbklkIjoiQ0hBTkdFLU1FIiwidHJhbnNhY3Rpb25LZXkiOiJDSEFOR0UtTUUiLCJlbnZpcm9ubWVudCI6InNhbmRib3gifQ=="}}`)
	configuration.Initialise()
	errs := Validate(rawJson, "v1")
	// _ = PrettyPrint(errs)
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, errs["Definition.MetaData.Name"], "Name must be a lowercase string")
}

// Tests Emptying string/string* validation (SingleUseTokenPassword *string, singleUseTokenUsername string)
func TestEmptyFields(t *testing.T) {
	configuration.Initialise()

	// "singleUseTokenPassword":"","singleUseTokenUsername":""
	rawJson := []byte(`{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"paysafe","configuration":"eyJhY3F1aXJlciI6ImNoYW5nZS1tZSIsIm1lcmNoYW50VXJsIjoiY2hhbmdlLW1lIiwiYWNjb3VudElEIjoiY2hhbmdlLW1lIiwiYXBpVXNlcm5hbWUiOiJjaGFuZ2UtbWUiLCJhcGlQYXNzd29yZCI6ImNoYW5nZS1tZSIsImVudmlyb25tZW50IjoiTU9DSyIsImNvdW50cnkiOiIiLCJjdXJyZW5jeSI6IlVTRCIsInVzZVZhdWx0IjpmYWxzZSwic2luZ2xlVXNlVG9rZW5QYXNzd29yZCI6IiIsInNpbmdsZVVzZVRva2VuVXNlcm5hbWUiOiIifQ=="}}`)
	errs := Validate(rawJson, "v1")
	assert.Equal(t, len(errs), 0)

	// "singleUseTokenUsername":""
	rawJson = []byte(`{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"paysafe","configuration":"eyJtZXJjaGFudFVybCI6ImNoYW5nZS1tZSIsImFjcXVpcmVyIjoiY2hhbmdlLW1lIiwiYWNjb3VudElEIjoiY2hhbmdlLW1lIiwiYXBpVXNlcm5hbWUiOiJjaGFuZ2UtbWUiLCJhcGlQYXNzd29yZCI6ImNoYW5nZS1tZSIsImVudmlyb25tZW50IjoiTU9DSyIsImNvdW50cnkiOiIiLCJjdXJyZW5jeSI6IlVTRCIsInVzZVZhdWx0IjpmYWxzZSwic2luZ2xlVXNlVG9rZW5Vc2VybmFtZSI6IiJ9"}}`)
	errs = Validate(rawJson, "v1")
	assert.Equal(t, 1, len(errs))
	assert.Equal(t, errs["PaySafeCredentials.SingleUseTokenPassword"], "SingleUseTokenPassword is a required field")

	// "singleUseTokenPassword":""
	rawJson = []byte(`{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"paysafe","configuration":"eyJtZXJjaGFudFVybCI6ImNoYW5nZS1tZSIsImFjcXVpcmVyIjoiY2hhbmdlLW1lIiwiYWNjb3VudElEIjoiY2hhbmdlLW1lIiwiYXBpVXNlcm5hbWUiOiJjaGFuZ2UtbWUiLCJhcGlQYXNzd29yZCI6ImNoYW5nZS1tZSIsImVudmlyb25tZW50IjoiTU9DSyIsImNvdW50cnkiOiIiLCJjdXJyZW5jeSI6IlVTRCIsInVzZVZhdWx0IjpmYWxzZSwic2luZ2xlVXNlVG9rZW5QYXNzd29yZCI6IiJ9"}}`)
	errs = Validate(rawJson, "v1")
	assert.Equal(t, 0, len(errs))
}

// Test for invalid predicates
func TestInvalidPredicates(t *testing.T) {

	var tests = []struct {
		operator selector.PredicateOperator
		values   []string
		valid    bool
	}{
		{selector.PredicateOperatorEqual, []string{}, false},
		{selector.PredicateOperatorEqual, []string{"one"}, true},
		{selector.PredicateOperatorEqual, []string{"one", "two"}, false},

		{selector.PredicateOperatorNotEqual, []string{}, false},
		{selector.PredicateOperatorNotEqual, []string{"one"}, true},
		{selector.PredicateOperatorNotEqual, []string{"one", "two"}, false},

		{selector.PredicateOperatorIn, []string{}, false},
		{selector.PredicateOperatorIn, []string{"one"}, false},
		{selector.PredicateOperatorIn, []string{"one", "two"}, true},

		{selector.PredicateOperatorNotIn, []string{}, false},
		{selector.PredicateOperatorNotIn, []string{"one"}, false},
		{selector.PredicateOperatorNotIn, []string{"one", "two"}, true},

		{selector.PredicateOperatorInLike, []string{}, false},
		{selector.PredicateOperatorInLike, []string{"one"}, false},
		{selector.PredicateOperatorInLike, []string{"one", "two"}, true},

		{selector.PredicateOperatorNotInLike, []string{}, false},
		{selector.PredicateOperatorNotInLike, []string{"one"}, false},
		{selector.PredicateOperatorNotInLike, []string{"one", "two"}, true},
	}

	configuration.Initialise()

	for _, v := range tests {

		values, _ := json.Marshal(v.values)
		rawJson := []byte(`{"kind":"Connector","metadata":{"projectId":"change-me","name":"change-me","displayName":"","description":"","annotations":null,"labels":null,"disabled":false},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.label.unsubscribed","operator":"` + string(v.operator) + `","values":` + string(values) + `}]},"spec":{"library":"paysafe","configuration":"eyJtZXJjaGFudFVybCI6ImNoYW5nZS1tZSIsImFjcXVpcmVyIjoiY2hhbmdlLW1lIiwiYWNjb3VudElEIjoiY2hhbmdlLW1lIiwiYXBpVXNlcm5hbWUiOiJjaGFuZ2UtbWUiLCJhcGlQYXNzd29yZCI6ImNoYW5nZS1tZSIsImVudmlyb25tZW50IjoiTU9DSyIsImNvdW50cnkiOiIiLCJjdXJyZW5jeSI6IlVTRCIsInVzZVZhdWx0IjpmYWxzZSwic2luZ2xlVXNlVG9rZW5QYXNzd29yZCI6IiIsInNpbmdsZVVzZVRva2VuVXNlcm5hbWUiOiIifQ=="}}`)

		errs := Validate(rawJson, "v1")
		if v.valid {
			assert.Equal(t, 0, len(errs))
		} else {
			fmt.Println(errs)
			assert.Equal(t, 1, len(errs))
		}
	}
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
