package camp

import (
	"strings"
	"testing"
)

func TestCountInclusiveParts(t *testing.T) {

	input := `
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`
	want := 2

	rdr := strings.NewReader(input[1:])
	got := CountInclusiveParts(rdr)
	if got != want {
		t.Errorf("got %d but wanted %d", got, want)
	}

}
