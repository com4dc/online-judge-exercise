package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Scan Standard Input
var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s := nextLine()
	result := 0

	for _, r := range s {
		s, _ := strconv.Atoi(string(r))

		if s == 1 {
			result++
		}
	}
	fmt.Println(result)
}
