// https://leetcode.com/problems/third-maximum-number/
import (
	"sort"
)

func thirdMax(nums []int) int {
	// sort
	sort.Ints(nums)
	// remove duplicate
	removed := removeDuplicate(nums)

	index := len(removed) - 3
	if index < 0 {
		return removed[len(removed)-1]
	} else {
		return removed[index]
	}
}

func removeDuplicate(args []int) []int {
	results := make([]int, 0, len(args))
	encountered := map[int]bool{}
	for i := 0; i < len(args); i++ {
		if !encountered[args[i]] {
			encountered[args[i]] = true
			results = append(results, args[i])
		}
	}
	return results
}