package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type Reader struct {
	BasePath string
}

func (p *Reader) ReadExample(day int) []string {
	path := p.buildPath(day, false)
	return p.readLines(path)
}

func (p *Reader) ReadReal(day int) []string {
	path := p.buildPath(day, true)
	return p.readLines(path)
}

func (p *Reader) readLines(path string) []string {
	lines := []string{}

	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if line != "" {
				lines = append(lines, line)
			}
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
	return lines
}

func (p *Reader) buildPath(day int, real bool) string {
	if real {
		return filepath.Join(p.BasePath, fmt.Sprintf("%d-real.txt", day))
	}
	return filepath.Join(p.BasePath, fmt.Sprintf("%d.txt", day))
}
