package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strings"
)

type Data struct {
  Winners []int
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
    var myCardWinningNumbers []string
		line := scanner.Text()
    if line == "" {
      continue
    }
    d.Lines = append(d.Lines, line)
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
    d.Winners = append(d.Winners, len(myCardWinningNumbers))
  }
  return d, nil
}

func (d *Data)copyCards(i int, wins int, copies *map[int][]int, numCards *int) {
  *numCards += wins
  for j := 1; j <= wins; j++ {
    if i + j > len(d.Winners) {
      continue
    }
    (*copies)[i+j] = append((*copies)[i+j], d.Winners[i+j])
  }
}

func part2() {
  d, err := parse("input.prod")
  if err != nil {
    fmt.Println("error parsing")
  }

  points := 0
  copies := map[int][]int{}
  numCards := len(d.Lines)

  for i, wins := range d.Winners {
    if wins > 0 {
      cardPoints := 1
      for j := 1; j < wins; j++ {
          cardPoints *= 2
      }
      points += cardPoints
    }
    d.copyCards(i, wins, &copies, &numCards)

    for n, copy := range copies {
      if n > i {
        continue
      }
      for _, copyInst := range copy {
        d.copyCards(n, copyInst, &copies, &numCards)
      }
      delete(copies, n)
    }
  }
  fmt.Printf("Part 1: %d\n", points)
  fmt.Printf("Part 2: %d\n", numCards)
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
  part2()
}
