// https://leetcode.com/problems/majority-element/

func majorityElement(nums []int) int {
	m := map[int]int{}
	max := 0
	res := 0

	for _, num := range nums {
		m[num] = m[num] + 1

		if m[num] > max {
			max = m[num]
			res = num
		}
	}
	return res
}
