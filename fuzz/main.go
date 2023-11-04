package main

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }

	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r), nil
}

func main() {
	startTime := time.Now()

	input := "The quick brown fox jumped over the lazy dog"
	rev, revErr := Reverse(input)
	doubleRev, doubleRevErr := Reverse(rev)
	fmt.Printf("debug original: %q\n", input)
	fmt.Printf("debug reversed: %q, err: %v\n", rev, revErr)
	fmt.Printf("debug reversed again: %q, err: %v\n", doubleRev, doubleRevErr)

	endTime := time.Now()
	executionTime := endTime.Sub(startTime).Seconds() * 1e3
	fmt.Printf("done in %.2f ms\n", executionTime)
}