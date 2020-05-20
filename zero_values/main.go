package main

import "fmt"

func main() {
	var x = 100
	var y = 5.2
	// go is a static typed language. you can not assgin different type of variable to each other. it will give compile time error.
	// x = y
	x = int(y)

	fmt.Println(x, y)
	//=====================
	var z int
	// z = "3" compile error
	z = 3
	fmt.Println(z)
	//========================
	// go initialise variable with zero values. there is no such un-initialise variable concept
	// this will help to reduce the bugs and errors.
	// zero values are like
	// int = 0
	// string ""
	// float 0
	//bool false
	//pointer nil
	var value int
	var name string
	var price float64
	var check bool

	fmt.Println(value, name, price, check)
}

/////////////////////////////////
// Types and Zero Values
// Go Playground: https://play.golang.org/p/zItROROXi64
/////////////////////////////////

// package main

// import "fmt"

// func main() {

//     // you must provide a type for each variable you declare or Go should be able to infer it
//     var a int = 10
//     var b = 15.5      // type inference (deduction)
//     c := "Gopher"     // short declaration, type inference
//     _, _, _ = a, b, c // Blank Identifier (_) to get rid of unused variable error

//     // Go is a Statically and Strong Typed Programming Language
//     // a = 3.14 -> error. A variable cannot change it's type
//     // a = b    -> error. It's not allowed to assign a type to another type

//     //** ZERO VALUES **//

//     // An uninitialized variable or empty variable  will get the so called ZERO VALUE
//     // The zero-value mechanism of Go ensures that a variable always holds a well defined value of its type
//     var value int                         // initialized with 0
//     var price float64                     // initialized with 0.0
//     var name string                       // initialized with empty string -> ""
//     var done bool                         // initialized with false
//     fmt.Println(value, price, name, done) // -> 0 0.0 ""  false
// }
