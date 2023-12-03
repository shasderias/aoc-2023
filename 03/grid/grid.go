package grid

import (
	"bufio"
	"fmt"
	"io"
)

type Point struct {
	X, Y int
}

type Rune struct {
	grid       [][]rune
	maxX, maxY int
	lenX, lenY int
}

func ParseRune(r io.Reader) (*Rune, error) {
	scanner := bufio.NewScanner(r)

	buf := make([][]rune, 0)

	var firstLineLen int

	for scanner.Scan() {
		line := scanner.Text()

		l := len(line)
		if l == 0 {
			break
		}

		if firstLineLen == 0 {
			firstLineLen = l
		} else if firstLineLen != l {
			return nil, fmt.Errorf("irregularly sized grid, first line: %d != current line %d", firstLineLen, l)
		}

		lineBuf := make([]rune, l, l)

		for i, r := range line {
			lineBuf[i] = r
		}

		buf = append(buf, lineBuf)
	}

	return &Rune{
		grid: buf,
		lenX: len(buf[0]), lenY: len(buf),
		maxX: len(buf[0]) - 1, maxY: len(buf) - 1,
	}, nil
}

func (g *Rune) Index(p Point) rune { return g.grid[p.Y][p.X] }
func (g *Rune) InBounds(p Point) bool {
	if p.X < 0 || p.X > g.maxX ||
		p.Y < 0 || p.Y > g.maxY {
		return false
	}
	return true
}

func (g *Rune) LenX() int { return g.lenX }
func (g *Rune) LenY() int { return g.lenY }
func (g *Rune) MaxX() int { return g.maxX }
func (g *Rune) MaxY() int { return g.maxY }
