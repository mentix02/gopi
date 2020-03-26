package gopi

import "testing"

var chars = [26]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z'}

func TestOrd(t *testing.T) {
	for num, char := range chars {
		if Ord(char) != int64(num+97) {
			t.Errorf("%d != %d", Ord(char), num+97)
		}
	}
}

func TestChr(t *testing.T) {
	for num, char := range chars {
		if Chr(int64(num+97)) != char {
			t.Errorf("%c != %c", Chr(int64(num+97)), char)
		}
	}
}
