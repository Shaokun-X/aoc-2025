package solver

import (
	"shaokun-x/aoc-2025/utils"
	"sort"
	"strconv"
	"strings"
)

type Solver9 struct {
	SolverSelector
}

func NewSolver9() *Solver9 {
	s := &Solver9{}
	s.SolverSelector.solver = s
	return s
}

func (s *Solver9) runSimple(data []string) string {
	coords := s.parse(data)
	n := len(coords)

	corners := Square{}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if getSurface(coords[i], coords[j]) > getSurface(corners[0], corners[1]) {
				corners = Square{coords[i], coords[j]}
			}
		}
	}

	result := getSurface(corners[0], corners[1])

	return strconv.Itoa(result)
}

func (s *Solver9) runNormal(data []string) string {
	coords := s.parse(data)
	n := len(coords)

	pairs := []Square{}
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			pairs = append(pairs, Square{coords[i], coords[j]})
		}
	}

	sort.Sort(BySurface(pairs))

	vectors := make([]Vector, n)
	for i := range n {
		j := i + 1
		if i == n-1 {
			j = 0
		}
		vectors[i] = Vector{coords[i], coords[j]}
	}

	// fmt.Println(pairs)
	// fmt.Println(vectors)

	corners := Square{}

	for _, sq := range pairs {
		inbound := true
		for _, v := range vectors {
			if sq.isRelevant(v) {
				// not sure if it is clock wise,
				// but at most try 2 times
				if v.isOutside(sq, true) {
					inbound = false
					break
				}
			}
		}
		if inbound {
			corners = sq
			break
		}
	}

	// fmt.Println(corners)

	result := getSurface(corners[0], corners[1])

	return strconv.Itoa(result)
}

type Tile [2]int

type Square [2]Tile

// check if a vector has intersection with a square
func (s Square) isRelevant(v Vector) bool {
	minX := min(s[0][0], s[1][0])
	maxX := max(s[0][0], s[1][0])
	minY := min(s[0][1], s[1][1])
	maxY := max(s[0][1], s[1][1])

	if v.isVertical() {
		// no intersection
		if v[0][0] < minX || v[0][0] > maxX {
			return false
		}
		maxVY := max(v[0][1], v[1][1])
		minVY := min(v[0][1], v[1][1])

		return overlaps([2]int{minVY, maxVY}, [2]int{minY, maxY})

	} else {
		// no intersection
		if v[0][1] < minY || v[0][1] > maxY {
			return false
		}
		maxVX := max(v[0][0], v[1][0])
		minVX := min(v[0][0], v[1][0])

		return overlaps([2]int{minVX, maxVX}, [2]int{minX, maxX})
	}
}

type BySurface []Square

func (a BySurface) Len() int      { return len(a) }
func (a BySurface) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BySurface) Less(i, j int) bool {
	return getSurface(a[i][0], a[i][1]) > getSurface(a[j][0], a[j][1])
}

// From Vector[0] to Vector[1]
type Vector [2]Tile

func (v *Vector) isVertical() bool {
	start, end := (*v)[0], (*v)[1]
	vertical := false
	if start[0] == end[0] {
		vertical = true
	}
	return vertical
}

func (v *Vector) isOutside(sq Square, clockwise bool) bool {
	start, end := (*v)[0], (*v)[1]
	vertical := v.isVertical()

	// increasing direciton, left to right or up to down
	dir := true
	if vertical {
		dir = end[1] > start[1]
	} else {
		dir = end[0] > start[0]
	}

	// check if overlaps
	if vertical {
		if max(sq[0][1], sq[1][1]) <= min(start[1], end[1]) || min(sq[0][1], sq[1][1]) >= max(start[1], end[1]) {
			return false
		}
	} else {
		if max(sq[0][0], sq[1][0]) <= min(start[0], end[0]) || min(sq[0][0], sq[1][0]) >= max(start[0], end[0]) {
			return false
		}
	}

	if clockwise {
		if vertical && dir {
			return sq[0][0] > start[0] || sq[1][0] > start[0]
		}
		if vertical && !dir {
			return sq[0][0] < start[0] || sq[1][0] < start[0]
		}
		if !vertical && dir {
			return sq[0][1] < start[1] || sq[1][1] < start[1]
		}
		if !vertical && !dir {
			return sq[0][1] > start[1] || sq[1][1] > start[1]
		}
	} else {
		if vertical && dir {
			return sq[0][0] < start[0] || sq[1][0] < start[0]
		}
		if vertical && !dir {
			return sq[0][0] > start[0] || sq[1][0] > start[0]
		}
		if !vertical && dir {
			return sq[0][1] > start[1] || sq[1][1] > start[1]
		}
		if !vertical && !dir {
			return sq[0][1] < start[1] || sq[1][1] < start[1]
		}
	}

	return false
}

func (s *Solver9) parse(data []string) []Tile {
	result := []Tile{}
	for _, line := range data {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		result = append(result, [2]int{x, y})
	}
	return result
}

func getSurface(c1, c2 Tile) int {
	return (utils.Abs(c1[0]-c2[0]) + 1) * (utils.Abs(c1[1]-c2[1]) + 1)
}

// Compare if 2 open ranges overlap. Note that the range array must be in ascending order.
func overlaps(r1, r2 [2]int) bool {
	if r1[0] == r2[0] {
		// empty range is considered overlapping in this case
		return true
	}
	var smaller, larger [2]int
	if r1[0] < r2[0] {
		smaller = r1
		larger = r2
	} else {
		smaller = r2
		larger = r1
	}
	return larger[0] < smaller[1]
}
