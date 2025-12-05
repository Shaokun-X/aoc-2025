package solver

import (
	"slices"
	"strconv"
	"strings"
)

type Solver5 struct {
	SolverSelector
}

func NewSolver5() *Solver5 {
	s := &Solver5{}
	s.SolverSelector.solver = s
	return s
}

func (s *Solver5) runSimple(data []string) string {
	ranges, ids := s.parse(data)
	merged := s.mergeRanges(ranges)
	result := 0
	for _, id := range ids {
		for _, rg := range merged {
			if id >= rg[0] && id <= rg[1] {
				result += 1
				break
			}
		}
	}
	// fmt.Println(merged, ids, result)
	return strconv.Itoa(result)
}

func (s *Solver5) runNormal(data []string) string {
	ranges, _ := s.parse(data)
	merged := s.mergeRanges(ranges)
	result := 0
	for _, rg := range merged {
		result += rg[1] - rg[0] + 1
	}
	return strconv.Itoa(result)
}

// Return a list of ranges and a list of ids
func (s *Solver5) parse(lines []string) ([][2]int, []int) {
	var ranges [][2]int
	var ids []int
	for _, line := range lines {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			left, _ := strconv.Atoi(parts[0])
			right, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, [2]int{left, right})
		} else {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		}
	}
	return ranges, ids
}

func (s *Solver5) mergeRanges(ranges [][2]int) [][2]int {
	var merged [][2]int
	for _, rg := range ranges {
		// nothing to compare with, insert
		if len(merged) == 0 {
			merged = append(merged, rg)
		} else {
			for i := 0; i < len(merged); i++ {
				if rg[0] >= merged[i][0] && rg[1] <= merged[i][1] {
					// merged[i] includes rg
					break
				} else if rg[1] < merged[i][0]-1 {
					// smaller (rg < merged[i]) insert at i
					merged = slices.Insert(merged, i, rg)
					break
				} else if rg[0] > merged[i][1]+1 {
					// larger skip
					if i == len(merged)-1 {
						// larger and at the end, append
						merged = append(merged, rg)
					} else {
						continue
					}
				} else {
					// merge
					merged[i][1] = max(rg[1], merged[i][1])
					merged[i][0] = min(rg[0], merged[i][0])
					var j int
					// mark delete till which index
					n := i + 1
					for j = i + 1; j < len(merged); j++ {
						if merged[i][1] >= merged[j][0]-1 {
							merged[i][1] = max(merged[i][1], merged[j][1])
							n += 1
						}
					}
					merged = slices.Delete(merged, i+1, n)
					break
				}
			}
		}
	}
	return merged
}
