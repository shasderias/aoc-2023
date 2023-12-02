package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	matchBlue  = regexp.MustCompile(`(\d+) blue`)
	matchRed   = regexp.MustCompile(`(\d+) red`)
	matchGreen = regexp.MustCompile(`(\d+) green`)
)

func main() {
	if err := star1("input.txt"); err != nil {
		panic(err)
	}
	if err := star2("input.txt"); err != nil {
		panic(err)
	}
}

func star1(inputPath string) error {
	rawBytes, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(rawBytes)

	acc := 0

	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")

		var gameNum int
		if _, err := fmt.Sscanf(parts[0], "Game %d", &gameNum); err != nil {
			return err
		}

		draws := strings.Split(parts[1], ";")

		for _, draw := range draws {
			if findColorDraw(matchRed, draw) > 12 {
				goto nextBag
			}
			if findColorDraw(matchGreen, draw) > 13 {
				goto nextBag
			}
			if findColorDraw(matchBlue, draw) > 14 {
				goto nextBag
			}
		}

		acc += gameNum

	nextBag:
	}

	fmt.Println(acc)

	return nil
}

func star2(inputPath string) error {
	rawBytes, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(rawBytes)

	acc := 0

	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ": ")

		var gameNum int
		if _, err := fmt.Sscanf(parts[0], "Game %d", &gameNum); err != nil {
			return err
		}

		var minRed, minGreen, minBlue int

		draws := strings.Split(parts[1], ";")

		for _, draw := range draws {
			red := findColorDraw(matchRed, draw)
			if red > minRed {
				minRed = red
			}
			green := findColorDraw(matchGreen, draw)
			if green > minGreen {
				minGreen = green
			}
			blue := findColorDraw(matchBlue, draw)
			if blue > minBlue {
				minBlue = blue
			}
		}

		acc += minRed * minGreen * minBlue
	}

	fmt.Println(acc)

	return nil
}

func findColorDraw(colorRegex *regexp.Regexp, s string) int {
	matches := colorRegex.FindStringSubmatch(s)
	if len(matches) > 2 {
		panic(fmt.Errorf("color shown more than once: %s\n", s))
	}
	if len(matches) == 0 {
		return 0
	}
	count, err := strconv.Atoi(matches[1])
	if err != nil {
		panic(err)
	}
	return count
}
