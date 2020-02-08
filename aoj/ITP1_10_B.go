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
	num_s := strings.Split(nextLine(), " ")

	a, _ := strconv.ParseFloat(num_s[0], 64)
	b, _ := strconv.ParseFloat(num_s[1], 64)
	d, _ := strconv.ParseFloat(num_s[2], 64)

	r := d / 180.0 * math.Pi

	c := math.Sqrt(a*a + b*b - 2.0*a*b*math.Cos(r))

	h := b * math.Sin(r)

	S := a * h / 2.0
	L := a + b + c

	fmt.Printf("%.8f\n", S)
	fmt.Printf("%.8f\n", L)
	fmt.Printf("%.8f\n", h)
}
