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
	fmt.Print("Testing templates: ")
	count := 0
	for k := range Templates {
		fmt.Print("X")
		_, err := Generate(k, "v1", false)
		if err != nil {
			fmt.Println(err)
		}
		assert.Equal(t, err, nil)
		count++
	}
	fmt.Printf(" - %v\n", count)
}
