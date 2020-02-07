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
	c, _ := strconv.Atoi(nextLine())

	t_score := 0
	h_score := 0

	for i := 0; i < c; i++ {
		phase := strings.Split(nextLine(), " ")

		comp := strings.Compare(phase[0], phase[1])

		if comp > 0 {
			t_score += 3
		} else if comp == 0 {
			t_score++
			h_score++
		} else {
			h_score += 3
		}
	}
	fmt.Printf("%d %d\n", t_score, h_score)
}
