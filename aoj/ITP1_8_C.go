package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Scane Standard Input
var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	r := make(map[string]int)
	r = initMap(r)

	for {
		s := nextLine()
		if s == "" {
			break
		}

		for _, ss := range s {
			if ('a' <= ss && ss <= 'z') || ('A' <= ss && ss <= 'Z') {
				tmp := strings.ToLower(string(ss))
				r[tmp] = r[tmp] + 1
			}
		}
	}
	a := List{}
	for k, v := range r {
		e := Entry{k, v}
		a = append(a, e)
	}
	sort.Sort(a)

	for _, e := range a {
		fmt.Println(e.key + " : " + strconv.Itoa(e.value))
	}
}

func initMap(m map[string]int) map[string]int {
	for i := 0; i < 26; i++ {
		m[string(rune('a'+i))] = 0
	}
	return m
}

// 以下Keyをソートするためのコード
type Entry struct {
	key   string
	value int
}

type List []Entry

func (l List) Len() int {
	return len(l)
}
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
func (l List) Less(i, j int) bool {
	return (l[i].key < l[j].key)
}
