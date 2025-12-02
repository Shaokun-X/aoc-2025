package solver

import (
	"strconv"
	"strings"
)

type Solver2 struct {
	SolverSelector
}

func NewSolver2() *Solver2 {
	s := &Solver2{}
	s.SolverSelector.solver = s
	return s
}

func (s *Solver2) runSimple(data []string) string {
	result := 0
	ranges := s.parse(data)
	for _, r := range ranges {
		left, right := r[0], r[1]
		for _, seg := range s.divide(left, right) {
			dig := s.getDigits(seg[0])
			if dig%2 != 0 {
				continue
			}
			pow := dig / 2
			// valid split
			start := s.getNextInvalidRoot(seg[0], pow)
			end := s.getPreviousInvalidRoot(seg[1], pow)

			for j := start; j <= end; j++ {
				result += s.buildInvalid(j, 2)
			}
		}
	}
	return strconv.Itoa(result)
}

func (s *Solver2) runNormal(data []string) string {
	result := 0
	ranges := s.parse(data)
	for _, r := range ranges {
		left, right := r[0], r[1]
		// println(left, right)
		for _, seg := range s.divide(left, right) {
			// dig is guaranteed to be greater than 1 by s.divide
			dig := s.getDigits(seg[0])
			invalid := make(map[int]bool)
			for i := dig - 1; i > 0; i-- {
				if dig%i == 0 {
					// println(seg[0], seg[1], i)
					// valid split
					start := s.getNextInvalidRoot(seg[0], i)
					end := s.getPreviousInvalidRoot(seg[1], i)
					// println(start, end)
					for j := start; j <= end; j++ {
						invalid[s.buildInvalid(j, dig/i)] = true
					}
				}
			}
			for iv := range invalid {
				// println("iv", iv)
				result += iv
			}
		}
	}

	return strconv.Itoa(result)
}

func (s *Solver2) parse(data []string) [][2]int {
	var result [][2]int
	pairs := strings.Split(data[0], ",")
	for _, range_ := range pairs {
		parts := strings.Split(range_, "-")
		left, _ := strconv.Atoi(parts[0])
		right, _ := strconv.Atoi(parts[1])
		result = append(result, [2]int{left, right})
	}
	return result
}

// Find the root of the closest invalid number bigger than given number. If the given number is already an invalid, return its root.
func (s *Solver2) getNextInvalidRoot(n int, pow int) int {
	nums := s.decompose(n, pow)

	for _, num := range nums {
		if num == nums[0] {
			continue
		}
		if num > nums[0] {
			// if nums[0] is 99 this will not run, so no risk of increasing digits
			return nums[0] + 1
		}
		return nums[0]
	}

	return nums[0]
}

// Find the root of the closest invalid number smaller than given number. If the given number is already an invalid, return its root.
//
// When the given number is a power of 10, return 0.
func (s *Solver2) getPreviousInvalidRoot(n int, pow int) int {
	nums := s.decompose(n, pow)

	for _, num := range nums {
		if num == nums[0] {
			continue
		}
		if num < nums[0] {
			// special case, nums[0] is power of 10
			if nums[0] == s.pow(10, pow-1) {
				return 0
			}
			return nums[0] - 1
		}
		return nums[0]
	}
	return nums[0]
}

func (s *Solver2) buildInvalid(root, repeat int) int {
	digs := s.getDigits(root)
	result := 0
	for i := 0; i < repeat; i++ {
		result += root * s.pow(10, i*digs)
	}
	return result
}

func (s *Solver2) getDigits(n int) int {
	var digs int
	for digs = 0; n > 0; n /= 10 {
		digs++
	}
	return digs
}

// Break a number by digits, pow is the number of digits.
func (s *Solver2) decompose(n int, pow int) []int {
	var result []int
	base := 1
	for i := 0; i < pow; i++ {
		base *= 10
	}
	m := n
	for m > 0 {
		result = append([]int{m % base}, result...)
		m /= base
	}
	return result
}

// Break given range to segments so that numbers in each range segment have same number of digits.
//
// 1 digit range is ignored.
func (s *Solver2) divide(lower, upper int) [][2]int {
	ldig := s.getDigits(lower)
	udig := s.getDigits(upper)

	// at least 2 digits is meaningful in this problem
	ldig = max(2, ldig)

	var result [][2]int
	for i := ldig; i <= udig; i++ {
		start := max(lower, s.pow(10, i-1))
		end := min(upper, s.pow(10, i)-1)
		result = append(result, [2]int{start, end})
	}

	return result
}

func (s *Solver2) pow(n, p int) int {
	if p == 0 {
		return 1
	}
	base := n
	for i := 0; i < p-1; i++ {
		n *= base
	}
	return n
}
