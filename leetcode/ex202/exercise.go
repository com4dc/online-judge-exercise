// https://leetcode.com/problems/happy-number/

import (
	"math"
	"strconv"
)

func isHappy(n int) bool {

	m := map[int]int{}

	s := n
	for s >= 0 {
		res := splitNumber(s)
		tmp := 0
		for _, value := range res {
			tmp += value * value
		}
		if tmp == 1 {
			return true
		}
		s = tmp

		_, exist := m[tmp]
		if exist {
			return false
		} else {
			m[tmp] = 1
		}

	}
	return false
}

func splitNumber(n int) (result []int) {
	count := len(strconv.Itoa(n))
	result = make([]int, count)

	if count == 1 {
		result[0] = 0
		result = append(result, n%10)
		return
	}

	for i := 0; i < count; i++ {
		if i < count-1 {
			tmp := int(math.Pow(10, float64(count-(i+1))))
			result[i] = n / tmp
			n = n % tmp
		} else {
			result[i] = n % 10
		}
	}
	return
}
