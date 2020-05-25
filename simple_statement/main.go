package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("20")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if j, err := strconv.Atoi("500"); err == nil {
		fmt.Println(j)
	} else {
		fmt.Println(err)
	}

	//Convert cmdline input int to float
	if args := os.Args; len(args) != 2 {
		fmt.Println("pls provide an int input")
	} else if i, err = strconv.Atoi(args[1]); err != nil {
		fmt.Println("pls enter a valid int input")
	} else {
		fmt.Println("input:", i)
		fmt.Printf("converted float %f, %T", float64(i), float64(i))
	}
}
