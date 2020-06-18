package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int64 {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return int64(n)
}

func binarySearch(x, n int64, arr []int64) bool {

	var reSearch func(b, e int64) bool

	reSearch = func(b, e int64) bool {
		if b >= e {
			return false
		}
		m := (b + e) / 2
		if arr[m] == x {
			return true
		}
		if arr[m] < x {
			return reSearch(m+1, e)
		}
		return reSearch(b, m)
	}

	return reSearch(0, n)
}

func parseArrInt(l int64, strArr []string) []int64 {
	var intArr []int64

	for i := 0; i < len(strArr); i++ {
		n, _ := strconv.Atoi(strArr[i])
		intArr = append(intArr, int64(n))

	}
	return intArr
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	sarr := make([]int64, n)
	for i, _ := range sarr {
		sarr[i] = nextInt()
	}

	q := nextInt()
	tarr := make([]int64, q)
	for i, _ := range tarr {
		tarr[i] = nextInt()
	}

	count := 0
	for _, t := range tarr {
		if binarySearch(t, n, sarr) {
			count++
		}
	}

	fmt.Println(count)
}
