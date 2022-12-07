package calories

import (
	"bufio"
	"io"
	"sort"
	"strconv"
)

func TopCalories(reader io.Reader) (most_calories []int) {

	scanner := bufio.NewScanner(reader)
	total_calories := 0
	max_calories := 0

	for scanner.Scan() {
		calories, err := strconv.Atoi(scanner.Text())

		if err != nil {
			if max_calories < total_calories {
				max_calories = total_calories
			}
			most_calories = append(most_calories, total_calories)
			total_calories = 0
		} else {
			total_calories += calories
		}

	}
	sort.Ints(most_calories)
	return
}
