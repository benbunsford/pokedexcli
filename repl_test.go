package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected []string
	}{
		"leading/trailing spaces": {input: "     hello     world     ", expected: []string{"hello", "world"}},
		"newlines":                {input: "  \nhello \n  world\n  ", expected: []string{"hello", "world"}},
		"no whitespace":           {input: "helloworld", expected: []string{"helloworld"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(tc.input)
			diff := cmp.Diff(tc.expected, actual)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
