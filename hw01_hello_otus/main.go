package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	phrase := "Hello, OTUS!"
	reversedPhrase := reverse.String(phrase)

	fmt.Println(reversedPhrase)
}
