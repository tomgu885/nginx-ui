package validator

import (
	"fmt"
	"testing"
)

func TestIsDigital(t *testing.T) {
	tests := []string{
		"1",
		"1.1",
		"1,",
		"12.3",
		"0",
		"199",
		"10_234",
		"234a",
		"John",
	}

	for _, s := range tests {
		ok := IsDigital(s)
		fmt.Printf("%s => %t\n", s, ok)
	}
}
