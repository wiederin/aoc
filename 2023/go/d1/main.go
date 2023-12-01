package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "regexp"
    "strconv"
)

func main() {
    part2()
}


func part1() {
    f, err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    pattern := regexp.MustCompile(`[0-9]`) 
    sum := 0
    for scanner.Scan() {
        fmt.Printf("line: %s\n", scanner.Text())
        digits := pattern.FindAllString(scanner.Text(), -1)
        fmt.Printf("digits: %s\n", digits)
        calibrationValue, err := strconv.Atoi(digits[0] + digits[len(digits)-1])        
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("calibrationValue: %d\n", calibrationValue)
        sum = sum + calibrationValue
        fmt.Printf("sum: %d\n", sum)
    }
    fmt.Printf("final sum: %d\n", sum)

    if err := scanner.Err(); err!= nil {
        log.Fatal(err)
        fmt.Println("here")
    }
}

func part2() {
	replacements := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
        "1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}
    f, err := os.Open("input")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()
    scanner := bufio.NewScanner(f)
    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        var first, last int
    fromStart:
    	for i := 0; i < len(line); i++ {
			for pattern, n := range replacements {
				if i+len(pattern) > len(line) {
					continue
				}
				if line[i:i+len(pattern)] == pattern {
					first = n
					break fromStart 
				}
			}
		}

    fromEnd:
		for i := len(line) - 1; i >= 0; i-- {
			for pattern, n := range replacements {
				if i-len(pattern)+1 < 0 {
					continue
				}
				if line[i-len(pattern)+1:i+1] == pattern {
					last = n
					break fromEnd 
				}
			}
		}
		sum += first*10 + last 
	}
    fmt.Printf("final sum: %d\n", sum)

    if err := scanner.Err(); err!= nil {
        log.Fatal(err)
    }
    
}
