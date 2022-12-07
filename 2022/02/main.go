package main

import (
	"fmt"
	"os"

	"github.com/sandinv/advent_of_code/2022/02/tournament"
)

func main() {

	input, err := os.Open("input")
	if err != nil {
		fmt.Printf("Error by opening the input %s\n", err)
		os.Exit(-1)
	}

	score := tournament.SimulateStrategy(input)
	fmt.Printf("The score is: %d\n", score)
}
