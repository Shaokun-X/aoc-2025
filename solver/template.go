package solver

type SolverTemplate struct {
	SolverSelector
}

func NewSolverTemplate() *SolverTemplate {
	s := &SolverTemplate{}
	s.SolverSelector.solver = s
	return s
}

func (s *SolverTemplate) runSimple(data []string) string {
	return ""
}

func (s *SolverTemplate) runNormal(data []string) string {
	return ""
}
