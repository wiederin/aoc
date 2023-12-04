package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strings"
)

type Data struct {
  Lines []string
}

func parse(fileName string) (*Data, error) {
  f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

  d := &Data{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
    if line == "" {
      continue
    }
    d.Lines = append(d.Lines, line)
  }
  return d, nil
}

func part1() {
  d, err := parse("input.prod")
  if err != nil {
    fmt.Println("error parsing")
  }

  points := 0

  for _, line := range d.Lines {
    var myCardWinningNumbers []string
    parts := strings.Split(line, ":")
    if len(parts) != 2 {
      log.Fatal("Invalid input")
    }
    cardNumbers := parts[1]
    numberParts := strings.Split(cardNumbers, "|")
    if len(numberParts) != 2 {
      log.Fatal("Invalid input")
    }
    winningNumbers := strings.TrimSpace(numberParts[0])
    myNumbers := strings.TrimSpace(numberParts[1])
    winningNumbersList := strings.Split(winningNumbers, " ")
    myNumbersList := strings.Split(myNumbers, " ")
    // for each number check if in winningNumbers
    for _, num := range myNumbersList {
      if num == "" {
        continue
      }
      for _, wNum := range winningNumbersList {
        if wNum == num {
          myCardWinningNumbers = append(myCardWinningNumbers, num)
        }
      }
    }
    cardPoints := 0
    for _, _ = range myCardWinningNumbers {
      if cardPoints == 0 {
        cardPoints = 1
        continue
      }
      cardPoints = cardPoints * 2
    }
    points += cardPoints
  }
  fmt.Printf("Part 1: %d\n", points)
}

func main() {
  part1()
}
