package main

import (
	"testing"

	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world!", " ", "!12345"}

	for _, testcase := range testcases {
		f.Add(testcase)
	}

	f.Fuzz(func(t *testing.T, orig string) {
		rev, err := Reverse(orig); if err != nil {
			return
		}

		doubleRev, err := Reverse(rev); if err != nil {
			return
		}

		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(orig), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))
		
		if orig != doubleRev {
			t.Errorf("Before %q, after %q", orig, doubleRev)
		}

		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}