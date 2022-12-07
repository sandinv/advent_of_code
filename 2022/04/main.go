package main

import (
	"fmt"
	"os"

	"github.com/sandinv/advent_of_code/2022/04/camp"
)

func main() {

	input, err := os.Open("input")
	if err != nil {
		fmt.Printf("Error opening the file %s\n", err)
		os.Exit(-1)
	}
	defer input.Close()

	fmt.Printf("Total of inclusive jobs: %d\n", camp.CountInclusiveParts(input))
	_, err = input.Seek(0, 0)
	if err != nil {
		fmt.Printf("Error reading the file %s\n", err)
		os.Exit(-1)
	}
	fmt.Printf("Total of overlap jobs: %d\n", camp.CountOverlapParts(input))
}
