package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Scane Standard Input
var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s := nextLine()

	r := make(map[string]int)

	for _, ss := range s {
		r[string(ss)] = r[string(ss)] + 1
	}

	for k, v := range r {
		fmt.Println(k + " : " + strconv.Itoa(v))
	}
}
