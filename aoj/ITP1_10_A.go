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

	x1, _ := strconv.ParseFloat(num_s[0], 64)
	y1, _ := strconv.ParseFloat(num_s[1], 64)
	x2, _ := strconv.ParseFloat(num_s[2], 64)
	y2, _ := strconv.ParseFloat(num_s[3], 64)

	r := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
	fmt.Printf("%.8f\n", r)
}
