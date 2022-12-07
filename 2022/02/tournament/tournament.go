package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

var UndefinedHand = errors.New("Unknown hand")
var UndefinedResult = errors.New("Unknown result")

const (
	rock    = 1
	paper   = 2
	scissor = 3
	win     = 'Z'
	draw    = 'Y'
	lose    = 'X'
)

type hand int

type result uint

func fromHand(x rune) (hand, error) {
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

func SimulateStrategy(reader io.Reader, guess bool) (points int) {

	scanner := bufio.NewScanner(reader)
	var player hand
	for scanner.Scan() {
		round := scanner.Text()
		if round == "" || len(round) != 3 {
			continue
		}

		elf, err := fromHand(rune(round[0]))
		if err != nil {
			fmt.Println(elf)
			return 0
		}

		if guess {
			player, err = fromHand(rune(round[2]))
			if err != nil {
				return 0
			}
		} else {
			result, err := fromResult(rune(round[2]))
			if err != nil {
				return 0
			}
			player = guessHand(elf, result)
		}
		points += player.battle(elf)
	}
	return
}

func fromResult(state rune) (result, error) {
	switch state {
	case 'X':
		return result(lose), nil
	case 'Y':
		return result(draw), nil
	case 'Z':
		return result(win), nil
	default:
		fmt.Println("Unknown")
		return result(0), UndefinedResult
	}
}

func guessHand(elf hand, r result) hand {
	switch r {
	case result(win):
		if int(elf)+1 > 3 {
			return hand(1)
		} else {
			return hand(elf + 1)
		}
	case result(lose):
		if int(elf)-1 == 0 {
			return hand(3)
		} else {
			return hand(elf - 1)
		}
	case result(draw):
		return elf
	}
	return hand(0)
}
