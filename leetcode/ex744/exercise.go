// https://leetcode.com/problems/find-smallest-letter-greater-than-target/
func nextGreatestLetter(letters []byte, target byte) byte {
    result := target
    target = target % letters[len(letters) - 1]
    for _, value := range letters {

        if value > target {
            result = value
            break;
        }
    }
    return result
}
