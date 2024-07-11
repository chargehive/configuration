package connector

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestConnectorValidate(t *testing.T) {
	validate := validator.New()

	t.Run("empty library", func(t *testing.T) {
		con := &Connector{
			ProcessingState: ProcessingStateLive,
			Library:         "",
			Configuration:   []byte("xxxx"),
			ConfigID:        "",
			ConfigAuth:      "",
			EnablePCIB:      false,
			SCAConnectorID:  "",
		}
		if err := validate.Struct(con); err != nil {
			t.Error(err)
		}
	})

	t.Run("invalid library", func(t *testing.T) {
		con := &Connector{
			ProcessingState: ProcessingStateLive,
			Library:         "chargehive",
			Configuration:   []byte("xxxx"),
			ConfigID:        "",
			ConfigAuth:      "",
			EnablePCIB:      false,
			SCAConnectorID:  "",
		}
		if err := validate.Struct(con); err == nil {
			t.Error("expected an error")
		}
	})

	t.Run("valid library", func(t *testing.T) {
		con := &Connector{
			ProcessingState: ProcessingStateLive,
			Library:         "checkout",
			Configuration:   []byte("xxxx"),
			ConfigID:        "",
			ConfigAuth:      "",
			EnablePCIB:      false,
			SCAConnectorID:  "",
		}
		if err := validate.Struct(con); err != nil {
			t.Error(err)
		}
	})

}
