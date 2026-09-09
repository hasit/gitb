package main

// ConcurrencyLastByte returns the final byte, or zero for an empty string.
func ConcurrencyLastByte(text string) byte {
	return text[len(text)-1]
}
