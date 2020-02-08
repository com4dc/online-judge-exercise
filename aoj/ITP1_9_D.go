package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Scan Standard Input
var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	str := nextLine()
	n, _ := strconv.Atoi(nextLine())

	for i := 0; i < n; i++ {
		m := strings.Split(nextLine(), " ")
		switch m[0] {
		case "print":
			{
				s_idx, _ := strconv.Atoi(m[1])
				e_idx, _ := strconv.Atoi(m[2])
				fmt.Println(str[s_idx : e_idx+1])
			}
		case "replace":
			{
				s_idx, _ := strconv.Atoi(m[1])
				e_idx, _ := strconv.Atoi(m[2])
				str = str[:s_idx] + m[3] + str[e_idx+1:]
			}
		case "reverse":
			{
				s_idx, _ := strconv.Atoi(m[1])
				e_idx, _ := strconv.Atoi(m[2])
				tmp := []rune(str[s_idx : e_idx+1])

				tt := ""
				for j := len(tmp) - 1; j >= 0; j-- {
					tt = tt + string(tmp[j])
				}
				str = str[:s_idx] + tt + str[e_idx+1:]
			}
		default:
		}
	}
}
