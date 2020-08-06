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
	_ = nextLine()
	ss := convert(strings.Split(nextLine(), " "))

	g := mergeSort(ss, 0, len(ss))

	fmt.Println(strings.Join(intArrayToStringArray(ss), " "))
	fmt.Println(g)
}

func convert(r []string) []int {
	aa := make([]int, len(r))
	for i, v := range r {
		aa[i], _ = strconv.Atoi(v)
	}
	return aa
}

func intArrayToStringArray(r []int) []string {
	aa := make([]string, len(r))
	for i, v := range r {
		aa[i] = strconv.Itoa(v)
	}
	return aa
}

func merge(A []int, left, mid, right int) (count int) {
	n1 := mid - left
	n2 := right - mid

	L := make([]int, n1+1)
	for i := 0; i < n1; i++ {
		L[i] = A[left+i]
	}
	L[n1] = math.MaxInt64

	R := make([]int, n2+1)
	for i := 0; i < n2; i++ {
		R[i] = A[mid+i]
	}
	R[n2] = math.MaxInt64

	i := 0
	j := 0
	for k := left; k < right; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i = i + 1
		} else {
			A[k] = R[j]
			j = j + 1
		}
		count++
	}
	return
}

func mergeSort(A []int, left, right int) int {
	count := 0
	if left+1 < right {
		mid := (left + right) / 2
		count += mergeSort(A, left, mid)
		count += mergeSort(A, mid, right)
		count += merge(A, left, mid, right)
	}
	return count
}
