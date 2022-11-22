package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	counter := 0
	previous := float64(0)

	for {
		counter++
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-previous) < 1e-15 {
			break
		}
		previous = z
		fmt.Println(z)
	}
	fmt.Printf("Arrived at number after %d attempts\n", counter)
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
