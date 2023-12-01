package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
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
		var d1, d2 string

		for _, char := range scanner.Text() {
			if char >= '0' && char <= '9' {
				d2 = string(char)
				if d1 == "" {
					d1 = string(char)
				}
			}
		}

		num, err := strconv.Atoi(d1 + d2)
		if err != nil {
			return err
		}

		acc += num
	}

	fmt.Println(acc)
	return nil
}

var digitMap = map[string]string{
	"1":     "1",
	"2":     "2",
	"3":     "3",
	"4":     "4",
	"5":     "5",
	"6":     "6",
	"7":     "7",
	"8":     "8",
	"9":     "9",
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
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

		line := scanner.Bytes()

		p1 := 0
		p2 := min(len(line), 5)
		l := len(line)

		var d1, d2 string

		for p1 < l {
			snip := line[p1:p2]

			d := findDigit(snip)

			if d != "" {
				d2 = d
				if d1 == "" {
					d1 = d
				}
			}

			p1++
			if p2 < l {
				p2++
			}

		}

		num, err := strconv.Atoi(d1 + d2)
		if err != nil {
			return err
		}

		acc += num
	}

	fmt.Println(acc)
	return nil
}

func findDigit(buf []byte) string {
	for str, digit := range digitMap {
		if bytes.HasPrefix(buf, []byte(str)) {
			return digit
		}
	}
	return ""
}
