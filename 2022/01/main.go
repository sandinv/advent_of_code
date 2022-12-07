package main

import (
	"fmt"
	"os"

	"github.com/sandinv/advent_of_code/2022/01/calories"
)

func main() {

	input_file, err := os.Open("input")

	if err != nil {
		fmt.Printf("Cannot open input file %s\n", err)
		os.Exit(-1)
	}

	cal := calories.TopCalories(input_file)

	fmt.Printf("Most calories carried: %v\n", cal[len(cal)-1])
	fmt.Printf("Total of calories from the top 3 elves %d\n", SumTop3(cal))

}

func SumTop3(calories []int) (total int) {
	for i := 0; i < 3; i++ {
		index := len(calories) - 1
		total += calories[index-i]
	}
	return
}
