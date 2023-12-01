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
    }


}
