package main

import (
	"fmt"
	"os"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to quiz game.")
	opt := os.Args[1]

	switch opt {
	case "-h":
		fmt.Println("Answer the small quiz questions as soon as you see them.")
	default:
		fmt.Println("Starting Quiz...Hoooh..Press 's' key to start. ")
		var response int
		fmt.Scanf("%c", &response)
		switch response {
		case 's' | 'S':
			fmt.Println("Starting")
		default:
			fmt.Println("Quiting")
			return
		}

		fmt.Println("Quiting")
	}
}
