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
	s := nextLine() // cyclic
	p := nextLine()

	slen := len(s)
	plen := len(p)

	// sの開始位置をシーク
	for i := 0; i < len(s); i++ {
		short := (slen - i) - plen

		mihon := ""
		if short > 0 {
			mihon = s[i : i+plen]
		} else {
			mihon = s[i:slen] + s[:-1*short]
		}

		if p == mihon {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
