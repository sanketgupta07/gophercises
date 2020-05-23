package main

import (
	"fmt"
)

func main() {
	const days int = 7 // use of constant will not be chekced like variables.
	const secondsInHour int = 3600

	var duration int = 55 //in hour
	fmt.Printf("The duration is %v\n", duration*secondsInHour)

	// Constants belogns to compile time while variables to runtime
	x, y := 5, 0
	_, _ = x, y
	// fmt.Println(x / y) // runtile error

	const a = 5
	const b = 0
	// fmt.Println(a / b) // compilation error

	//Multiple constants declation
	const x1, y1 int = 55, 22
	const n, m = 56, 66
	// Group contstants
	const (
		min1 = -600
		min2 = 200
		min3 = -300
	)
	fmt.Println(min1, min2, min3)

	// Group contstants
	// unassigned constants get the type and value from previous contant
	const (
		min4 = -200
		min5
		min6
	)
	fmt.Println(min4, min5, min6)

	// Constant Rules
	// 1.  Can not change constants
	const v int = 100
	// v = 50// compile time error

	// 2. Constant can not be initiate at runtime
	// const pow = math.Pow(2, 3) // Compile time error

	// 3. Constatnats can not be initiate using variables.
	var t = 3
	// const tc = t compile time error
	_ = t

	// 4. Constant can be assigned value using builtin functions like len("abc")

	const length = len("math") // len function belongs to compile time.

	fmt.Println("lenght const", length)

	// =========Untyped Constants=============
	const aa float64 = 5.3 // typeed constant
	const bb = 2.4         // Untyped constant

	// constant experssion
	const cc = aa * bb
	const str = "Hello" + "Go"
	fmt.Println(str, cc)

	// Go is strong typed language
	const dd int = 5
	// const ee float64 = dd * 5.0 // complile time error

	//  But untyped const can be used to break above scenario

	const xx = 5
	const yy = 2.3 * 5
	fmt.Printf("Type of xx %T, yy %T\n", xx, yy) // Untyped constant do not fllow the strict type rule.

	var ii int = xx     // xx change to int
	var jj float64 = xx // xx change to float64 behind the scene go change the type of xx like var jj float64 = float64(xx)
	var kk byte = xx

	fmt.Println(ii, jj, kk)

	// go will assign the type to untyped constant as per the used scenario

	const r = 5         // default type in for r as it is assigned to int literal 5
	var rr = r          // go will assign rr type to default r type (int)
	var rrr float64 = r // go will change the type to float64 for r and assign the value to rrr

	fmt.Printf("Type of rr: %T rrr %T\n", rr, rrr)
}
