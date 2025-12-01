package solver

type Solver1 struct {
	SolverSelector
}

func NewSolver1() *Solver1 {
	s := &Solver1{}
	s.SolverSelector.solver = s

	return s
}

func (s *Solver1) runSimple(data []string) string {
	return "simple"
}

func (s *Solver1) runNormal(data []string) string {
	return "normal"
}
