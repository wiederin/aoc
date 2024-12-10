package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func checkUpdate(update []string, rules map[int][]string) bool {
	valid := true
	for i, page := range update {
		if i >= 1 {
			num, err := strconv.Atoi(page)
			if err != nil {
				fmt.Println("error parsing page number")
				continue
			}
			var before []string = update[0:i]
			cantBeBefore := rules[num]
			for _, loc := range before {
				if slices.Contains(cantBeBefore, loc) {
					valid = false
				}
			}
		}
	}
	return valid
}

func reorder(slice []string, i, j int) []string {
	if i < 0 || i >= len(slice) || j < 0 || j >= len(slice) {
		return slice
	}

	element := slice[i]
	slice = append(slice[:i], slice[i+1:]...)
	slice = append(slice[:j], append([]string{element}, slice[j:]...)...)
	return slice
}

func fixUpdate(update []string, rules map[int][]string) []string {
	for i, page := range update {
		if i >= 1 {
			num, err := strconv.Atoi(page)
			if err != nil {
				fmt.Println("error parsing page number")
				continue
			}
			var before []string = update[0:i]
			cantBeBefore := rules[num]
			for j, loc := range before {
				if slices.Contains(cantBeBefore, loc) {
					proposedUpdate := reorder(update, i, j)
					if checkUpdate(proposedUpdate, rules) {
						return proposedUpdate
					} else {
						return fixUpdate(proposedUpdate, rules)
					}
				}
			}
		}
	}
	return update
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	rulePattern := regexp.MustCompile(`\d+|\d+`)
	updatePattern := regexp.MustCompile(`\d+`)

	rulesFlag := true
	rules := make(map[int][]string)
	sum := 0
	sum2 := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			rulesFlag = false
			continue
		}
		if rulesFlag {
			ruleStrings := rulePattern.FindAllString(scanner.Text(), -1)
			if len(ruleStrings) != 2 {
				continue
			} else {
				key, err := strconv.Atoi(ruleStrings[0])
				if err != nil {
					fmt.Println("error parsing rule")
				}
				rules[key] = append(rules[key], ruleStrings[1])
			}
		} else {
			updateStrings := updatePattern.FindAllString(scanner.Text(), -1)
			valid := checkUpdate(updateStrings, rules)
			if valid {
				// find middle
				loc := (len(updateStrings) - 1) / 2
				middle, err := strconv.Atoi(updateStrings[loc])
				if err != nil {
					fmt.Println("error finding middle")
					continue
				}
				sum += middle
			} else {
				fmt.Printf("update: %s\n", updateStrings)
				validUpdate := fixUpdate(updateStrings, rules)
				fmt.Printf("valid update: %s\n", validUpdate)
				loc := (len(updateStrings) - 1) / 2
				middle, err := strconv.Atoi(updateStrings[loc])
				if err != nil {
					fmt.Println("error finding middle")
					continue
				}
				sum2 += middle
			}
		}

	}
	fmt.Println("sum: ", sum)
	fmt.Println("sum2: ", sum2)
}
