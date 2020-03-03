package utils

import (
	"fmt"
	"github.com/chargehive/configuration"
	"github.com/go-playground/assert/v2"
	"testing"
)

// generate all config templates and validate them as they're generated
func TestGenerate(t *testing.T) {
	configuration.Initialise()
	fmt.Println("Testing templates: ")
	for k := range Templates {
		output, err := Generate(k, "v1", false)
		fmt.Printf("%-30v%s\n", k, output)
		if err != nil {
			fmt.Println(err)
		}
		assert.Equal(t, err, nil)
	}
}
