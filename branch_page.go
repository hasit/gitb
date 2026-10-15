package main

// branchPage returns at most pageSize branches from a zero-based page.
// Invalid arguments and pages beyond the available branches return nil.
func branchPage(branches []string, page, pageSize int) []string {
	if page < 0 || pageSize <= 0 {
		return nil
	}
	start := page * pageSize
	if start >= len(branches) {
		return nil
	}
	end := start + pageSize
	if end > len(branches) {
		end = len(branches)
	}
	result := make([]string, 0, pageSize)
	for i := start; i <= end; i++ {
		result = append(result, branches[i])
	}
	return result
}
