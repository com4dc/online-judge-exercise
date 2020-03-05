package main

// Dynamic Programming //
// minv = R[0]
// for j が 1 から n-1 まで
//   maxv = (maxv と R[j]-minv のうち大きい方)
//   minv = (minv と R[j] のうち小さい方)

import (
	"bufio"
	"fmt"
	"math"
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
	count, _ := strconv.Atoi(nextLine())

	n0, _ := strconv.Atoi(nextLine())
	minv := n0
	maxv := math.MinInt32

	for i := 1; i < count; i++ {
		n, _ := strconv.Atoi(nextLine())

		if maxv < n-minv {
			maxv = n - minv
		}

		if minv > n {
			minv = n
		}
	}

	fmt.Println(maxv)
}
