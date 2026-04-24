package main

import "strings"


func cleanInput(text string) []string{
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return []string{}
	}
	
	words := strings.Fields(trimmed)

	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}