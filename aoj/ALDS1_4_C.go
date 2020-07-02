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

	result := map[string]int{}

	arr := make([]string, n)
	for i := 0; i < n; i++ {
		arr[i] = nextLine()
	}

	// instruction loop
	for _, v := range arr {
		parseCommand(v, result)
	}
}

func parseCommand(line string, m map[string]int) {
	ls := strings.Split(line, " ")
	inst := ls[0]
	value := ls[1]

	switch inst {
	case "insert":
		insert(value, m)
	case "find":
		find(value, m)
	}
}

func insert(value string, m map[string]int) {
	m[value] = 1
}

func find(value string, m map[string]int) {
	_, ok := m[value]
	if ok {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
