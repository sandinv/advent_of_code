package main

import (
	"fmt"
	"os"

	"github.com/sandinv/advent_of_code/2022/02/tournament"
)

func main() {

	input, err := os.Open("input")
	defer input.Close()
	if err != nil {
		fmt.Printf("Error by opening the input %s\n", err)
		os.Exit(-1)
	}

	score := tournament.SimulateStrategy(input, true)
	fmt.Printf("The score is for part I is: %d\n", score)
	_, err = input.Seek(0, 0)
	if err != nil {
		fmt.Printf("Error while reading the input %s\n", err)
		os.Exit(-1)
	}
	score = tournament.SimulateStrategy(input, false)
	fmt.Printf("The score is for part II is: %d\n", score)
}
