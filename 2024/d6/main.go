package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func guardMove(col, row int, direction string, guardMap map[int]string) (newCol, newRow int, newDirection string, onMap bool) {
	switch direction {
	case "up":
		if row == 0 {
			fmt.Println("out of bounds at the bottom")
			return col, row, direction, false
		}
		if guardMap[row-1][col] != '#' {
			fmt.Println("move up")
			return col, row - 1, direction, true
		}
		fmt.Println("turn")
		if col+1 == len(guardMap[row])-1 {
			fmt.Println("out of bounds at the right")
			return col, row, direction, false
		}
		if guardMap[row][col+1] != '#' {
			fmt.Println("move right")
			return col + 1, row, "right", true
		}
		if row == len(guardMap) {
			fmt.Println("out of bounds at the bottom")
			return col, row, direction, false
		}
		if guardMap[row+1][col] != '#' {
			fmt.Println("move down")
			return col, row + 1, "down", true
		}
		if col == 0 {
			fmt.Println("out of bounds at the left")
			return col, row, direction, false
		}
		if guardMap[row][col-1] != '#' {
			fmt.Println("move left")
			return col - 1, row, "down", true
		}
		return col, row, "right", true
	case "right":
		if col == len(guardMap[row])-1 {
			fmt.Println("out of bounds at the right")
			return col, row, direction, false
		}
		if guardMap[row][col+1] != '#' {
			fmt.Println("move right")
			return col + 1, row, direction, true
		}
		fmt.Println("turn")
		return col, row, "down", true
	case "down":
		if row == len(guardMap)-1 {
			fmt.Println("out of bounds at the top")
			return col, row, direction, false
		}
		if guardMap[row+1][col] != '#' {
			fmt.Println("move down")
			return col, row + 1, direction, true
		}
		fmt.Println("turn right")
		return col, row, "left", true
	case "left":
		if col == 0 {
			fmt.Println("out of bounds at the left")
			return col, row, direction, false
		}
		if guardMap[row][col-1] != '#' {
			fmt.Println("move left")
			return col - 1, row, direction, true
		}
		fmt.Println("turn up")
		return col, row, "up", true
	}
	return col, row, direction, false
}

func main() {
	f, err := os.Open("example_input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	guardMap := make(map[int]string)
	count := 1
	initGuardPosRow := 0
	initGuardPosCol := 0
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		guardMap[row] = line
		// find ^
		for col, car := range line {
			if car == '^' {
				initGuardPosRow = row
				initGuardPosCol = col
			}
		}

		row++
	}
	fmt.Printf("initial guard position: %d, %d\n", initGuardPosCol, initGuardPosRow)

	direction := "up"
	visited := make(map[string]bool)
	for {
		posKey := fmt.Sprintf("%d,%d", initGuardPosCol, initGuardPosRow)
		fmt.Println(posKey)
		if visited[posKey] {
			fmt.Println("inifinte loop detected, exiting")
		}
		visited[posKey] = true
		guardPosCol, guardPosRow, newDirection, onMap := guardMove(initGuardPosCol, initGuardPosRow, direction, guardMap)
		fmt.Printf("%d, %d, %s\n", initGuardPosCol, initGuardPosRow, direction)
		if onMap {
			initGuardPosCol = guardPosCol
			initGuardPosRow = guardPosRow
			direction = newDirection
			count++
		} else {
			fmt.Println("off map")
			break
		}
	}

	fmt.Println("count ", len(visited))
}
