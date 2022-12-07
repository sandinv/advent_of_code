package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

var UndefinedHand = errors.New("Unknown hand")

const (
	rock    = 1
	paper   = 2
	scissor = 3
)

type hand int

func fromRune(x rune) (hand, error) {
	switch x {
	case 'A', 'X':
		return hand(rock), nil
	case 'B', 'Y':
		return hand(paper), nil
	case 'C', 'Z':
		return hand(scissor), nil
	default:
		fmt.Println("Unknown")
		return hand(0), UndefinedHand
	}
}

func (h hand) battle(e hand) int {
	var result int
	if h == e {
		result = 3
	}
	switch int(h) - int(e) {
	case 1, -2:
		// you win
		result = 6
	case -1, 2:
		// you lose
		result = 0
	}
	return result + int(h)
}

func SimulateStrategy(reader io.Reader) (points int) {

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		round := scanner.Text()
		if round == "" || len(round) != 3 {
			continue
		}
		elf, err := fromRune(rune(round[0]))
		if err != nil {
			fmt.Println(elf)
			return 0
		}
		player, err := fromRune(rune(round[2]))
		if err != nil {
			return 0
		}
		points += player.battle(elf)
	}
	return
}
