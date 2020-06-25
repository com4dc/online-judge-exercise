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

	// 内部で回す関数を定義
	var reSearch func(b, e int64) bool
	reSearch = func(b, e int64) bool {
		// startとendのindexが合致、または逆転してれば評価終了。該当なし
		if b >= e {
			return false
		}
		// startとendの中間点を計算
		m := (b + e) / 2
		// 合致すれば true
		if arr[m] == x {
			return true
		}
		// 対象よりも低い数ならmの次のindexからendまでで再計算
		if arr[m] < x {
			return reSearch(m+1, e)
		}
		// 対象よりも大きな数ならbからmまでで再計算
		return reSearch(b, m)
	}
	// 開始は0から最大値indexまで
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
	// 比較元の Array を構成
	n := nextInt()
	sarr := make([]int64, n)
	for i, _ := range sarr {
		sarr[i] = nextInt()
	}

	// 比較対象の Array を構成
	q := nextInt()
	tarr := make([]int64, q)
	for i, _ := range tarr {
		tarr[i] = nextInt()
	}

	// 比較対象を回しながらバイナリサーチ
	count := 0 // true の場合インクリメント
	for _, t := range tarr {
		if binarySearch(t, n, sarr) {
			count++
		}
	}

	fmt.Println(count)
}
