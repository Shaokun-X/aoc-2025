package solver

import (
	"shaokun-x/aoc-2025/utils"
	"strconv"
)

type Solver3 struct {
	SolverSelector
}

func NewSolver3() *Solver3 {
	s := &Solver3{}
	s.SolverSelector.solver = s
	return s
}

func (s *Solver3) runSimple(data []string) string {
	banks := s.parse(data)
	result := 0
	for _, bank := range banks {
		// fmt.Println(bank)
		tops := s.getTopN(bank, 2)
		joltage := s.buildJoltage(tops)
		result += joltage
		// println(joltage)
	}
	return strconv.Itoa(result)
}

func (s *Solver3) runNormal(data []string) string {
	banks := s.parse(data)
	result := 0
	for _, bank := range banks {
		// fmt.Println(bank)
		tops := s.getTopN(bank, 12)
		joltage := s.buildJoltage(tops)
		result += joltage
		// println(joltage)
	}
	return strconv.Itoa(result)
}

func (s *Solver3) parse(data []string) [][]int {
	var result [][]int
	for _, line := range data {
		var bats []int
		for _, n := range line {
			num, _ := strconv.Atoi(string(n))
			bats = append(bats, num)
		}
		result = append(result, bats)
	}
	return result
}

// Find top N that maintains the order in the original array.
func (s *Solver3) getTopN(nums []int, n int) []int {
	tops := make([]int, n)
	start, end := 0, len(nums)-n+1
	for range n {
		t, i := s.getTop(nums[start:end])
		tops = append([]int{t}, tops...)

		start += i + 1
		end += 1
	}
	return tops
}

// Find the largest number and its index
func (s *Solver3) getTop(nums []int) (int, int) {
	max := -1
	index := -1
	for i, n := range nums {
		if n > max {
			max = n
			index = i
		}
	}
	return max, index
}

func (s *Solver3) buildJoltage(tops []int) int {
	joltage := 0
	for i, t := range tops {
		joltage += t * utils.Pow(10, i)
	}
	return joltage
}
