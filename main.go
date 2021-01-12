package main

import (
	"fmt"
	"math"
)

// Basics
var c, python, java = true, false, "no!"

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func convertTypes() {
	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)
}

func inferType() {
	v := 42.1
	g := 0.867 + 0.5i
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("g is of type %T\n", g)
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func numericConstants() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func MySqrt(x float64) float64 {
	z := 1.0
	for {
		z = x - (x*x-2)/(2*x)
		fmt.Println(z)
		if math.Abs(z-x) < 0.0001 {
			break
		}
		x = z
	}
	return z
}

func main() {
	var i, j int = 1, 2
	k := 3

	a, b := swap("hello", "world")

	fmt.Println(a, b)
	fmt.Println(i, j, k, c, python, java)

	convertTypes()
	inferType()

	const Truth = true
	fmt.Println("Go rules?", Truth)

	numericConstants()
	fmt.Println(MySqrt(2))
}
