package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"

	"github.com/shasderias/aoc-2023/03/grid"
)

var (
	PointNil = grid.Point{-1, -1}
)

func main() {
	if err := star1("input.txt"); err != nil {
		panic(err)
	}
	if err := star2("input.txt"); err != nil {
		panic(err)
	}
}

type Number struct {
	P0, P1 grid.Point
	Num    int
}

func star1(inputPath string) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	g, err := grid.ParseRune(f)
	if err != nil {
		return err
	}

	numbers, err := findNumbers(g)
	if err != nil {
		return err
	}

	acc := 0

	for _, num := range numbers {
		for y := num.P0.Y - 1; y <= num.P1.Y+1; y++ {
			for x := num.P0.X - 1; x <= num.P1.X+1; x++ {
				if !g.InBounds(grid.Point{x, y}) {
					continue
				}
				pt := g.Index(grid.Point{x, y})
				if unicode.IsDigit(pt) || pt == '.' {
					continue
				}

				acc += num.Num
				goto nextNum
			}
		}
	nextNum:
	}

	fmt.Println(acc)

	return nil
}

func star2(inputPath string) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	g, err := grid.ParseRune(f)
	if err != nil {
		return err
	}

	numbers, err := findNumbers(g)
	if err != nil {
		return err
	}

	gearPoints := make([]grid.Point, 0)

	for y := 0; y < g.LenY(); y++ {
		for x := 0; x < g.LenX(); x++ {
			if g.Index(grid.Point{x, y}) == '*' {
				gearPoints = append(gearPoints, grid.Point{x, y})
			}
		}
	}

	acc := 0

	for _, gear := range gearPoints {
		numMap := map[int]struct{}{}

		for y := gear.Y - 1; y <= gear.Y+1; y++ {
			for x := gear.X - 1; x <= gear.X+1; x++ {
				if !g.InBounds(grid.Point{x, y}) {
					continue
				}
				if x == gear.X && y == gear.Y {
					continue
				}
				num, ok := findNumber(grid.Point{x, y}, numbers)
				if ok {
					numMap[num.Num] = struct{}{}
				}
			}
		}

		if len(numMap) == 2 {
			ratio := 0
			for num := range numMap {
				if ratio == 0 {
					ratio = num
				} else {
					ratio *= num
				}
			}
			acc += ratio
		}
	}

	fmt.Println(acc)

	return nil
}

func pointOnLine(pt, l0, l1 grid.Point) bool {
	if pt.Y != l0.Y {
		return false
	}
	if pt.X < l0.X || pt.X > l1.X {
		return false
	}
	return true
}

func findNumber(pt grid.Point, numbers []Number) (Number, bool) {
	for _, num := range numbers {
		if pointOnLine(pt, num.P0, num.P1) {
			return num, true
		}
	}
	return Number{}, false
}

func findNumbers(g *grid.Rune) ([]Number, error) {
	numbers := make([]Number, 0)

	for y := 0; y < g.LenY(); y++ {
		numRunStart := PointNil
		numStr := ""
		for x := 0; x < g.LenX(); x++ {
			pt := g.Index(grid.Point{x, y})

			if unicode.IsDigit(pt) {
				numStr += string(pt)
				if numRunStart == PointNil {
					numRunStart = grid.Point{x, y}
				}
			} else {
				if len(numStr) > 0 {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						return nil, err
					}
					numbers = append(numbers, Number{
						P0:  numRunStart,
						P1:  grid.Point{x - 1, y},
						Num: num,
					})
					numRunStart = PointNil
					numStr = ""
				}
			}
		}
		if len(numStr) > 0 {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, Number{
				P0:  numRunStart,
				P1:  grid.Point{g.MaxX(), y},
				Num: num,
			})
			numRunStart = PointNil
			numStr = ""
		}
	}

	return numbers, nil
}
