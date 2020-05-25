package main

import "fmt"

type km float64
type mile float64

func main() {
	type speed uint
	var s1 speed = 20
	var s2 speed = 10

	fmt.Println(s1 - s2)

	var x uint
	// x = s1 //compile time error because s1 is speed not uint
	x = uint(s1)
	_ = x

	var s3 speed = speed(x)
	_ = s3

	s3 = s3 / 2
	fmt.Println(s3)

	var puneToMumbai km = 150
	var distanceInMiles mile
	distanceInMiles = mile(puneToMumbai) / .623
	fmt.Println(distanceInMiles)
}
