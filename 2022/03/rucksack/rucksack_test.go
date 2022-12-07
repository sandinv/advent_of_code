package rucksack

import (
	"fmt"
	"strings"
	"testing"
)

func TestProcessRucksacks(t *testing.T) {
	rucksacks := `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`
	want := 157
	reader := strings.NewReader(rucksacks)

	got := GetTotalPriorities(reader)

	if got != want {
		t.Errorf("wanted %d but got %d", want, got)
	}

}

func TestFindSharedItem(t *testing.T) {

	cases := []struct {
		input string
		want  rune
	}{
		{input: "vJrwpWtwJgWrhcsFMMfFFhFp", want: 'p'},
		{input: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", want: 'L'},
		{input: "PmmdzqPrVvPwwTWBwg", want: 'P'},
		{input: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", want: 'v'},
		{input: "ttgJtRGJQctTZtZT", want: 't'},
		{input: "CrZsJsPPZsGzwwsLwLmpwMDw", want: 's'},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("Rucksack %s", test.input), func(t *testing.T) {
			got := FindSharedItem(test.input)
			if got != test.want {
				t.Errorf("wanted %v but got %v", test.want, got)
			}
		})
	}
}

func TestGetPriority(t *testing.T) {
	cases := []struct {
		name  string
		input rune
		want  int
	}{
		{name: "lowercase", input: 'a', want: 1},
		{name: "last lowercase", input: 'z', want: 26},
		{name: "uppercase", input: 'A', want: 27},
		{name: "last uppercase", input: 'Z', want: 52},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			got := GetPriority(test.input)
			if got != test.want {
				t.Errorf("wanted %d but got %d", test.want, got)
			}
		})
	}
}
