package main

import "strings"

// ConcurrencyFirstBranch returns the first branch name, or empty for blank input.
func ConcurrencyFirstBranch(input string) string {
	fields := strings.Fields(input)
	if len(fields) == 0 { return "" }
	return fields[0]
}
