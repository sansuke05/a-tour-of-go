package main

import "fmt"

var c, python, java = true, false, "no!"

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	var i, j int = 1, 2
	k := 3
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(i, j, k, c, python, java)
}
