package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Scan Standard Input
var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	n, _ := strconv.Atoi(nextLine())
	s := strings.Split(nextLine(), " ")
	r := convert(s, n)

	result := 0
	for {
		rr, b := allEvent(r)
		if !b {
			break
		}
		r = rr
		result++
	}
	fmt.Println(result)
}

func allEvent(r []int) ([]int, bool) {
	for i, a := range r {
		if a%2 == 1 {
			return r, false
		} else {
			r[i] = a / 2
		}
	}
	return r, true
}

func convert(s []string, n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(s[i])
		r[i] = a
	}
	return r
}
