package main

import (
	"bufio"
	"fmt"
	"os"
)

// Scane Standard Input
var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s := nextLine()

	result := make([]rune, len(s))

	for i, c := range s {
		result[i] = alphabetConverter(c)
	}

	fmt.Println(string(result))
}

// 文字コードの判定がわからず以下を参考にした
// https://hhelibex.hatenablog.jp/entry/2017/10/17/214438
func alphabetConverter(in rune) rune {

	if 'a' <= in && in <= 'z' {
		return in - ('a' - 'A')
	} else if 'A' <= in && in <= 'Z' {
		return in + ('a' - 'A')
	} else {
		return in
	}
}
