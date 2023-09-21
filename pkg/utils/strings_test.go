package utils

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicatedSpace(t *testing.T) {
	tests := []string{
		"hello world ",
		" hello  world",
		"hello  from  another     world",
	}

	for _, s := range tests {
		n := RemoveDuplicatedSpace(s)
		fmt.Printf("|%s| \t = \t|%s|\n", s, n)
	}
}
