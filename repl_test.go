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

	for _, c := range cases {
        actual := cleanInput(c.input)

        // 1. Check length
        if len(actual) != len(c.expected) {
            t.Errorf("cleanInput(%q) returned %d words, expected %d words",
                c.input, len(actual), len(c.expected))
            continue // skip further checks for this case
        }

        // 2. Check each word
        for i := range actual {
            if actual[i] != c.expected[i] {
                t.Errorf("cleanInput(%q) word %d = %q, expected %q",
                    c.input, i, actual[i], c.expected[i])
            }
        }
    }
}