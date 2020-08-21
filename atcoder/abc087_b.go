package main

import (
	"bufio"
	"fmt"
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
	A, _ := strconv.Atoi(nextLine()) // 500
	B, _ := strconv.Atoi(nextLine()) // 100
	C, _ := strconv.Atoi(nextLine()) // 50

	X, _ := strconv.Atoi(nextLine()) // 合計

	fmt.Println(check(A, B, C, X))
}

// 頭悪すぎる
func check(a, b, c, x int) (count int) {

	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {

				ta := 500 * i
				tb := 100 * j
				tc := 50 * k

				if (ta + tb + tc) == x {
					count++
				}
			}
		}
	}
	return
}
