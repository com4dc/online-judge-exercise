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

func (d *Dice) find(num string) int {
	switch num {
	case d.one:
		return 1
	case d.two:
		return 2
	case d.three:
		return 3
	case d.four:
		return 4
	case d.five:
		return 5
	case d.six:
		return 6
	default:
		return -1
	}
}

func main() {
	// Create Dice Constructor
	dice_s := strings.Split(nextLine(), " ")

	q_num, _ := strconv.Atoi(nextLine())

	for i := 0; i < q_num; i++ {
		dice := &Dice{dice_s[0], dice_s[1], dice_s[2], dice_s[3], dice_s[4], dice_s[5]}
		q_s := strings.Split(nextLine(), " ")

		switch q_s[0] {
		case dice.two:
			dice.opN()
		case dice.three:
			dice.opW()
		case dice.four:
			dice.opE()
		case dice.five:
			dice.opS()
		case dice.six:
			{
				dice.opN()
				dice.opN()
			}
		}

		switch q_s[1] {
		case dice.one:
			dice.opS()
		case dice.three:
			{
				dice.opN()
				dice.opW()
				dice.opS()
			}
		case dice.four:
			{
				dice.opN()
				dice.opE()
				dice.opS()
			}

		case dice.five:
			{
				dice.opS()
				dice.opS()
				dice.opE()
				dice.opE()
			}
		case dice.six:
			dice.opN()
		}

		fmt.Println(dice.three)
	}
}
