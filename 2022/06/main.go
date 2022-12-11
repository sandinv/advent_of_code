package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sandinv/advent_of_code/2022/06/markers"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		fmt.Printf("Error opening the file %s\n", err)
		os.Exit(-1)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Scan()

	fmt.Printf("Part I: %d", markers.FindMarker(scanner.Text()))

}
