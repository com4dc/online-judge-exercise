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
	_, _ = strconv.Atoi(nextLine())
	S_seq := strings.Split(nextLine(), " ")

	q, _ := strconv.Atoi(nextLine())
	T_seq := strings.Split(nextLine(), " ")

	count := 0
	cvt := convertIntSeq(S_seq)

	for i := 0; i < q; i++ {
		target, _ := strconv.Atoi(T_seq[i])

		res, tmp := getAndRemove(cvt, target)
		copy(cvt, tmp)
		if res {
			count++
		}
	}
	fmt.Println(count)
}

func convertIntSeq(seq []string) []int {
	s := make([]int, len(seq))
	for i, r := range seq {
		s[i], _ = strconv.Atoi(r)
	}
	return s
}

func getAndRemove(seq []int, target int) (bool, []int) {
	for idx, s := range seq {
		// fmt.Printf("s=%d, target=%d\n", s, target)
		if target == s {
			res := make([]int, 0)

			// fmt.Printf("bingo! idx=%d, s=%d, target=%d\n", idx, s, target)

			if idx-1 >= 0 {
				res = append(res, seq[:idx]...)
				res = append(res, seq[idx+1:]...)
			} else {
				res = append(res, seq[idx+1:]...)
			}
			return true, res
		}
	}
	return false, seq
}
