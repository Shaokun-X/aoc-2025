package solver

type Solver interface {
	runSimple([]string) string
	runNormal([]string) string
}

type SolverSelector struct {
	solver Solver
}

type config struct {
	normal bool
}

type RunOption func(*config)

func WithNormal() RunOption {
	return func(c *config) {
		c.normal = true
	}
}

func (ss *SolverSelector) Run(data []string, opts ...RunOption) string {
	cfg := config{
		normal: false,
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	if cfg.normal {
		return ss.solver.runNormal(data)
	}
	return ss.solver.runSimple(data)
}
