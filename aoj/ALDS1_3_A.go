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
	lines := strings.Split(nextLine(), " ")
	r := process(lines)
	fmt.Println(r)
}

func process(lines []string) string {

	queue := make([]string, 0)
	new_stack := &Stack{queue}

	for _, element := range lines {

		switch element {
		case "+":
			a, _ := strconv.Atoi(new_stack.Pop())
			b, _ := strconv.Atoi(new_stack.Pop())
			new_stack = new_stack.Push(strconv.Itoa(b + a))
		case "-":
			a, _ := strconv.Atoi(new_stack.Pop())
			b, _ := strconv.Atoi(new_stack.Pop())
			new_stack = new_stack.Push(strconv.Itoa(b - a))
		case "*":
			a, _ := strconv.Atoi(new_stack.Pop())
			b, _ := strconv.Atoi(new_stack.Pop())
			new_stack = new_stack.Push(strconv.Itoa(b * a))
		case "/":
			a, _ := strconv.Atoi(new_stack.Pop())
			b, _ := strconv.Atoi(new_stack.Pop())
			new_stack = new_stack.Push(strconv.Itoa(b / a))
		default:
			new_stack = new_stack.Push(element)
		}
		// fmt.Printf("%d, %s\n", result, element)
	}
	return new_stack.Pop()
}

type Stack struct {
	stack []string
}

func (s *Stack) Pop() string {
	last_idx := len(s.stack) - 1
	// get last element
	p := s.stack[last_idx]
	// update stack
	n := make([]string, last_idx)
	copy(n, s.stack)
	s.stack = n
	return p
}

func (s *Stack) Push(e string) *Stack {
	s.stack = append(s.stack, e)
	return s
}
