package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{
			input:  "heLlo WorlD",
			output: []string {
				"hello",
				"world",
			},
		},
	}

	for _ , cs := range cases {
		actual := cleanInput(cs.input)
		expected := cs.output

		if len(actual) != len(expected) {
			t.Errorf("the lengths are not equal: %v vs %v", len(actual), len(expected))
			continue
		}

		for i := range actual {
			actualWord := actual[i]
			expectedWord := expected[i]

			if actualWord != expectedWord {
				t.Errorf("%v does not match %v", actualWord, expectedWord)
			}
		}
	}
}