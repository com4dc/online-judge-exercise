// https://leetcode.com/problems/reverse-vowels-of-a-string/
// this go program is tooooooooooo slooooooooooooooow
func reverseVowels(s string) string {
    index := []int{}
    chars := []string{}
    v := []string{"a", "i", "u", "e", "o", "A", "I", "U", "E", "O"}
    
    for i, c := range s {
        if contains(v, string(c)) == true {
            index = append(index, i)
            chars = append(chars, string(c))
        }
    }
    
    // 変換するものがない
    if len(index) == 0 {
        return s
    }
    
    result := ""
    fIdx := index[0]
    fChar := chars[len(chars)-1]                
        
    for i, c := range s {
        
        if i == fIdx {
            result = result + fChar
            index = index[1:len(index)]
            chars = chars[:len(index)]
        } else {
            result = result + string(c)
        }
        
        if (len(chars) > 0) {
            fIdx = index[0]
            fChar = chars[len(chars)-1]             
        }
    
    }
    
    return result
}


func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
