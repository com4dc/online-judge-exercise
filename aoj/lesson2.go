package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := inputIO()
	res := strings.Split(s, " ")

	a, errA := strconv.Atoi(res[0])
	b, errB := strconv.Atoi(res[1])

	if errA != nil || errB != nil {
		panic("Atoi err")
	}

	fmt.Println(a*b, a*2+b*2)
}

func inputIO() (result string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	result = scanner.Text()
	return
}
