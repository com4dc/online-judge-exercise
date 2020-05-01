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

type Command struct {
	inst string
	val  int
}

type LinkedList struct {
	head int
	tail int
	list []int
}

func main() {
	n, _ := strconv.Atoi(nextLine())

	list := LinkedList{head: -1, tail: -1, list: make([]int, 0)}

	for i := 0; i < n; i++ {
		command := nextLine()

		c := parseCommand(command)

		switch c.inst {
		case "insert":
			insert(c.val, &list)
		case "delete":
			delete(c.val, &list)
		case "deleteFirst":
			deleteFirst(&list)
		case "deleteLast":
			deleteLast(&list)
		}
	}

	ans := make([]string, len(list.list))
	for idx := len(list.list) - 1; idx >= 0; idx-- {
		tmp := list.list[idx]

		ansIdx := len(list.list) - 1 - idx
		ans[ansIdx] = strconv.Itoa(tmp)
	}

	fmt.Println(strings.Join(ans, " "))
}

func parseCommand(command string) Command {
	c := strings.Split(command, " ")

	if len(c) == 2 {
		v, _ := strconv.Atoi(c[1])

		return Command{c[0], v}
	}
	return Command{c[0], -1}
}

func insert(n int, l *LinkedList) {
	l.list = append(l.list, n)
	l.head = n

	if l.tail == -1 {
		l.tail = n
	}
}

func delete(n int, l *LinkedList) {
	for idx := len(l.list) - 1; idx >= 0; idx-- {
		v := l.list[idx]

		if n == v {

			if idx == 0 {
				l.head = n
			} else if idx == len(l.list)-1 {
				l.tail = n
			}

			l.list = append(l.list[:idx], l.list[idx+1:]...)
			return
		}
	}
}

func deleteFirst(l *LinkedList) {

	if len(l.list) == 1 {
		l.tail = l.list[0]
		l.list = l.list[1:]
		return
	}
	l.head = l.list[len(l.list)-2]
	l.list = l.list[:len(l.list)-1]
}

func deleteLast(l *LinkedList) {

	if len(l.list) == 1 {
		l.tail = l.list[0]
		l.list = l.list[1:]
		return
	}

	l.tail = l.list[1]
	l.list = l.list[1:]
}
