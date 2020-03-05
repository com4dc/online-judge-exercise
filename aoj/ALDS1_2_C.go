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
	s_line := strings.Split(nextLine(), " ")

	C := make([]Card, 0)
	for _, s := range s_line {
		sign := s
		value, _ := strconv.Atoi(string(s[1]))
		c := Card{sign: sign, value: value}
		C = append(C, c)
	}

	C, _ = bubbleSort(C, N)

	result := make([]string, len(C))
	for idx, s := range C {
		result[idx] = s.sign
	}

	fmt.Println(strings.Join(result, " "))
}

type Card struct {
	sign  string
	value int
}

func bubbleSort(C []Card, N int) ([]Card, int) {
	count := 0
	flag := 1

	for flag == 1 {
		flag = 0
		for j := N - 1; j > 0; j-- {
			if C[j].value < C[j-1].value {
				tmp := C[j]
				C[j] = C[j-1]
				C[j-1] = tmp
				flag = 1
				count++
			}
		}
	}
	return C, count
}

func selectionSort(C []Card, N int) ([]Card, int) {

	count := 0
	for i := 0; i < N; i++ {
		minj := i
		for j := i; j < N; j++ {
			if C[j].value < C[minj].value {
				minj = j
			}
		}

		// changed value i
		if i != minj {
			count++
		}

		tmp := C[i]
		C[i] = C[minj]
		C[minj] = tmp
	}

	return C, count
}

func stable?(C []Card) {

}