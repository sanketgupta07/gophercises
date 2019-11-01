package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type problem struct {
	statement string
	answer    string
}

var problems []problem
var score int

func loadQuestions() {
	fmt.Println("Hang on.. loading questions")
	lines, err := os.Open("problems.csv")

	if err != nil {
		log.Fatalln("Could not open problems.csv file", err)
	}

	line := csv.NewReader(lines)
	for {
		record, err := line.Read()
		if err == io.EOF {
			break
		}

		// make([]problem, len(problems)+1)
		p := problem{statement: record[0], answer: record[1]}
		problems = append(problems, p)
	}
}
func main() {
	// reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to quiz game.")
	opt := os.Args[1]

	switch opt {
	case "-h":
		fmt.Println("Answer the small quiz questions as soon as you see them.")
	default:
		fmt.Print("Starting Quiz...Hoooh..Press 's' key to start:")
		var response int
		fmt.Scanf("%c", &response)
		switch response {
		case 's' | 'S':
			go loadQuestions()
			time.Sleep(5 * time.Second)
			fmt.Println("Starting")
		default:
			fmt.Println("Quiting")
			return
		}

		fmt.Println("Your Score:", score, "/", len(problems))
	}
}
