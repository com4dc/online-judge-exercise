// https://leetcode.com/problems/detect-capital/
import "unicode"

func detectCapitalUse(word string) bool {

	upper, upper2, result := true, true, false
	for i := 0; i < len(word); i++ {

		if i == 0 {
			upper = unicode.IsUpper(rune(word[i]))

			if len(word) == 1 {
				return true
			}
		} else if i == 1 {
			upper2 = unicode.IsUpper(rune(word[i]))

			if len(word) == 2 {
				return upper == upper2 || (upper == true && upper2 == false)
			}
		} else {
			if upper == true && upper2 == true && unicode.IsUpper(rune(word[i])) {
				result = true
			} else if upper == true && upper2 == false && unicode.IsLower(rune(word[i])) {
				result = true
			} else if upper == false && upper2 == false && unicode.IsLower(rune(word[i])) {
				result = true
			} else {
				return false
			}
		}
	}
	return result
}