package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		if _, exists := wordCount[word]; exists {
			wordCount[word]++
		} else {
			wordCount[word] = 1
		}
	}

	return wordCount
}

func main() {
	wc.Test(WordCount)
}
