package solver

import (
	"fmt"
	"shaokun-x/aoc-2025/utils"
	"strconv"
	"strings"
)

type Solver6 struct {
	SolverSelector
}

func NewSolver6() *Solver6 {
	s := &Solver6{}
	s.SolverSelector.solver = s
	return s
}

type Expression struct {
	Numbers  []int
	Operator string
}

func (exp *Expression) add() int {
	result := 0
	for _, n := range exp.Numbers {
		result += n
	}
	return result
}

func (exp *Expression) multiply() int {
	result := 1
	for _, n := range exp.Numbers {
		result *= n
	}
	return result
}

func (exp *Expression) Execute() int {
	switch exp.Operator {
	case "+":
		return exp.add()
	case "*":
		return exp.multiply()
	}
	fmt.Println("Error, unknown operator")
	return 0
}

func (s *Solver6) parse(data []string) []Expression {
	var exps []Expression

	n := len(strings.Fields(data[0]))
	for range n {
		exps = append(exps, Expression{make([]int, len(data)-1), ""})
	}

	for i, line := range data {
		parts := strings.Fields(line)
		for j, p := range parts {
			if i < len(data)-1 {
				num, _ := strconv.Atoi(p)
				exps[j].Numbers[i] = num
			} else {
				exps[j].Operator = p
			}
		}
	}
	return exps
}

func (s *Solver6) parseVertical(data []string) []Expression {
	width := len(data[0])
	height := len(data)

	n := len(strings.Fields(data[height-1]))
	exps := make([]Expression, n)
	for i := range n {
		exps[i] = Expression{make([]int, 0), ""}
	}

	// expression index
	expi := 0
	for i := width - 1; i >= 0; i-- {
		var digits []int
		for j := range height - 1 {
			if string(data[j][i]) == " " {
				continue
			}
			num, _ := strconv.Atoi(string(data[j][i]))
			digits = append(digits, num)
		}

		num := 0
		for j, d := range digits {
			num += d * utils.Pow(10, len(digits)-1-j)
		}

		if num > 0 {
			exps[expi].Numbers = append(exps[expi].Numbers, num)
		}

		if string(data[height-1][i]) != " " {
			exps[expi].Operator = string(data[height-1][i])
			expi += 1
		}
	}
	return exps
}

func (s *Solver6) runSimple(data []string) string {
	exps := s.parse(data)
	result := 0
	for _, exp := range exps {
		result += exp.Execute()
	}
	return strconv.Itoa(result)
}

func (s *Solver6) runNormal(data []string) string {
	exps := s.parseVertical(data)
	result := 0
	for _, exp := range exps {
		result += exp.Execute()
	}
	return strconv.Itoa(result)
}
