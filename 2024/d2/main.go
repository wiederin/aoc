package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	f, err := os.Open("example_input")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0

	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())
		safe := true
		increasing := true
		digits := strings.Split(scanner.Text(), " ")
		fmt.Printf("digits: %s\n", digits)
		for i, digitString := range digits {
			if len(digits) == i+1 {
				break
			}
			digit, err := strconv.Atoi(digitString)
			next, err := strconv.Atoi(digits[i+1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("digit: %d\n", digit)
			if i == 0 {
				if next < digit {
					increasing = false
				}
			}
			if increasing {
				if next <= digit {
					safe = false
				} else {
					if next-digit > 3 {
						safe = false
					}
				}
			} else {
				if next >= digit {
					safe = false
				} else {
					if digit-next > 3 {
						safe = false
					}
				}
			}

		}
		fmt.Printf("increasing: %t\n", increasing)
		fmt.Printf("safe: %t\n", safe)
		if safe {
			count += 1
		}
	}
	fmt.Printf("count: %d\n", count)
}

func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	increasing := levels[1] > levels[0]
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		if diff > 3 || diff < -3 {
			return false
		}
		if increasing && levels[i+1] <= levels[i] {
			return false
		}
		if !increasing && levels[i+1] >= levels[i] {
			return false
		}
	}
	return true
}

func part2() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		strLevels := strings.Split(line, " ")
		levels := make([]int, len(strLevels))

		for i, s := range strLevels {
			levels[i], err = strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
		}

		//fmt.Printf("Checking report: %v\n", levels)

		if isSafe(levels) {
			//fmt.Printf("Report is safe as-is: %v\n", levels)
			count++
			continue
		}

		safeWithDampener := false
		for i := 0; i < len(levels); i++ {
			subLevels := append([]int{}, levels[:i]...)
			subLevels = append(subLevels, levels[i+1:]...)

			//fmt.Printf("  Testing subsequence (removing %d): %v\n", levels[i], subLevels)

			if isSafe(subLevels) {
				//fmt.Printf("  Subsequence is safe: %v\n", subLevels)
				safeWithDampener = true
				break
			}
		}

		if safeWithDampener {
			//fmt.Printf("Report is safe with dampener: %v\n", levels)
			count++
		} else {
			//fmt.Printf("Report is unsafe: %v\n", levels)
		}
	}

	fmt.Printf("Count of safe reports: %d\n", count)
}

func main() {
	part2()
}
