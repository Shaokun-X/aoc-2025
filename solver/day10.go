package solver

import (
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

type Solver10 struct {
	SolverSelector
}

func NewSolver10() *Solver10 {
	s := &Solver10{}
	s.SolverSelector.solver = s
	return s
}

func (s *Solver10) runSimple(data []string) string {
	n := len(data)
	lights, buttons, _ := s.parse(data)

	result := 0
	for i := range n {
		l, btns := lights[i], buttons[i]
		result += s.findShortest(l, btns)
	}

	return strconv.Itoa(result)
}

func (s *Solver10) runNormal(data []string) string {
	n := len(data)
	_, buttons, joltages := s.parse(data)

	result := 0

	for m := range n {
		jlts, btns := joltages[m], buttons[m]

		an := []float64{}
		for _, btn := range btns {
			col := make([]float64, len(jlts))
			for _, l := range btn {
				col[l] = 1
			}
			an = append(an, col...)
		}

		bn := make([]float64, len(jlts))
		for i := 0; i < len(jlts); i++ {
			bn[i] = float64(jlts[i])
		}

		A := mat.NewDense(len(btns), len(jlts), an).T()
		X := mat.NewDense(len(btns), 1, nil)
		B := mat.NewDense(len(jlts), 1, bn)
		// fmt.Println(B, X)
		if err := X.Solve(A, B); err != nil {
			fmt.Println("Error solving system:", err)
		} else {
			fmt.Println(X)
		}
	}

	return strconv.Itoa(result)
}

func (s *Solver10) findShortest(l Lights, btns Buttons) int {
	start := make(State, len(l))
	dm := Distance{start.Hash(): 0}
	queue := []State{start}

	for len(queue) > 0 {
		ss := StatesSorter{States: queue, DistanceMap: dm}
		sort.Sort(ss)
		state := queue[0]
		queue = queue[1:]
		dist := dm.Get(state.Hash())

		if state.Equals(State(l)) {
			return dist
		}

		nbs := state.GetNeighbors(btns, State(l))
		for _, nb := range nbs {
			nbHash := nb.Hash()
			if !dm.Has(nbHash) {
				queue = append(queue, nb)
			}
			if dm.Get(nbHash) > dist+1 {
				dm.Set(nbHash, dist+1)
			}
		}
	}
	return -1
}

func (s *Solver10) parse(data []string) ([]Lights, []Buttons, []Joltages) {
	lights := []Lights{}
	buttons := []Buttons{}
	joltages := []Joltages{}
	for _, line := range data {
		l, b, j := s.parseOne(line)
		lights = append(lights, l)
		buttons = append(buttons, b)
		joltages = append(joltages, j)
	}
	return lights, buttons, joltages
}

func (s *Solver10) parseOne(line string) (Lights, Buttons, Joltages) {
	var lights Lights
	var buttons Buttons
	var joltages Joltages
	parts := strings.Fields(line)

	ligthPart := parts[0]
	for i := 1; i < len(ligthPart)-1; i++ {
		switch ligthPart[i] {
		case '.':
			lights = append(lights, false)
		case '#':
			lights = append(lights, true)
		}
	}

	joltagePart := strings.Trim(parts[len(parts)-1], "{}")
	subParts := strings.SplitSeq(joltagePart, ",")
	for p := range subParts {
		n, _ := strconv.Atoi(p)
		joltages = append(joltages, n)
	}

	buttonParts := parts[1 : len(parts)-1]
	for _, p := range buttonParts {
		btn := Button{}
		for sp := range strings.SplitSeq(strings.Trim(p, "()"), ",") {
			n, _ := strconv.Atoi(sp)
			btn = append(btn, n)
		}
		buttons = append(buttons, btn)
	}
	return lights, buttons, joltages
}

type Lights []bool
type Button []int
type Buttons []Button
type Joltages []int

type State []bool

func (s *State) GetNeighbors(btns Buttons, target State) []State {
	result := []State{}
	for _, btn := range btns {
		diff := s.Xor(target)
		if diff.IsRelevant(btn) {
			newState := s.Apply(btn)
			result = append(result, newState)
		}
	}
	return result
}

func (s *State) Xor(st State) State {
	result := make(State, len(st))
	for i, l := range *s {
		result[i] = l != st[i]
	}
	return result
}

func (s *State) Equals(st State) bool {
	if len(*s) != len(st) {
		return false
	}

	for i := range len(*s) {
		if (*s)[i] != st[i] {
			return false
		}
	}

	return true
}

func (s *State) IsRelevant(btn Button) bool {
	for _, l := range btn {
		if (*s)[l] {
			return true
		}
	}
	return false
}

func (s *State) Hash() string {
	var sb strings.Builder
	sb.Grow(len(*s))
	for _, b := range *s {
		if b {
			sb.WriteRune('1')
		} else {
			sb.WriteRune('0')
		}
	}
	return sb.String()
}

func (s *State) Apply(btn Button) State {
	newState := slices.Clone(*s)
	for _, i := range btn {
		newState[i] = !newState[i]
	}
	return newState
}

type StatesSorter struct {
	States      []State
	DistanceMap Distance
}

func (ss StatesSorter) Len() int      { return len(ss.States) }
func (ss StatesSorter) Swap(i, j int) { ss.States[i], ss.States[j] = ss.States[j], ss.States[i] }
func (ss StatesSorter) Less(i, j int) bool {
	iv, jv := ss.DistanceMap.Get(ss.States[i].Hash()), ss.DistanceMap.Get(ss.States[j].Hash())
	return iv < jv
}

type Distance map[string]int

func (d Distance) Get(k string) int {
	if d.Has(k) {
		return d[k]
	}
	return math.MaxInt64
}

func (d Distance) Set(k string, v int) {
	d[k] = v
}

func (d Distance) Has(k string) bool {
	_, exists := d[k]
	return exists
}
