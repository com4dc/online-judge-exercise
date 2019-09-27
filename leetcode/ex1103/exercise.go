// https://leetcode.com/problems/distribute-candies-to-people
func distributeCandies(candies int, num_people int) []int {
	result := make([]int, num_people, 1000)
	for i, rest := 0, candies; rest > 0; i++ {
		rest = rest - (i + 1)
		index := i % num_people
		if rest >= 1 {
			result[index] = result[index] + (i + 1)
		} else {
			result[index] = result[index] + rest + (i + 1)
		}
	}
	return result
}
