package main

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

	seq := make([]int, 0)

	for i := 0; i < count; i++ {
		n, _ := strconv.Atoi(nextLine())
		seq = append(seq, n)
	}

	max := math.MinInt32

	for idx, v := range seq {
		t := findMax(seq[idx+1:])
		r := t - v

		if max < r {
			max = r
		}
	}

	fmt.Println(max)
}

func findMax(s []int) int {
	max := math.MinInt32
	for _, target := range s {
		if max < target {
			max = target
		}
	}
	return max
}
