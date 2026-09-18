package main

// ConcurrencyClampIndex bounds index to a valid position in a nonempty slice.
func ConcurrencyClampIndex(index, length int) int {
	if index < 0 { return 0 }
	if index > length { return length - 1 }
	return index
}
