package rucksack

import (
	"bufio"
	"io"
)

const lowecase = 96
const uppercase = 64 - 26

func GetTotalPriorities(rdr io.Reader) (total int) {

	scanner := bufio.NewScanner(rdr)

	for scanner.Scan() {
		rucksacks := scanner.Text()
		if rucksacks == "" {
			continue
		}
		total += GetPriority(FindSharedItem(rucksacks))
	}

	return
}
func FindSharedItem(rucksacks string) rune {

	half := len(rucksacks) / 2
	first_rucksack := rucksacks[:half]
	second_rucksack := rucksacks[half:]

	for i := 0; i < len(first_rucksack); i++ {
		for j := 0; j < len(second_rucksack); j++ {
			if first_rucksack[i] == second_rucksack[j] {
				return rune(first_rucksack[i])
			}
		}

	}
	return ' '
}

func GetPriority(item rune) int {

	if item >= 'a' && item <= 'z' {
		return int(item) - lowecase
	}
	return int(item) - uppercase
}
