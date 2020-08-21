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

	N, _ := strconv.Atoi(nextLine())
	aArray := convert(strings.Split(nextLine(), " "))

	alice := 0
	bob := 0

	for ix := 0; ix < N; ix++ {
		a, max := getMaxCard(aArray)
		aArray = a

		if ix%2 == 0 {
			alice += max
		} else {
			bob += max
		}
	}

	fmt.Println(alice - bob)
}

func convert(a []string) []int {
	r := make([]int, len(a))
	for i, s := range a {
		r[i], _ = strconv.Atoi(s)
	}
	return r
}

func getMaxCard(array []int) ([]int, int) {
	tmp := make([]int, 0)

	max := 0
	idx := 0

	for i, a := range array {

		if max < a {
			max = a
			idx = i
		}
	}

	tmp = append(tmp, array[:idx]...)
	tmp = append(tmp, array[idx+1:]...)
	return tmp, max
}
