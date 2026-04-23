package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct{
		input string
		expected []string
	}{
		{
			input: " Hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input: "   ",
			expected: []string{},
		},
		{
			input: "Single",
			expected: []string{"single"},
		},
	}
}