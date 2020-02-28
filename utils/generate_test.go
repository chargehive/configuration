package utils

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGenerate(t *testing.T) {
	fmt.Println("Testing templates:")
	for k := range Templates {
		fmt.Printf("- %v\n", k)
		result, err := Generate(k, false)
		fmt.Println(string(result))
		if err != nil {
			fmt.Println(err)
		}
		assert.Equal(t, err, nil)
	}
}
