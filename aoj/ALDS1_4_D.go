package main

// typed from
// https://onlinejudge.u-aizu.ac.jp/solutions/problem/ALDS1_4_D/review/2985436/nmurakami/Go
// TLE出る

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
	ls := strings.Split(nextLine(), " ")
	n, _ := strconv.Atoi(ls[0])
	k, _ := strconv.Atoi(ls[1])

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(nextLine())
		arr[i] = a
	}

	e := calcMinTruckSize(k, arr)
	fmt.Println(e)
}

func loadable(truckNum, truckSize int, elem []int) bool {

	truckCount := 0
	sum := 0
	for i := 0; i < len(elem); i++ {
		if sum+elem[i] > truckSize {
			truckCount++

			// truckCount がトラックの数を超えていたらfalse
			if truckCount >= truckNum {
				return false
			}
			sum = 0
		}
		sum += elem[i] // 積載していく
	}
	// トラック内に収まればOK
	return true
}

//  最小のトラックサイズを計算する
func calcMinTruckSize(numTrucks int, elem []int) int {
	low, high := 0, 0
	for i := 0; i < len(elem); i++ {
		if low < elem[i] {
			low = elem[i] // low は最小の値が設定される
		}
		high += elem[i] // high はひたすら要素をたしていく

		for low+1 <= high {
			mid := (low + high) / 2 // 中間値
			if loadable(numTrucks, mid, elem) {
				high = mid
			} else {
				low = mid + 1 // mid の +1 した値を low に指定
			}
		}
	}
	return high
}
