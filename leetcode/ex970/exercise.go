// https://leetcode.com/problems/powerful-integers

import (
	"math"
)

func powerfulIntegers(x int, y int, bound int) []int {
	x1 := make([]int, 0, 0)
	i := 0
	for {
		t := int(math.Pow(float64(x), float64(i)))

		// 0乗は1なので必ず最大値よりは1小さい値が限界値のはず
		// それと同じ数値が出た場合ループしてるので打ち切る
		if t > bound-1 || contains(x1, t) {
			break
		}
		x1 = append(x1, t)
		i++
	}

	x2 := make([]int, 0, 0)
	j := 0
	for {
		t := int(math.Pow(float64(y), float64(j)))

		// 0乗は1なので必ず最大値よりは1小さい値が限界値のはず
		// それと同じ数値が出た場合ループしてるので打ち切る
		if t > bound-1 || contains(x2, t) {
			break
		}
		x2 = append(x2, t)
		j++
	}

	result := make([]int, 0, 0)
	for k := 0; k < len(x1); k++ {
		for l := 0; l < len(x2); l++ {
			// 足し算組み合わせ。限界値を超えたら打ち切る
			temp := x1[k] + x2[l]
			if temp > bound {
				break
			}

			if contains(result, temp) == true {
				continue
			} else {
				result = append(result, temp)
			}
		}
	}

	return result
}

// 配列内に同じ要素があるかどうかをチェック
func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
