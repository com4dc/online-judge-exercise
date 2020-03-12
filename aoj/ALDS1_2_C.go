package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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
	N, _ := strconv.Atoi(nextLine())
	s_line := strings.Split(nextLine(), " ")

	C := make([]Card, 0)
	for _, s := range s_line {
		sign := s
		value, _ := strconv.Atoi(string(s[1]))
		c := Card{sign: sign, value: value}
		C = append(C, c)
	}

	sort_b := make([]Card, len(C))
	copy(sort_b, C)
	sort_b, _ = bubbleSort(sort_b, N)

	sort_s := make([]Card, len(C))
	copy(sort_s, C)
	sort_s, _ = selectionSort(sort_s, N)

	m_origin := same_value(C)
	b_origin := signed_order(sort_b)

	stable := true
	for k, v := range b_origin {
		s := m_origin[k]
		// fmt.Printf("s=%s, v=%s\n", strings.Join(s, " "), strings.Join(v, " "))
		if !reflect.DeepEqual(s, v) {
			stable = false
			break
		}
	}

	fmt.Println(stringify(sort_b))
	if stable {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}

	s_origin := signed_order(sort_s)
	stable = true
	for k, v := range s_origin {
		s := m_origin[k]
		// fmt.Printf("s=%s, v=%s\n", strings.Join(s, " "), strings.Join(v, " "))
		if !reflect.DeepEqual(s, v) {
			stable = false
			break
		}
	}

	fmt.Println(stringify(sort_s))
	if stable {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
}

type Card struct {
	sign  string
	value int
}

func stringify(C []Card) string {
	r := make([]string, len(C))
	for i, v := range C {
		r[i] = v.sign
	}
	return strings.Join(r, " ")
}

func bubbleSort(C []Card, N int) ([]Card, int) {
	count := 0
	flag := 1

	for flag == 1 {
		flag = 0
		for j := N - 1; j > 0; j-- {
			if C[j].value < C[j-1].value {
				tmp := C[j]
				C[j] = C[j-1]
				C[j-1] = tmp
				flag = 1
				count++
			}
		}
	}
	return C, count
}

func selectionSort(C []Card, N int) ([]Card, int) {

	count := 0
	for i := 0; i < N; i++ {
		minj := i
		for j := i; j < N; j++ {
			if C[j].value < C[minj].value {
				minj = j
			}
		}

		// changed value i
		if i != minj {
			count++
		}

		tmp := C[i]
		C[i] = C[minj]
		C[minj] = tmp
	}

	return C, count
}

func signed_order(C []Card) map[int][]string {

	m := make(map[int][]string, 0)

	prev := C[0]
	for i := 1; i < len(C); i++ {
		if prev.value == C[i].value {
			if len(m[C[i].value]) == 0 {
				m[C[i].value] = append(m[C[i].value], prev.sign)
			}
			m[C[i].value] = append(m[C[i].value], C[i].sign)
		}
		prev = C[i]
	}
	return m
}

func same_value(C []Card) map[int][]string {
	m := map[int][]string{}

	for _, v := range C {
		m[v.value] = append(m[v.value], v.sign)
	}
	return m
}
