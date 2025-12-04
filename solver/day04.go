package solver

import (
	"strconv"
	"strings"
)

type Solver4 struct {
	SolverSelector
}

func NewSolver4() *Solver4 {
	s := &Solver4{}
	s.SolverSelector.solver = s
	return s
}

func (s *Solver4) runSimple(data []string) string {
	grid := s.parse(data)
	width := len(grid[0])
	height := len(grid)

	count := 0
	// i is row, j is column
	for i := range height {
		for j := range width {
			if s.isPaper(grid[i][j]) && s.isAccessible(grid, i, j) {
				count += 1
			}
		}
	}
	return strconv.Itoa(count)
}

func (s *Solver4) runNormal(data []string) string {
	grid := s.parse(data)
	width := len(grid[0])
	height := len(grid)

	result := 0

	for {
		count := 0
		// i is row, j is column
		for i := range height {
			for j := range width {
				if s.isPaper(grid[i][j]) && s.isAccessible(grid, i, j) {
					count += 1
					grid[i][j] = "."
				}
			}
		}
		if count == 0 {
			break
		}
		result += count
	}
	return strconv.Itoa(result)
}

func (s *Solver4) parse(data []string) [][]string {
	var result [][]string
	for _, line := range data {
		result = append(result, strings.Split(line, ""))
	}
	return result
}

func (s *Solver4) isPaper(char string) bool {
	return char == "@"
}

// i is row, j is column
func (s *Solver4) isAccessible(grid [][]string, i, j int) bool {
	width := len(grid[0])
	height := len(grid)

	// limit to edges
	left := max(0, j-1)
	right := min(width-1, j+1)
	up := max(0, i-1)
	down := min(height-1, i+1)

	adjPapers := 0
	for m := up; m <= down; m++ {
		for n := left; n <= right; n++ {
			if m == i && n == j {
				continue
			}
			if s.isPaper(grid[m][n]) {
				adjPapers += 1
			}
		}
	}

	return adjPapers < 4
}
