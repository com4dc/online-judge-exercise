package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (n ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v",
		float64(n))
}

func Sqrt(x float64) (float64, error) {
	var e error
	if x < 0.0 {
		e = ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, e
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
