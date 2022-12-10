package main

import (
	"fmt"
	"os"

	"github.com/sandinv/advent_of_code/2022/05/stacks"
)

func main() {

	input, err := os.Open("input")
	if err != nil {
		fmt.Printf("Error opening the file %s\n", err)
		os.Exit(-1)
	}
	defer input.Close()

	top, err := stacks.PartI(input)
	if err != nil {
		fmt.Printf("Error processing the scenario %s\n", err)
	}
	fmt.Printf("Top of the stacks: %q\n", top)

}
