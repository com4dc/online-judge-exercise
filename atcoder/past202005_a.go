package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Scan Standard Input
var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	a := nextLine()
	b := nextLine()

	if a == b {
		fmt.Println("same")
	} else if strings.ToLower(a) == strings.ToLower(b) {
		fmt.Println("case-insensitive")
	} else {
		fmt.Println("different")
	}
}
