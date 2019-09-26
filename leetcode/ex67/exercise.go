// https://leetcode.com/problems/add-binary/
import "strconv"

func addBinary(a string, b string) string {

	a1, b1 := Reverse(a), Reverse(b)
	result := ""
	cl := 0

	for i, j := 0, 0; i < len(a1) || j < len(b1); i, j = i+1, j+1 {
		if i < len(a1) && j < len(b1) {
			ar := ByteToInt(a1[i])
			br := ByteToInt(b1[i])
			val := ar + br + cl
			result = strconv.Itoa(val%2) + result
			cl = val / 2
		} else {
			if i < len(a1) {
				aa := ByteToInt(a1[i])
				val1 := aa + cl
				cl = val1 / 2
				result = strconv.Itoa(val1%2) + result
			}
			if j < len(b1) {
				bb := ByteToInt(b1[i])
				val2 := bb + cl
				cl = val2 / 2
				result = strconv.Itoa(val2%2) + result
			}
		}
	}
	if cl > 0 {
		return strconv.Itoa(cl) + result
	} else {
		return result
	}
}

func ByteToInt(a byte) (res int) {
	tmp := string(a)
	res, _ = strconv.Atoi(tmp)
	return
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
