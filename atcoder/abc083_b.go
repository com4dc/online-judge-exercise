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

	in := strings.Split(nextLine(), " ")

	N, _ := strconv.Atoi(in[0])
	A, _ := strconv.Atoi(in[1])
	B, _ := strconv.Atoi(in[2])

	result := 0

	for i := 0; i <= N; i++ {

		digit := calcDigit(i)
		r := calc(i, digit)

		if r >= A && r <= B {
			result += i
		}
	}
	fmt.Println(result)
}

func calcDigit(x int) (digit int) {
	digit = len(strconv.Itoa(x))
	return
}

func calc(n, digit int) (result int) {
	for i := 0; i < digit; i++ {
		div := int(math.Pow10(i))
		result += (n / div) % 10
	}
	return
}
