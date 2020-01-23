package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func createNewSheet(a, b int) (sheet [][]string) {
	sheet = make([][]string, a+1)
	for i := 0; i < a+1; i++ {
		sheet[i] = make([]string, b+1)
	}
	return
}

func outputArray(sheet []string) {
	fmt.Println(strings.Join(sheet[:], " "))
}

func main() {
	s := nextLine()

	res := strings.Split(s, " ")

	a, _ := strconv.Atoi(res[0])
	b, _ := strconv.Atoi(res[1])

	// 多重配列の初期化
	sheet := createNewSheet(a, b)

	for i := 0; i < a; i++ {
		line := nextLine()
		lines := strings.Split(line, " ")

		r := sheet[i]

		total := 0
		for j := 0; j < b; j++ {
			v, _ := strconv.Atoi(lines[j])

			r[j] = strconv.Itoa(v)

			total = total + v
		}
		r[b] = strconv.Itoa(total)
		outputArray(r)
	}

	// 足し合わせた数があるので+1行+1列多い
	for j := 0; j <= b; j++ {
		total := 0
		for i := 0; i <= a; i++ {
			v, _ := strconv.Atoi(sheet[i][j])
			total = total + v
		}
		sheet[a][j] = strconv.Itoa(total)
	}
	outputArray(sheet[a])
}
