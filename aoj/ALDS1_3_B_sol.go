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

type Process struct {
	name string
	time int
}

const MAX_LEN = 100000

type Queue struct {
	container [MAX_LEN]Process
	head      int
	tail      int
}

func (q *Queue) isEmpty() bool {
	return q.head == q.tail
}

func (q *Queue) enqueue(name string, time int) {
	q.container[q.tail] = Process{name: name, time: time}
	q.tail = (q.tail + 1) % MAX_LEN // Overflowしないように最大値で余剰
}

func (q *Queue) dequeue() Process {
	if q.isEmpty() {
		fmt.Println("Empty")
	}
	process := q.container[q.head]
	q.head = (q.head + 1) % MAX_LEN // Overflowしないように最大値で余剰
	return process
}

func main() {
	q := Queue{head: 0, tail: 0}
	s := strings.Split(nextLine(), " ")

	n, _ := strconv.Atoi(s[0])
	time_slice, _ := strconv.Atoi(s[1])

	for i := 0; i < n; i++ {
		str := strings.Split(nextLine(), " ")
		name := str[0]
		time, _ := strconv.Atoi(str[1])

		q.enqueue(name, time)
	}

	cpu := 0

	for {
		if q.isEmpty() {
			break
		}

		processing := q.dequeue()
		if processing.time <= time_slice {
			cpu += processing.time
			fmt.Printf("%s %d\n", processing.name, cpu)
		} else {
			q.enqueue(processing.name, processing.time-time_slice)
			cpu += time_slice
		}
	}

}
