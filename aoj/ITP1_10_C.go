package main

import (
	"bufio"
	"fmt"
	"math"
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
	for {
		n, _ := strconv.Atoi(nextLine())

		if n == 0 {
			return
		}

		ss := strings.Split(nextLine(), " ")

		m := average(ss)

		s_total := 0.0
		for i := 0; i < n; i++ {
			s, _ := strconv.ParseFloat(ss[i], 64)
			s_total += (s - m) * (s - m)
		}
		x := math.Sqrt(s_total / (float64)(n))
		fmt.Printf("%.8f\n", x)
	}
}

func average(xs []string) float64 {
	total := 0.0
	for _, v := range xs {
		a, _ := strconv.ParseFloat(v, 64)
		total += a
	}
	return total / (float64)(len(xs))
}
