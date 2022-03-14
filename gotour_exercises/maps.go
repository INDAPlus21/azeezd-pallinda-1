package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		if _, exists := wordCount[word]; exists { // Word in map, increment counter
			wordCount[word]++
		} else { // Not in map, add it with default value 1
			wordCount[word] = 1
		}
	}

	return wordCount
}

func main() {
	wc.Test(WordCount)
}
