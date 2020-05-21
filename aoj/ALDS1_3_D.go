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

type Stack struct {
	stack []int
}

func (stack *Stack) push(val int) {
	stack.stack = append(stack.stack, val)
}

func (stack *Stack) pop() int {
	if len(stack.stack) == 0 {
		return -1
	}

	val := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[0 : len(stack.stack)-1]
	return val
}

func (stack *Stack) empty() bool {
	if len(stack.stack) == 0 {
		return true
	}
	return false
}

type Stack2 struct {
	stack []Tuple
}

type Tuple struct {
	index int
	area  int
}

func (stack *Stack2) push(val Tuple) {
	stack.stack = append(stack.stack, val)
}

func (stack *Stack2) pop() Tuple {
	if len(stack.stack) == 0 {
		return Tuple{-1, 0}
	}

	val := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[0 : len(stack.stack)-1]
	return val
}

func (stack *Stack2) top() Tuple {
	if len(stack.stack) == 0 {
		return Tuple{-1, 0}
	}

	val := stack.stack[len(stack.stack)-1]
	return val
}

func (stack *Stack2) empty() bool {
	if len(stack.stack) == 0 {
		return true
	}
	return false
}

func main() {
	line := nextLine()

	s1 := Stack{}
	s2 := Stack2{}
	total := 0 // 面積の総和
	for i, r := range line {
		rr := string(r)

		switch rr {
		case "\\":
			s1.push(i) // indexをStackに追加
		case "/":
			b := s1.pop()
			if b == -1 {
				continue
			}

			area := i - b
			total += area

			if !s2.empty() {

				if s2.top().index < b {
					s2.push(Tuple{b, area})
				} else {
					t2 := s2.pop()
					new_area := area

					for t2.index > b {
						new_area += t2.area
						t2 = s2.pop()
					}
					s2.push(t2)
					s2.push(Tuple{b, new_area})
				}
			} else {
				s2.push(Tuple{b, area})
			}
		default:
			// no op
		}
	}

	fmt.Println(total)

	out := make([]string, 0)
	if len(s2.stack) == 0 {
		fmt.Println(0)
		return
	}
	for _, a := range s2.stack {
		if a.area == 0 {
			continue
		}
		out = append(out, strconv.Itoa(a.area))
	}
	fmt.Printf("%d %s\n", len(out), strings.Join(out, " "))
}
