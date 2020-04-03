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
	n, _ := strconv.Atoi(nextLine())

	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = nextLine()
	}

	// fmt.Println(strings.Join(a, ","))

	hoge := shellSort(a, n)
	fmt.Println(strings.Join(hoge, ","))
}

func insertionSort(A []string, n, g int) []string {
	r := make([]string, len(A))
	copy(r, A)
	for i := g; i < n-1; i++ {
		v, _ := strconv.Atoi(r[i])
		j := i - g

		for {
			rv, _ := strconv.Atoi(r[j])
			if j >= 0 && rv > v {
				r[j+g] = r[j]
				j = j - g
			} else {
				break
			}
		}
		r[j+g] = strconv.Itoa(v)
	}
	return r
}

func shellSort(A []string, n int) []string {

	count := 0
	m := len(A)

	result := make([]string, len(A))
	copy(result, A)

	G := make([]int, 0)

	i := 2
	for {
		if m < 1 {
			break
		}
		m = m / i
		G = append(G, m)
		fmt.Println(m)
	}

	fmt.Printf("%d\n", len(G))

	for i := 0; i < m-1; i++ {
		count++
		result = insertionSort(result, n, G[i])
	}

	return result
}
