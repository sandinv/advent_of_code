package stacks

import (
	"reflect"
	"strings"
	"testing"
)

func TestPartI(t *testing.T) {
	input := `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

	got, err := PartI(strings.NewReader(input))
	assertNoError(t, err)
	want := "CMZ"
	if got != want {
		t.Errorf("wanted %q but got %q", want, got)
	}

}

func TestCreateStacks(t *testing.T) {

	input := `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3
`[1:]
	content := ReadScenario(strings.NewReader(input))
	got := CreateStacks(content)

	want := &Stacks{
		{'Z', 'N'},
		{'M', 'C', 'D'},
		{'P'},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %+v but got %+v", want, got)
	}

}

func TestReadInstructions(t *testing.T) {
	input := `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`[1:]

	content := ReadScenario(strings.NewReader(input))
	got := ReadInstructions(content)
	want := []Instruction{
		{count: 1, src: 2, dst: 1},
		{count: 3, src: 1, dst: 3},
		{count: 2, src: 2, dst: 1},
		{count: 1, src: 1, dst: 2},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted %+v but got %+v", want, got)
	}

}

func TestCrate(t *testing.T) {

	t.Run("Push an item into the Stack", func(t *testing.T) {

		stack := &Stack{}
		stack.Push('A')
		want := &Stack{'A'}

		if !reflect.DeepEqual(stack, want) {
			t.Errorf("wanted %+v but got %+v", *want, *stack)
		}
	})

	t.Run("Push multiple items into the Stack", func(t *testing.T) {
		stack := &Stack{}
		stack.Push('A', 'B')
		want := &Stack{'A', 'B'}
		if !reflect.DeepEqual(stack, want) {
			t.Errorf("wanted %+v but got %+v", *want, *stack)
		}
	})

	t.Run("Pop one item from the Stack", func(t *testing.T) {
		stack := &Stack{'A'}
		want := rune('A')
		got, err := stack.Pop()

		assertNoError(t, err)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %+v but got %+v", want, got)
		}

	})

	t.Run("Pop two items from the Stack", func(t *testing.T) {
		stack := &Stack{'A', 'B'}
		want := []rune{'B', 'A'}
		got, err := stack.PopN(2)

		assertNoError(t, err)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %+v but got %+v", want, got)
		}

	})

	t.Run("Pop without enough capacity", func(t *testing.T) {
		stack := &Stack{}
		_, err := stack.Pop()
		assertError(t, err)
	})

}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Did not expected an error %s", err)
	}
}
func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("Did expected an error")
	}
}
