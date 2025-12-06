package solver

import (
	"strconv"
)

type Solver1 struct {
	SolverSelector
}

func NewSolver1() *Solver1 {
	s := &Solver1{}
	s.SolverSelector.solver = s
	return s
}

func (s *Solver1) parse(data []string) []int {
	var steps []int

	for _, line := range data {
		minus := string(line[0]) == "L"
		var step int
		step, _ = strconv.Atoi(line[1:])
		if minus {
			step = -step
		}
		steps = append(steps, step)
	}
	return steps
}

func (s *Solver1) runSimple(data []string) string {
	pos := 50
	numbers := 100
	steps := s.parse(data)

	result := 0
	for _, step := range steps {
		pos += step
		pos = pos % numbers
		if pos < 0 {
			pos += numbers
		}
		if pos == 0 {
			result += 1
		}
	}

	return strconv.Itoa(result)
}

func (s *Solver1) runNormal(data []string) string {
	pos := 50
	numbers := 100
	steps := s.parse(data)

	result := 0
	for _, step := range steps {
		abs := step
		if abs < 0 {
			abs = -abs
		}
		rounds := abs / numbers
		remainder := step % numbers

		result += rounds

		// pos + remainder is always less than 2 rounds
		if pos != 0 && (pos+remainder <= 0 || pos+remainder > numbers-1) {
			result += 1
		}
		pos += remainder
		pos = pos % numbers
		if pos < 0 {
			pos += numbers
		}
	}

	return strconv.Itoa(result)
}
