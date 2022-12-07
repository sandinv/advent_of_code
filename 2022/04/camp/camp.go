package camp

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type assignment struct {
	init, end int
}

func fromString(asg string) assignment {
	rng := strings.Split(asg, "-")

	init, err := strconv.ParseInt(rng[0], 10, 32)
	if err != nil {
		return assignment{0, 0}
	}
	end, err := strconv.ParseInt(rng[1], 10, 32)
	if err != nil {
		return assignment{0, 0}
	}

	return assignment{
		init: int(init),
		end:  int(end),
	}

}
func (s assignment) in(another assignment) bool {
	return s.init >= another.init && s.end <= another.end
}
func (s assignment) overlap(another assignment) bool {
	return s.init >= another.init && s.init <= another.end
}

func CountInclusiveParts(rdr io.Reader) (total int) {

	scanner := bufio.NewScanner(rdr)

	for scanner.Scan() {
		line := scanner.Text()
		assignments := strings.Split(line, ",")
		first_assignment := fromString(assignments[0])
		second_assignment := fromString(assignments[1])
		if first_assignment.in(second_assignment) || second_assignment.in(first_assignment) {
			total++
		}
	}
	return
}

func CountOverlapParts(rdr io.Reader) (total int) {
	scanner := bufio.NewScanner(rdr)

	for scanner.Scan() {
		line := scanner.Text()
		assignments := strings.Split(line, ",")
		first_assignment := fromString(assignments[0])
		second_assignment := fromString(assignments[1])
		if first_assignment.overlap(second_assignment) || second_assignment.overlap(first_assignment) {
			total++
		}
	}

	return
}
