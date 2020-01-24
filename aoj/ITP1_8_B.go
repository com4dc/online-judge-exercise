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
		s := nextLine()
		// 0だったら処理しない
		if s == "0" {
			break
		}

		// 1文字ずつに分解
		result := 0
		for _, ss := range s {
			i, _ := strconv.Atoi(string(ss))
			result = result + i
		}

		fmt.Println(strconv.Itoa(result))
	}
}
