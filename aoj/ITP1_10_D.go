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

	n, _ := strconv.Atoi(nextLine())
	xs := strings.Split(nextLine(), " ")
	ys := strings.Split(nextLine(), " ")

	M := manhattan(xs, ys, n)
	E := euclid(xs, ys, n)
	M3 := minkowski_3(xs, ys, n)
	C := chebychev(xs, ys, n)

	fmt.Printf("%.8f\n", M)
	fmt.Printf("%.8f\n", E)
	fmt.Printf("%.8f\n", M3)
	fmt.Printf("%.8f\n", C)
}

func manhattan(xs, ys []string, n int) (distance float64) {
	for i := 0; i < n; i++ {
		x, _ := strconv.ParseFloat(xs[i], 64)
		y, _ := strconv.ParseFloat(ys[i], 64)

		distance += math.Abs(x - y)
	}
	return
}

func euclid(xs, ys []string, n int) (distance float64) {
	tmp := 0.0
	for i := 0; i < n; i++ {
		x, _ := strconv.ParseFloat(xs[i], 64)
		y, _ := strconv.ParseFloat(ys[i], 64)

		a := math.Abs(x - y)
		tmp += a * a
	}
	distance = math.Sqrt(tmp)
	return
}

func minkowski_3(xs, ys []string, n int) (distance float64) {
	tmp := 0.0
	for i := 0; i < n; i++ {
		x, _ := strconv.ParseFloat(xs[i], 64)
		y, _ := strconv.ParseFloat(ys[i], 64)

		a := math.Abs(x - y)
		tmp += a * a * a
	}
	distance = math.Pow(tmp, 1.0/3.0)
	return
}

func chebychev(xs, ys []string, n int) (distance float64) {

	for i := 0; i < n; i++ {
		x, _ := strconv.ParseFloat(xs[i], 64)
		y, _ := strconv.ParseFloat(ys[i], 64)

		a := math.Abs(x - y)
		if distance < a {
			distance = a
		}
	}
	return
}
