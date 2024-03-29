package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
  "strconv"
)

type Node struct {
	x, y      int
	Number    int
	Neighbors []*Node
}

type Data struct {
  Lines []string
}

type symbol struct {
  val byte
  x, y int
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

func isSymbol(c byte) bool {
  if c == 0 {
    return false
  }
  if c >= '0' && c <= '9' {
    return false
  }
  if c == '.' {
    return false
  }
  return true
}

func (d *Data) get(x, y int) byte {
  if y < 0 || x < 0 {
    return 0
  }
  if y >= len(d.Lines) {
    return 0
  }
  if x >= len(d.Lines[y]) {
    return 0
  }
  return d.Lines[y][x]
}

func (d *Data) checkAround(i, j, k int, curNumber string) []symbol {
  var out []symbol

  if val := d.get(k, i - 1); isSymbol(val) {
    out = append(out, symbol{val, k, i - 1})
  }
  if val := d.get(k, i + 1); isSymbol(val) {
    out = append(out, symbol{val, k, i + 1})
  }
  if k == j - len(curNumber) {
    if val := d.get(k-1, i-1); isSymbol(val) {
      out = append(out, symbol{val, k - 1, i - 1})
    }
    if val := d.get(k-1, i+1); isSymbol(val) {
      out = append(out, symbol{val, k - 1, i + 1})
    }
    if val := d.get(k-1, i); isSymbol(val) {
      out = append(out, symbol{val, k - 1, i})
    }
  }
  if k == j - 1 {
    if val := d.get(k + 1, i - 1); isSymbol(val) {
      out = append(out, symbol{val, k + 1, i - 1})
    }
    if val := d.get(k + 1, i + 1); isSymbol(val) {
      out = append(out, symbol{val, k + 1, i + 1})
    }
    if val := d.get(k + 1, i); isSymbol(val) {
      out = append(out, symbol{val, k + 1, i})
    }
  }

  return out
}

func part1() {
  d, err := parse("input.prod")
  if err != nil {
    fmt.Println("error parsing")
  }
	//starNodeIndex := map[string]*Node{}

	partNumberSum := 0

	for i, line := range d.Lines {
    curNumber := ""
    // Looping through the characters and checking if they are numbers
    for j := 0; j < len(line); j++ {
      char := line[j]
      if char >= '0' && char <= '9' {
        curNumber += string(char)
        if j+1 != len(line) {
            continue
        }
        j++
      }
      if curNumber == "" {
        continue
      }
      number, err := strconv.Atoi(curNumber)
      if err != nil {
        fmt.Errorf("atoi %q; %w", char, err)
      }
      for k := j - len(curNumber); k < j; k++ {
        if val := d.checkAround(i, j, k, curNumber); len(val) > 0 {
          partNumberSum += number
          break
        }
      }
      curNumber = ""
    }
  }

	fmt.Printf("Part 1: %d\n", partNumberSum)
}

func part2() {
  d, err := parse("input.prod")
  if err != nil {
    fmt.Println("error parsing")
  }

  starNodeIndex := map[string]*Node{}
	gearRatioSum := 0

	for i, line := range d.Lines {
    curNumber := ""
    // Looping through the characters and checking if they are numbers
    for j := 0; j < len(line); j++ {
      char := line[j]
      if char >= '0' && char <= '9' {
        curNumber += string(char)
        if j+1 != len(line) {
            continue
        }
        j++
      }
      if curNumber == "" {
        continue
      }
      number, err := strconv.Atoi(curNumber)
      if err != nil {
        fmt.Errorf("atoi %q; %w", char, err)
      }
      for k := j - len(curNumber); k < j; k++ {
        if val := d.checkAround(i, j, k, curNumber); len(val) > 0 {
          for _, c := range val {
            if c.val != '*' {
              continue
            }
            starNode := starNodeIndex[fmt.Sprintf("%d/%d", c.x, c.y)]
            if starNode == nil {
              starNode = &Node {
                Number: -1,
                x: c.x,
                y: c.y,
              }
              starNodeIndex[fmt.Sprintf("%d/%d", c.x, c.y)] = starNode
            }
            starNode.Neighbors = append(starNode.Neighbors, &Node{Number: number, x: j, y: i})
          }
          break
        }
      }
      curNumber = ""
    }
  }
  for _, gearPart := range starNodeIndex {
    if len(gearPart.Neighbors) == 2 {
      gearRatioSum += gearPart.Neighbors[0].Number * gearPart.Neighbors[1].Number
    }
  }

	fmt.Printf("Part 1: %d\n", gearRatioSum)
}


func main() {
  part2()
}
