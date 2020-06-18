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
	n, _ := strconv.Atoi(nextLine())
	S := strings.Split(nextLine(), " ")

	q, _ := strconv.Atoi(nextLine())
	T := strings.Split(nextLine(), " ")

	count := 0
	for i := 0; i < q; i++ {
		tmp_S := make([]string, n)
		t := T[i]

		copy(tmp_S, S)

		if binarySearch(tmp_S, t) {
			count++
		}
	}
	fmt.Println(count)
}

func binarySearch(l []string, target string) bool {

	fmt.Printf("l : %s, target : %s\n", l, target)
	last_idx := len(l) - 1

	if len(l) == 0 {
		return false
	}

	half_idx := last_idx / 2

	t := l[half_idx]
	if t == target {
		fmt.Printf("Hit Target : %s\n", target)
		return true
	} else {
		t_n, _ := strconv.Atoi(t)
		target_n, _ := strconv.Atoi(target)

		a := make([]string, 0)
		if t_n > target_n {
			a = deepcopy(l[:half_idx])
		} else {
			a = deepcopy(l[half_idx+1:])
		}

		if len(a) == 1 {
			fmt.Printf("length1 l = %s, target = %s\n", a, target)
			return a[0] == target
		}

		// time.Sleep(time.Second * 1)

		return binarySearch(a, target)
	}
}

func deepcopy(source []string) []string {
	dst := make([]string, 0)
	for _, s := range source {
		dst = append(dst, s)
	}
	return dst
}
