package stacks

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type (
	Stack       []rune
	Instruction struct {
		count int
		src   int
		dst   int
	}
	Stacks []Stack
)

var NotEnoughCapacity = errors.New("Not enough capacity")
var InvalidScenario = errors.New("Instructions are not correct")

func PartI(rdr io.Reader) (string, error) {
	content := ReadScenario(rdr)
	stacks := CreateStacks(content)
	instructions := ReadInstructions(content)
	err := stacks.ApplyInstructions(instructions)
	if err != nil {
		return "", err
	}
	return stacks.GetTop(), nil

}

func ReadInstructions(content []string) []Instruction {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	result := []Instruction{}
	for _, line := range content {
		if !strings.Contains(line, "move") {
			continue
		}
		matches := re.FindStringSubmatch(line)
		count, _ := strconv.ParseInt(matches[1], 10, 64)
		src, _ := strconv.ParseInt(matches[2], 10, 64)
		dst, _ := strconv.ParseInt(matches[3], 10, 64)
		result = append(result, Instruction{
			count: int(count),
			src:   int(src),
			dst:   int(dst),
		})
	}
	return result

}

func (s *Stacks) ApplyInstructions(instructions []Instruction) error {
	for _, ins := range instructions {
		items, err := (*s)[ins.src-1].PopN(ins.count)
		if err != nil {
			return InvalidScenario
		}
		(*s)[ins.dst-1].Push(items...)
	}
	return nil

}

func (s *Stacks) GetTop() string {
	var build strings.Builder
	for _, stack := range *s {
		if len(stack) > 0 {
			build.WriteString(string(stack[len(stack)-1]))
		}
	}
	return build.String()

}

func CreateStacks(content []string) *Stacks {
	result := Stacks{}
	scenario := map[int][]rune{}
	num_line := 0
	num_column := 0
	for _, line := range content {
		if strings.Contains(line, "1") {
			break
		}
		num_column = 0
		for i := 0; i < len(line); i += 4 {
			item := line[i+1]
			if item != ' ' {
				scenario[num_column] = append(scenario[num_column], rune(item))
			}
			num_column++
		}
		num_line++
	}
	for i := 0; i < num_column; i++ {
		stack := Stack{}
		for j := len(scenario[i]) - 1; j >= 0; j-- {
			stack = append(stack, rune(scenario[i][j]))
		}
		result = append(result, stack)

	}
	return &result

}

func (s *Stack) Push(a ...rune) {

	for _, r := range a {
		*s = append(*s, r)
	}
}

func (s *Stack) PopN(count int) ([]rune, error) {

	result := []rune{}
	for i := 0; i < count; i++ {
		item, err := s.Pop()
		if err != nil {
			return nil, err
		}
		result = append(result, item)

	}
	return result, nil

}

func (s *Stack) Pop() (rune, error) {
	if len(*s) == 0 {
		return 0, NotEnoughCapacity
	}
	index := len(*s) - 1

	item := (*s)[index]
	*s = (*s)[:index]

	return item, nil
}

func ReadScenario(rdr io.Reader) (content []string) {
	scanner := bufio.NewScanner(rdr)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			content = append(content, line)
		}
	}
	return
}
