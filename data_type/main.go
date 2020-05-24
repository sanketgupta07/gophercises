package main

import "fmt"

func main() {
	var i1 int8 = 100
	fmt.Printf("i1 type %T\n", i1)
	// var i2 int8 = 129
	// fmt.Printf("i2 type %T", i2) //.\main.go:8:16: constant 129 overflows int8

	var i3 uint16 = 65535
	fmt.Printf("i3 type %T\n", i3)

	var f1, f2, f3 = 2.2, -.5, 3.
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	// rune is alias for int32 and byte is for int8
	var r rune = 'a'
	fmt.Printf("r type %T\n and value of r is %v\n", r, r)

	// bool type
	var b bool = true
	fmt.Printf("b type %T\n", b)
	//string type
	var s string = "Hello go"
	fmt.Printf("s type %T\n", s)

	//array type -- fixed size like 4 in below example
	var numbers = [4]int{1, 2, 3, 4}
	fmt.Printf("numbers type %T\n", numbers)

	//slice type -- dynamic in size
	var num = []int{1, 2, 3, 4}
	fmt.Printf("num type %T\n", num)

	// map type
	balances := map[string]float64{
		"INR": 55.3,
		"USD": 88.1,
	}
	fmt.Printf("balances type %T\n", balances)

	// struct type

	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("you type %T\n", you)

	//Pointer type
	var value int = 53
	ptr := &value
	fmt.Printf("ptr type %T\n and value of ptr is %v\n", ptr, ptr)
	var ptr1 = &value
	fmt.Printf("ptr1 type %T\n and value of ptr1 is %v\n", ptr1, ptr1)

	//function type
	fmt.Printf("f type %T\n", f)
}

func f() {

}
