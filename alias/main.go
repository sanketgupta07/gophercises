package main

import "fmt"

func main() {
	var a uint8 = 10
	var b byte // byte is an alias for uint8 so we can assign value to each other without any converion
	b = a
	_ = b

	//Alias decalaration
	type second = uint
	var hour second = 3600
	fmt.Println("Mins in a hour ", hour/60)
}
