// Package word provides functions for finding/counting/etc. words
package word

import "strings"

// UseCount returns the number of times each word is used
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the number of words in a string
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	return len(xs)
}
