package tournament

import (
	"strings"
	"testing"
)

func TestSimulateStrategy(t *testing.T) {

	t.Run("simple example", func(t *testing.T) {

		input := `
A Y
B X
C Z
`
		want := 15

		reader := strings.NewReader(input)
		got := SimulateStrategy(reader)
		if want != got {
			t.Errorf("wanted %d but got %d", want, got)
		}

	})

}
