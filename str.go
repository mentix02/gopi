package gopi

// Chr returns a character with ordinal i.
func Chr(i int64) rune {
	return rune(i)
}

// Ord returns the Unicode point for a character.
func Ord(c rune) int64 {
	return int64(c)
}
