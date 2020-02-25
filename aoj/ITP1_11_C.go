package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// Scan Standard Input
var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

type Dice struct {
	one   string
	two   string
	three string
	four  string
	five  string
	six   string
}

func (d *Dice) opN() *Dice {
	tmp := d.one // copy

	d.one = d.two
	d.two = d.six
	d.six = d.five
	d.five = tmp

	return d
}

func (d *Dice) opE() *Dice {
	tmp := d.one // copy

	d.one = d.four
	d.four = d.six
	d.six = d.three
	d.three = tmp

	return d
}

func (d *Dice) opS() *Dice {
	tmp := d.one // copy

	d.one = d.five
	d.five = d.six
	d.six = d.two
	d.two = tmp
	return d
}

func (d *Dice) opW() *Dice {
	tmp := d.one // copy

	d.one = d.three
	d.three = d.six
	d.six = d.four
	d.four = tmp
	return d
}

func main() {

	in1 := strings.Split(nextLine(), " ")
	in2 := strings.Split(nextLine(), " ")

	if reflect.DeepEqual(in1, in2) {
		fmt.Println("Yes")
		return
	}

	// create two dices
	dice1 := &Dice{in1[0], in1[1], in1[2], in1[3], in1[4], in1[5]}
	dice2 := &Dice{in2[0], in2[1], in2[2], in2[3], in2[4], in2[5]}

	// dice2を回転させてdice1になるかチェック
	detectTop(dice1, dice2)
	detectFront(dice1, dice2)

	if dice1.three == dice2.three && dice1.four == dice2.four {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func detectTop(dice1, dice2 *Dice) {
	switch dice1.one {
	case dice2.two:
		dice2.opN()
	case dice2.three:
		dice2.opW()
	case dice2.four:
		dice2.opE()
	case dice2.five:
		dice2.opS()
	case dice2.six:
		{
			dice2.opN()
			dice2.opN()
		}
	}
	return
}

func detectFront(dice1, dice2 *Dice) {
	switch dice1.two {
	case dice2.one:
		dice2.opS()
	case dice2.three:
		{
			dice2.opN()
			dice2.opW()
			dice2.opS()
		}
	case dice2.four:
		{
			dice2.opN()
			dice2.opE()
			dice2.opS()
		}

	case dice2.five:
		{
			dice2.opS()
			dice2.opS()
			dice2.opE()
			dice2.opE()
		}
	case dice2.six:
		dice2.opN()
	}
	return
}
