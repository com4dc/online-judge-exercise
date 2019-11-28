// wrong answer 2019/11/04
func convertToTitle(n int) string {
	return loop(n, 0)
}

func loop(num, carry int) string {
	result := (num - 1) + 'A'

	if result > 'Z' {
		carry++
		result = result - 'Z'
	} else {
		if carry == 0 {
			return string(result)
		} else {
			carry = (carry - 1) + 'A'
			return string(carry) + string(result)
		}
	}
	return loop(result, carry)
}
