import ("strconv")


func calPoints(ops []string) int {
    result := make([]int, 0) // 結果
    
    for _, value := range ops {
        switch value {
            case "D":
                // get last and Dounble and Add
                tmp := result[len(result) - 1]
                result = append(result, tmp * 2)
            case "C":
                // remove last
                result = result[:len(result) - 1]
            case "+":
                // sum last 2 valid value
                last1 := result[len(result) - 1]
                last2 := result[len(result) - 2]
                result = append(result, last1 + last2)
            default:
                i, _ := strconv.Atoi(value)
                result = append(result, i)
        }
    }
    
    sum := 0
    for _, x := range result {
        sum += x
    }
    
    return sum
}


