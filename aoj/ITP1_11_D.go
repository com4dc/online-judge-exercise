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

	n, _ := strconv.Atoi(nextLine())

	base := strings.Split(nextLine(), " ")
	// create two dices
	base_dice := &Dice{base[0], base[1], base[2], base[3], base[4], base[5]}

	for i := 1; i < n; i++ {
		nums := strings.Split(nextLine(), " ")

		if reflect.DeepEqual(base, nums) {
			fmt.Println("No")
			return
		}

		dice := &Dice{nums[0], nums[1], nums[2], nums[3], nums[4], nums[5]}

		// dice2を回転させてdice1になるかチェック
		detectTop(base_dice, dice)
		detectFront(base_dice, dice)

		if base_dice.one == dice.one &&
			base_dice.two == dice.two &&
			base_dice.three == dice.three &&
			base_dice.four == dice.four &&
			base_dice.five == dice.five &&
			base_dice.six == dice.six {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
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
