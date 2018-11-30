package main

import "fmt"

func main() {
	a := [5]int{25, 50, 75, 100, 125}
	fmt.Println("emp:", a)
	fmt.Println("set:", a)
	fmt.Println("get:", a[4], a[2])

	var twoD [2][3]int
	for b := 0; b < 3; b++ {
		for c := 0; c < 3; c++ {
			twoD[b][c] = b + c

		}
	}
	fmt.Println("2d:", twoD)

}
