package main

import (
	"fmt"
	"os"

	"golang.org/x/text/language"
)

func main() {
	for _, arg := range os.Args[1:] {
		tag, err := language.Parse(arg); if err != nil {
			fmt.Printf("error Could not parse %s due to %v\n", arg, err)
		} else if tag == language.Und {
			fmt.Printf("error Undefined %s\n", arg)
		} else {
			fmt.Printf("info %s tag %s\n", arg, tag)
		}
	}
}