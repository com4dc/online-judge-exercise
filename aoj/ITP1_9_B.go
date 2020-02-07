package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Scane Standard Input
var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {

	for {
		target := nextLine()

		if target == "-" {
			return
		}

		count, _ := strconv.Atoi(nextLine())

		for i := 0; i < count; i++ {
			h, _ := strconv.Atoi(nextLine())
			a := target[:h]
			b := target[h:]
			target = b + a
		}
		fmt.Println(target)
	}
}
