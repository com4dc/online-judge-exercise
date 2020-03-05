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
	N, err1 := strconv.Atoi(nextLine())
	if err1 != nil {
		panic("convert error")
	}

	s := strings.Split(nextLine(), " ")
	A := convertI(s)

	// sort
	R, count := bubbleSort(A, N)

	// output result
	fmt.Println(strings.Join(convertA(R), " "))
	fmt.Println(count)
}

func convertA(target []int) (r []string) {
	for _, v := range target {
		n := strconv.Itoa(v)
		r = append(r, n)
	}
	return
}

func convertI(target []string) (r []int) {
	for _, v := range target {
		n, _ := strconv.Atoi(v)
		r = append(r, n)
	}
	return
}

func bubbleSort(A []int, N int) ([]int, int) {
	flag := 1
	count := 0
	for flag == 1 {
		flag = 0
		for j := N - 1; j > 0; j-- {
			if A[j] < A[j-1] {
				tmp := A[j]
				A[j] = A[j-1]
				A[j-1] = tmp
				flag = 1
				count++
			}
		}
	}
	return A, count
}
