package main

import "strings"

// ConcurrencyFirstBranch returns the first branch name, or empty for blank input.
func ConcurrencyFirstBranch(input string) string {
	return strings.Fields(input)[0]
}
