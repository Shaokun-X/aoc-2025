package solver

import (
	"shaokun-x/aoc-2025/utils"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Solver8 struct {
	SolverSelector
}

func NewSolver8() *Solver8 {
	s := &Solver8{}
	s.SolverSelector.solver = s
	return s
}

func (s *Solver8) runSimple(data []string) string {
	coords := s.parse(data)
	dm := NewDistanceMap(coords)
	sort.Sort(dm)
	// fmt.Println(dm)

	// each circuit is a set of junction box indices
	circuits := []map[int]bool{}
	// number of connections
	count := 0
	for _, pair := range dm.Sorted {
		// fmt.Println(circuits)

		if count >= 1000 {
			break
		}
		// whether to start a new circuit
		new := true
		for i := 0; i < len(circuits); i++ {
			c := circuits[i]
			// one match, test if the another is already in another circuit
			if c[pair[0]] || c[pair[1]] {
				count += 1
				new = false
				// already in circuit
				if c[pair[0]] && c[pair[1]] {
					break
				}

				var other int
				if c[pair[0]] {
					other = pair[1]
				} else {
					other = pair[0]
				}
				merged := false
				for j := i + 1; j < len(circuits); j++ {
					otherC := circuits[j]
					if otherC[other] {
						// merge circuits
						for cdi := range otherC {
							c[cdi] = true
						}
						circuits = slices.Delete(circuits, j, j+1)
						merged = true
						break
					}
				}
				if !merged {
					c[other] = true
				}
			}
		}
		if new {
			count += 1
			newC := map[int]bool{pair[0]: true, pair[1]: true}
			circuits = append(circuits, newC)
		}
	}

	sort.Sort(ByMapLength(circuits))
	// fmt.Println(circuits)

	result := 1
	for i := range min(3, len(circuits)) {
		result *= len(circuits[i])
	}

	return strconv.Itoa(result)
}

func (s *Solver8) runNormal(data []string) string {
	coords := s.parse(data)
	dm := NewDistanceMap(coords)
	sort.Sort(dm)
	// fmt.Println(dm)

	// each circuit is a set of junction box indices
	circuits := []map[int]bool{}
	// indices of the last connected junction boxes
	var last [2]int
	// number of connections
	for _, pair := range dm.Sorted {
		// fmt.Println(circuits)

		if len(circuits) == 1 && len(circuits[0]) == len(coords) {
			break
		}
		// whether to start a new circuit
		new := true
		for i := 0; i < len(circuits); i++ {
			c := circuits[i]
			// one match, test if the another is already in another circuit
			if c[pair[0]] || c[pair[1]] {
				new = false
				last = pair
				// already in circuit
				if c[pair[0]] && c[pair[1]] {
					break
				}

				var other int
				if c[pair[0]] {
					other = pair[1]
				} else {
					other = pair[0]
				}
				merged := false
				for j := i + 1; j < len(circuits); j++ {
					otherC := circuits[j]
					if otherC[other] {
						// merge circuits
						for cdi := range otherC {
							c[cdi] = true
						}
						circuits = slices.Delete(circuits, j, j+1)
						merged = true
						break
					}
				}
				if !merged {
					c[other] = true
				}
			}
		}
		if new {
			newC := map[int]bool{pair[0]: true, pair[1]: true}
			circuits = append(circuits, newC)
		}
	}

	sort.Sort(ByMapLength(circuits))
	// fmt.Println(circuits)
	// fmt.Println(coords[last[0]], coords[last[1]])
	result := coords[last[0]].X * coords[last[1]].X

	return strconv.Itoa(result)
}

func (s *Solver8) parse(data []string) []Coordinates {
	result := []Coordinates{}
	for _, line := range data {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		result = append(result, Coordinates{x, y, z})
	}
	return result
}

type Coordinates struct {
	X, Y, Z int
}

type DistanceMap struct {
	// square of euclidean distance as fallback when mDistance is the same, mapped by indices in the original array
	EDistance map[[2]int]int
	// indices of pairs of coordinates sorted by distances
	Sorted      [][2]int
	coordinates []Coordinates
}

func NewDistanceMap(coords []Coordinates) *DistanceMap {
	n := len(coords)
	dm := &DistanceMap{}
	dm.EDistance = map[[2]int]int{}
	dm.Sorted = make([][2]int, n*(n-1)/2)

	si := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			mDist := SquareEuclideanDistance(coords[i], coords[j])
			dm.EDistance[[2]int{i, j}] = mDist
			dm.Sorted[si] = [2]int{i, j}
			si += 1
		}
	}
	dm.coordinates = coords
	return dm
}

func (dm DistanceMap) Len() int {
	return len(dm.Sorted)
}

func (dm DistanceMap) Swap(i, j int) {
	dm.Sorted[i], dm.Sorted[j] = dm.Sorted[j], dm.Sorted[i]
}

func (dm DistanceMap) Less(i, j int) bool {
	return dm.EDistance[dm.Sorted[i]] < dm.EDistance[dm.Sorted[j]]
}

func SquareEuclideanDistance(c1, c2 Coordinates) int {
	return utils.SquareEuclideanDistance(
		c1.X,
		c2.X,
		c1.Y,
		c2.Y,
		c1.Z,
		c2.Z,
	)
}

type ByMapLength []map[int]bool

func (a ByMapLength) Len() int           { return len(a) }
func (a ByMapLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMapLength) Less(i, j int) bool { return len(a[i]) > len(a[j]) }
