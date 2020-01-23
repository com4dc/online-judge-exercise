package main

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

func main() {
	s := nextLine()

	res := strings.Split(s, " ")

	i1, _ := strconv.Atoi(res[0])
	i2, _ := strconv.Atoi(res[1])
	i3, _ := strconv.Atoi(res[2])

	arr1, arr2 := generateArray(i1, i2, i3)

	arr1 = inputArray(arr1, i1, i2)
	arr2 = inputArray(arr2, i2, i3)

	result := productMatrix(arr1, arr2, i1, i3)

	// print
	outputMatrix(result)
}

func outputMatrix(result [][]string) {
	size := len(result)

	for i := 0; i < size; i++ {
		fmt.Println(strings.Join(result[i][:], " "))
	}
}

func productMatrix(arr1, arr2 [][]int, i1, i3 int) (result [][]string) {
	// result初期化
	result = make([][]string, i1)
	for i := 0; i < i1; i++ {
		result[i] = make([]string, i3)
	}

	for i := 0; i < i1; i++ {
		for j := 0; j < i3; j++ {
			// resultを作るためのループ
			result[i][j] = strconv.Itoa(decompose(arr1[i], arr2, j))
		}
	}
	return
}

func decompose(arr1 []int, arr2 [][]int, i int) int {
	add := 0
	size := len(arr1)

	for j := 0; j < size; j++ {
		add = add + arr1[j]*arr2[j][i]
	}
	return add
}

func inputArray(arr [][]int, i1, i2 int) [][]int {
	for i := 0; i < i1; i++ {
		input := nextLine()
		ss := strings.Split(input, " ")

		for j := 0; j < i2; j++ {
			hoge, _ := strconv.Atoi(ss[j])
			arr[i][j] = hoge
		}
	}
	return arr
}

func generateArray(i1, i2, i3 int) (array1, array2 [][]int) {
	// make i1 x i2 matrix
	array1 = make([][]int, i1)
	for i := 0; i < i1; i++ {
		array1[i] = make([]int, i2)
	}

	// make i2 x i3 matrix
	array2 = make([][]int, i2)
	for i := 0; i < i2; i++ {
		array2[i] = make([]int, i3)
	}
	return
}
