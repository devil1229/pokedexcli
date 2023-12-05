package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	//writing unit testcases for cleaninput function
	//creating the cases 
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
        
		//check if lengths are equal or not
		if len(actual) != len(expected) {
			t.Errorf("the lengths are not equal: %v vs %v", len(actual), len(expected))
			continue
		}
 
		//then check all the words in both slices 
		//we can run following for loop because we have checked the legth of the slices and they are same
		for i := range actual {
			actualWord := actual[i]
			expectedWord := expected[i]

			if actualWord != expectedWord {
				t.Errorf("%v does not match %v", actualWord, expectedWord)
			}
		}
	}
}