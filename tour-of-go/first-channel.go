package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)     // int型のchannelを作ってる？
	go sum(s[:len(s)/2], c) // 配列の前半をgoroutineで処理。結果はcに突っ込む
	go sum(s[len(s)/2:], c) // 配列の後半をgoroutineで処理。結果はcに突っ込む
	x, y := <-c, <-c        // receive from c

	fmt.Println(x, y, x+y)
}
