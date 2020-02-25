package main

import (
	"bufio"
	"fmt"
	"os"
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

	dice_s := strings.Split(nextLine(), " ")
	dice := &Dice{dice_s[0], dice_s[1], dice_s[2], dice_s[3], dice_s[4], dice_s[5]}

	operations := nextLine()
	for _, op := range operations {
		switch op {
		case 'N':
			dice = dice.opN()
		case 'E':
			dice = dice.opE()
		case 'S':
			dice = dice.opS()
		case 'W':
			dice = dice.opW()
		default:
		}
	}

	fmt.Println(dice.one)
}
