package calories

import (
	"strings"
	"testing"
)

func TestTopCalories(t *testing.T) {

	t.Run("simple example", func(t *testing.T) {

		example := `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
  `
		want := 24000

		rdr := strings.NewReader(example)
		got := TopCalories(rdr)

		if got[len(got)-1] != want {
			t.Errorf("want %d but got %d", want, got)
		}

	})

}
