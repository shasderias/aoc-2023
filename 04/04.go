package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	if err := star1(); err != nil {
		panic(err)
	}

	if err := star2(); err != nil {
		panic(err)
	}
}

func star1() error {
	acc := 0

	scanner := bufio.NewScanner(strings.NewReader(input))

	lineRegexp := regexp.MustCompile(`Card\s*(\d+):(.*)\|(.*)`)

	for scanner.Scan() {
		matches := lineRegexp.FindStringSubmatch(scanner.Text())
		if len(matches) != 4 {
			return fmt.Errorf("unexpected input, got %d matches: %s", len(matches), scanner.Text())
		}
		winNums := numStrToInts(matches[2])
		cardNums := numStrToInts(matches[3])

		var i, j, wins int

		for i < len(winNums) && j < len(cardNums) {
			if winNums[i] == cardNums[j] {
				wins++
				i++
				j++
				continue
			}

			if winNums[i] < cardNums[j] {
				i++
				continue
			}

			if winNums[i] > cardNums[j] {
				j++
				continue
			}
		}

		if wins > 0 {
			acc += pow(2, wins-1)
		}
	}

	fmt.Println(acc)

	return nil
}

func star2() error {
	scanner := bufio.NewScanner(strings.NewReader(input))

	lineRegexp := regexp.MustCompile(`Card\s*(\d+):(.*)\|(.*)`)

	cards := make([]*Card, 0)

	for scanner.Scan() {
		matches := lineRegexp.FindStringSubmatch(scanner.Text())
		if len(matches) != 4 {
			return fmt.Errorf("unexpected input, got %d matches: %s", len(matches), scanner.Text())
		}
		cardNum, err := strconv.Atoi(matches[1])
		if err != nil {
			return err
		}
		winNums := numStrToInts(matches[2])
		cardNums := numStrToInts(matches[3])

		var i, j, wins int

		for i < len(winNums) && j < len(cardNums) {
			if winNums[i] == cardNums[j] {
				wins++
				i++
				j++
				continue
			}

			if winNums[i] < cardNums[j] {
				i++
				continue
			}

			if winNums[i] > cardNums[j] {
				j++
				continue
			}
		}

		cards = append(cards, &Card{
			num:   cardNum,
			wins:  wins,
			count: 1,
		})
	}

	queue := make([]*Card, 0, len(cards))

	for _, c := range cards {
		queue = append(queue, c)
	}

	for len(queue) > 0 {
		card := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if card.wins > 0 {
			for i := card.num; i < card.num+card.wins; i++ {
				cards[i].count++
				queue = append(queue, cards[i])
			}
		}
	}

	acc := 0
	for _, c := range cards {
		acc += c.count
	}
	fmt.Println(acc)

	return nil
}

type Card struct {
	num   int
	wins  int
	count int
}

func numStrToInts(s string) []int {
	numStr := strings.Split(s, " ")
	nums := make([]int, 0, len(numStr))
	for _, ns := range numStr {
		if ns == "" {
			continue
		}
		num, err := strconv.Atoi(strings.TrimSpace(ns))
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}

	sort.Ints(nums)
	return nums
}
func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
