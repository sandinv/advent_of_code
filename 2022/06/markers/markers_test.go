package markers

import (
	"fmt"
	"testing"
)

func TestFindMarker(t *testing.T) {

	cases := []struct {
		input string
		want  int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, test := range cases {

		t.Run(fmt.Sprintf("Case_%q", test.input), func(t *testing.T) {
			got := FindMarker(test.input)

			if got != test.want {
				t.Errorf("expected %d but got %d", test.want, got)
			}

		})

	}
}
