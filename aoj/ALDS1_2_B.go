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
	N, _ := strconv.Atoi(nextLine())
	A := strings.Split(nextLine(), " ")

	r, count := selectionSort(A, N)
	fmt.Println(strings.Join(r, " "))
	fmt.Println(count)
}

func selectionSort(A []string, N int) ([]string, int) {
	tmpSeq := make([]int, len(A))
	count := 0
	for i, v := range A {
		n, _ := strconv.Atoi(v)
		tmpSeq[i] = n
	}

	for i := 0; i < N; i++ {
		minj := i
		for j := i; j < N; j++ {
			if tmpSeq[j] < tmpSeq[minj] {
				minj = j
			}
		}

		// changed value i
		if i != minj {
			count++
		}

		tmp := tmpSeq[i]
		tmpSeq[i] = tmpSeq[minj]
		tmpSeq[minj] = tmp
	}

	result := make([]string, len(A))
	for i, v := range tmpSeq {
		result[i] = strconv.Itoa(v)
	}
	return result, count
}
