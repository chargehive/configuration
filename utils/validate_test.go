package utils

import (
	"encoding/json"
	"fmt"
	"github.com/chargehive/configuration"
	"github.com/go-playground/assert/v2"
	"testing"
)

// test for additional unknown fields
func TestAdditionalUnknownFields(t *testing.T) {
	rawJson := []byte(`{"Kind":"Connector","metadata":{"projectId":"CHANGE-ME","bob":"cat","Name":"CHANGE-ME","uuid":"","displayName":"","description":"","annotations":null,"labels":null,"disabled":true},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"paypal-websitepaymentspro","configuration":"eyJhcGlVc2VybmFtZSI6IkNIQU5HRS1NRSIsImFwaVBhc3N3b3JkIjoiQ0hBTkdFLU1FIiwiYXBpU2lnbmF0dXJlIjoiQ0hBTkdFLU1FIiwic3VwcG9ydGVkQ3VycmVuY2llcyI6WyJVU0QiXSwiY2FyZGluYWxQcm9jZXNzb3JJRCI6bnVsbCwiY2FyZGluYWxNZXJjaGFudElEIjpudWxsLCJjYXJkaW5hbFRyYW5zYWN0aW9uUHciOm51bGwsImNhcmRpbmFsVHJhbnNhY3Rpb25VUkwiOm51bGwsImNhcmRpbmFsQVBJSWRlbnRpZmllciI6bnVsbCwiY2FyZGluYWxBUElLZXkiOm51bGwsImNhcmRpbmFsT3JnVW5pdElEIjpudWxsLCJlbnZpcm9ubWVudCI6InNhbmRib3gifQ=="}}`)
	configuration.Initialise()
	errs := Validate(rawJson, "v1")
	_ = PrettyPrint(errs)
	assert.Equal(t, errs["json"], "json: unknown field \"bob\"")
}

// test for missing field
func TestMissingFields(t *testing.T) {
	rawJson := []byte(`{"Kind":"Connector","metadata":{"Name":"CHANGE-ME","uuid":"","displayName":"","description":"","annotations":null,"labels":null},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equal","conversion":"","values":["GBP"]}]},"spec":{"library":"paypal-websitepaymentspro","configuration":"eyJhcGlQYXNzd29yZCI6IkNIQU5HRS1NRSIsImFwaVNpZ25hdHVyZSI6IkNIQU5HRS1NRSIsInN1cHBvcnRlZEN1cnJlbmNpZXMiOlsiVVNEIl0sImNhcmRpbmFsUHJvY2Vzc29ySUQiOm51bGwsImNhcmRpbmFsTWVyY2hhbnRJRCI6bnVsbCwiY2FyZGluYWxUcmFuc2FjdGlvblB3IjpudWxsLCJjYXJkaW5hbFRyYW5zYWN0aW9uVVJMIjpudWxsLCJjYXJkaW5hbEFQSUlkZW50aWZpZXIiOm51bGwsImNhcmRpbmFsQVBJS2V5IjpudWxsLCJjYXJkaW5hbE9yZ1VuaXRJRCI6bnVsbCwiZW52aXJvbm1lbnQiOiJzYW5kYm94In0="}}`)
	configuration.Initialise()
	if errs := Validate(rawJson, "v1"); len(errs) > 0 {
		_ = PrettyPrint(errs)
		assert.Equal(t, 2, len(errs))
		assert.Equal(t, errs["PayPalWebsitePaymentsProCredentials.APIUsername"], "APIUsername is a required field")
		assert.Equal(t, errs["Definition.MetaData.ProjectID"], "ProjectID is a required field")
	}
}

// test for invalid values in correct fields
func TestValidation(t *testing.T) {
	rawJson := []byte(`{"Kind":"Connector","metadata":{"projectId":"CHANGE-ME","Name":"CHANGE-ME","uuid":"","displayName":"","description":"","annotations":null,"labels":null,"disabled":true},"specVersion":"v1","selector":{"priority":50,"expressions":[{"key":"charge.amount.currency","operator":"Equals","conversion":"","values":["GBP"]}]},"spec":{"library":"paypal-websitepaymentspro","configuration":"eyJhcGlVc2VybmFtZSI6IkNIQU5HRS1NRSIsImFwaVBhc3N3b3JkIjoiQ0hBTkdFLU1FIiwiYXBpU2lnbmF0dXJlIjoiQ0hBTkdFLU1FIiwic3VwcG9ydGVkQ3VycmVuY2llcyI6WyJCT0IiXSwiY2FyZGluYWxQcm9jZXNzb3JJRCI6bnVsbCwiY2FyZGluYWxNZXJjaGFudElEIjpudWxsLCJjYXJkaW5hbFRyYW5zYWN0aW9uUHciOm51bGwsImNhcmRpbmFsVHJhbnNhY3Rpb25VUkwiOm51bGwsImNhcmRpbmFsQVBJSWRlbnRpZmllciI6bnVsbCwiY2FyZGluYWxBUElLZXkiOm51bGwsImNhcmRpbmFsT3JnVW5pdElEIjpudWxsLCJlbnZpcm9ubWVudCI6InNhbmRib3gifQ=="}}`)
	configuration.Initialise()
	if errs := Validate(rawJson, "v1"); len(errs) > 0 {
		_ = PrettyPrint(errs)
		assert.Equal(t, 2, len(errs))
		assert.Equal(t, errs["PayPalWebsitePaymentsProCredentials.SupportedCurrencies[0]"], "SupportedCurrencies[0] must be one of [AUD BRL CAD CZK DKK EUR HKD HUF INR ILS JPY MYR MXN TWD NZD NOK PHP PLN GBP RUB SGD SEK CHF THB USD]")
		assert.Equal(t, errs["Definition.Selector.Expressions[0].Operator"], "Operator must be one of [Equal NotEqual In NotIn Exists DoesNotExists Gt Lt]")
	}
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
