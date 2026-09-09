package main

import "strings"

// branchNames extracts branch names from the output of git branch.
func branchNames(output string) []string {
	lines := strings.Split(output, "\n")
	names := make([]string, 0, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "*" {
			fields = fields[1:]
		}
		names = append(names, fields[0])
	}
	return names
}
