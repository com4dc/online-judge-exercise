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
	a, _ := strconv.Atoi(nextLine())
	second := nextLine()
	bc := strings.Split(second, " ")
	b, _ := strconv.Atoi(bc[0])
	c, _ := strconv.Atoi(bc[1])
	s := nextLine()

	fmt.Printf("%d %s\n", a+b+c, s)
}
