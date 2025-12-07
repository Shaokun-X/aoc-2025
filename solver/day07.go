package solver

import (
	"fmt"
	"strconv"
)

type Solver7 struct {
	SolverSelector
}

func NewSolver7() *Solver7 {
	s := &Solver7{}
	s.SolverSelector.solver = s
	return s
}

func (s *Solver7) runSimple(data []string) string {
	width := len(data[0])

	result := 0

	input := map[int]bool{}
	for j, line := range data {
		if j%2 == 1 {
			continue
		}
		output := map[int]bool{}
		for i, char := range line {
			if char == 'S' {
				output[i] = true
			} else if input[i] {
				if char == '^' {
					result += 1
					if i-1 >= 0 {
						output[i-1] = true
					}
					if i+1 < width {
						output[i+1] = true
					}
				} else {
					output[i] = true
				}
			}
		}
		input = output
	}

	return strconv.Itoa(result)
}

func (s *Solver7) runNormal(data []string) string {
	width := len(data[0])
	height := len(data)

	nodes := []*Node{}

	// beam to its source nodes
	input := map[int][]*Node{}
	for j := 0; j < height+1; j++ {
		var line string
		if j == height {
			line = data[j-1]
		} else {
			line = data[j]
		}
		if j%2 == 1 {
			continue
		}
		output := map[int][]*Node{}
		for i, char := range line {
			if char == 'S' {
				n := NewNode(Position{j, i})
				nodes = append(nodes, n)
				output[i] = append(output[i], n)
			} else if input[i] != nil {
				// j == height means end nodes
				if char == '^' || j == height {
					n := NewNode(Position{j, i})
					parents := input[i]
					for _, p := range parents {
						p.Children = append(p.Children, n)
						n.Parents = append(n.Parents, p)
					}
					nodes = append(nodes, n)
					if i-1 >= 0 {
						output[i-1] = append(output[i-1], n)
					}
					if i+1 < width {
						output[i+1] = append(output[i+1], n)
					}
				} else {
					output[i] = append(output[i], input[i]...)
				}
			}
		}
		input = output
	}

	// from leaves to root
	for i := len(nodes) - 1; i >= 0; i-- {
		n := nodes[i]
		if len(n.Children) == 0 {
			n.Routes = 1
		} else {
			for _, ch := range n.Children {
				n.Routes += ch.Routes
			}
		}
	}

	return strconv.Itoa(nodes[0].Routes)
}

type Position struct {
	Row    int
	Column int
}

// Position is only needed for debug purpose
type Node struct {
	Position Position
	Parents  []*Node
	Children []*Node
	Routes   int
}

func NewNode(pos Position) *Node {
	n := &Node{}
	n.Routes = 0
	n.Position = pos
	return n
}

func (n *Node) String() string {
	return fmt.Sprintf("Node(%d, %d) %d", n.Position.Row+1, n.Position.Column+1, len(n.Children))
}
