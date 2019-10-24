// https://leetcode.com/problems/flipping-an-image/
func flipAndInvertImage(A [][]int) [][]int {
    
    result := make([][]int, 0)
        
    for _, array := range( A ){
        temp := make([]int, 0)
        for _, value := range( array ) {
           temp = append(temp, value)
        }
        result = append(result, reverse(temp))
    }
    return result
}

func reverse(tmp []int) []int {
    rr := make([]int, 0)
    for i := len(tmp) - 1; i >= 0; i-- {
        rr = append(rr, (tmp[i] + 1)%2)
    }
    return rr
}

