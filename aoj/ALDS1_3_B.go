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

type Process struct {
	name   string
	time   int
	wasted int
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	dfns := strings.Split(nextLine(), " ")
	C, _ := strconv.Atoi(dfns[0])
	LIMIT, _ := strconv.Atoi(dfns[1])

	q := createQueue(C)
	ends := queueProcess(q, LIMIT)

	for i := 0; i < len(ends); i++ {
		qq := ends[i]
		fmt.Printf("%s %d\n", qq.name, qq.time)
	}
}

func createQueue(C int) []Process {
	queue := make([]Process, C)
	for i := 0; i < C; i++ {
		lines := strings.Split(nextLine(), " ")

		p := lines[0]
		t, _ := strconv.Atoi(lines[1])
		queue[i] = Process{p, t, t}
	}
	return queue
}

func queueProcess(queue []Process, LIMIT int) []Process {

	end := make([]Process, 0)
	cpu := 0

	for {
		// queue が 0 になったら終了
		if len(queue) <= 0 {
			return end
		}

		// Queueの先頭を取得
		p := queue[0]

		if p.wasted <= LIMIT {
			cpu += p.wasted
			p.time = cpu
			end = append(end, p)

			// 先頭を削除してQueueを作り直す
			copied := make([]Process, len(queue)-1)
			copy(copied, queue[1:])
			queue = copied
		} else {
			cpu += LIMIT
			p.wasted = p.wasted - LIMIT
			queue = append(queue, p)
			queue = queue[1:]
		}

	}
}
