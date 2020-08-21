package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Scan Standard Input
var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {

	N, _ := strconv.Atoi(nextLine())

	d := make([]int, N)
	for ix := 0; ix < N; ix++ {
		di, _ := strconv.Atoi(nextLine())

		d[ix] = di
	}

	fmt.Println(len(sortUniq(d)))
}

func sortUniq(a []int) []int {
	r := make([]int, 0)

	sort.SliceStable(a, func(i, j int) bool { return a[i] < a[j] })

	tmp := 0
	for _, b := range a {

		if tmp == b {
			//skip
		} else {
			r = append(r, b)
		}
		tmp = b
	}
	return r
}
