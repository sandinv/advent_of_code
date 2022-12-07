package rucksack

import (
	"bufio"
	"io"
)

const lowecase = 96
const uppercase = 64 - 26

func GetTotalBadgetPriorities(rdr io.Reader) (total int) {
	scanner := bufio.NewScanner(rdr)
	var i int
	var rucksacks [3]string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		rucksacks[i] = line
		i++
		if i == 3 {
			i = 0
			badget := FindBadget(rucksacks)
			total += GetPriority(badget)
		}
	}

	return
}

func FindBadget(rucksacks [3]string) rune {
	for i := 0; i < len(rucksacks[0]); i++ {
		for j := 0; j < len(rucksacks[1]); j++ {
			if rucksacks[0][i] == rucksacks[1][j] {
				for k := 0; k < len(rucksacks[2]); k++ {
					if rucksacks[1][j] == rucksacks[2][k] {
						return rune(rucksacks[2][k])
					}
				}

			}

		}

	}
	return rune(0)
}

func GetTotalPriorities(rdr io.Reader) (total int) {

	scanner := bufio.NewScanner(rdr)

	for scanner.Scan() {
		rucksacks := scanner.Text()
		if rucksacks == "" {
			continue
		}
		total += GetPriority(FindSharedItem(rucksacks, 1))
	}

	return
}
func FindSharedItem(rucksacks string, times int) rune {

	half := len(rucksacks) / 2
	first_rucksack := rucksacks[:half]
	second_rucksack := rucksacks[half:]
	total := 0

	for i := 0; i < len(first_rucksack); i++ {
		for j := 0; j < len(second_rucksack); j++ {
			if first_rucksack[i] == second_rucksack[j] {
				total += 1
				if total == times {
					return rune(first_rucksack[i])
				}
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
