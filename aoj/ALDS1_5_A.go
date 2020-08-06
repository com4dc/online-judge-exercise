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
	_ = nextLine()
	A := convert(strings.Split(nextLine(), " "))

	q, _ := strconv.Atoi(nextLine())
	m := convert(strings.Split(nextLine(), " "))

	for i := 0; i < q; i++ {
		if exhaustiveSearch(m[i], A) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

// 再起処理を起動
func exhaustiveSearch(x int, A []int) bool {
	return solve(x, 0, A)
}

// 再起で回して検索
func solve(target int, index int, A []int) bool {
	if target == 0 {
		return true
	} else if target < 0 || index >= len(A) {
		return false
	} else {
		return solve(target-A[index], index+1, A) ||
			solve(target, index+1, A)
	}
}

func convert(src []string) []int {
	r := make([]int, len(src))
	for i, v := range src {
		r[i], _ = strconv.Atoi(v)
	}
	return r
}
