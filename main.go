package main

import "fmt"

func main() {
	amount := 999_999_00.00
	minper := 0.14
	maxper := 0.16
	mininc := amount * minper
	maxinc := amount * maxper
	fmt.Println(mininc)
	fmt.Println(maxinc)

}
