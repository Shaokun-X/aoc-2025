package main

import (
	"flag"
	"fmt"
	"os"
	solver "shaokun-x/aoc-2025/solver"
	reader "shaokun-x/aoc-2025/utils"
	"strconv"
)

type Selector interface {
	Run([]string, ...solver.RunOption) string
}

var solverRegistry = map[int]Selector{
	1: solver.NewSolver1(),
	2: solver.NewSolver2(),
	3: solver.NewSolver3(),
}

func main() {
	// get options
	normal := flag.Bool("n", false, "Whether to run the program for the normal challenge (the golden star).")
	real := flag.Bool("r", false, "Whether to run the program against the real dataset.")
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Error: Number of day is required.")
		os.Exit(1)
	}
	day, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	// read data
	rdr := reader.Reader{BasePath: "./data"}
	var data []string
	if *real {
		data = rdr.ReadReal(day)
	} else {
		data = rdr.ReadExample(day)
	}

	// call solver
	var result string

	slvr, exists := solverRegistry[day]

	if !exists {
		fmt.Printf("Error: Solver not found for day %d.\n", day)
		os.Exit(1)
	}
	if *normal {
		result = slvr.Run(data, solver.WithNormal())
	} else {
		result = slvr.Run(data)
	}

	fmt.Println(result)
}
