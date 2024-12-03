package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func part1() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	mulPattern := regexp.MustCompile(`mul\(\d+,\d+\)`)
	xyPattern := regexp.MustCompile(`\d+`)
	sum := 0 

	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())	
		muls := mulPattern.FindAllString(scanner.Text(), -1)
		fmt.Printf("muls: %s\n", muls)
		for _, mul := range muls {
			fmt.Printf("mul: %s\n", mul)	
			xy := xyPattern.FindAllString(mul, -1)
			fmt.Printf("xy: %s\n", xy)
			x, err := strconv.Atoi(xy[0])
			y, err := strconv.Atoi(xy[1])
			if err != nil {
				fmt.Println("error parsing x, y")
			}
			multiple := x * y 
			fmt.Printf("multiple: %d\n", multiple)
			sum += multiple
		}
	}

	fmt.Printf("sum: %d\n", sum)
}

func part2() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	mulPattern := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
	xyPattern := regexp.MustCompile(`\d+`)
	sum := 0 
	disabled := false

	for scanner.Scan() {
		fmt.Printf("line: %s\n", scanner.Text())	
		muls := mulPattern.FindAllString(scanner.Text(), -1)
		fmt.Printf("muls: %s\n", muls)
		for _, mul := range muls {
			fmt.Printf("mul: %s\n", mul)	
			xy := xyPattern.FindAllString(mul, -1)
			if len(xy) == 2 && !disabled {
				fmt.Printf("xy: %s\n", xy)
				x, err := strconv.Atoi(xy[0])
				y, err := strconv.Atoi(xy[1])
				if err != nil {
					fmt.Println("error parsing x, y")
				}
				multiple := x * y 
				fmt.Printf("multiple: %d\n", multiple)
				sum += multiple
			} else if mul == "don't()" {
				disabled = true
			} else if mul == "do()" {
				disabled = false
			}
		}
	}

	fmt.Printf("sum: %d\n", sum)
}

func main() {
	part2()
}
