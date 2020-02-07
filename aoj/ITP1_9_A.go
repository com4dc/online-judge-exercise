package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Scane Standard Input
var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {

	w := nextLine()

	wordCount := 0
	for {
		t := strings.Split(nextLine(), " ")
		if t[0] == "END_OF_TEXT" {
			fmt.Println(wordCount)
			return
		}

		for _, tt := range t {
			if w == strings.ToLower(tt) {
				wordCount++
			}
		}
	}
}
