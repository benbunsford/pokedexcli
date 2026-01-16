package main

import "strings"

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	cleaned := strings.Fields(lower)
	return cleaned
}
