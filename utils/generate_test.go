package utils

import (
	"fmt"
	"github.com/chargehive/configuration"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGenerate(t *testing.T) {
	configuration.Initialise()
	fmt.Println("Testing templates:")
	for k := range Templates {
		fmt.Printf("- %v\n", k)
		result, err := Generate(k, "v1", false)
		fmt.Println(string(result))
		if err != nil {
			fmt.Println(err)
		}
		assert.Equal(t, err, nil)
	}
}
