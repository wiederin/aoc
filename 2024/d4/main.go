package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getVerticalSlice(crossword [][]rune, i, j, wordLen int) []rune {
	var column []rune
	for k := 0; k < wordLen; k++ {
		if i+k < len(crossword) {
			column = append(column, crossword[i+k][j])
		}
	}
	return column
}

func getRightDiagonalSlice(crossword [][]rune, i, j, wordLen int) []rune {
	var column []rune
	for k := 0; k < wordLen; k++ {
		if i+k < len(crossword) && j+k < len(crossword[i+k]) {
			column = append(column, crossword[i+k][j+k])
		}
	}
	return column
}

func getLeftDiagonalSlice(crossword [][]rune, i, j, wordLen int) []rune {
	var column []rune
	for k := 0; k < wordLen; k++ {
		if i+k < len(crossword) && j-k >= 0 {
			column = append(column, crossword[i+k][j-k])
		}
	}
	return column
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func findXMAS(crossword [][]rune) []string {
	word := "XMAS"
	wordLen := len(word)
	var positions []string

	for i := range crossword {
		for j := range crossword[i] {
			// Horizontal
			if j+wordLen <= len(crossword[i]) {
				slice := string(crossword[i][j : j+wordLen])
				if slice == word || reverse(slice) == word {
					dir := "horizontal"
					if slice == reverse(word) {
						dir = "horizontal backward"
					}
					positions = append(positions, fmt.Sprintf("%s found at row %d, col %d", dir, i, j))
				}
			}

			// Vertical
			if i+wordLen <= len(crossword) {
				verticalSlice := getVerticalSlice(crossword, i, j, wordLen)
				if string(verticalSlice) == word || reverse(string(verticalSlice)) == word {
					dir := "vertical"
					if string(verticalSlice) == reverse(word) {
						dir = "vertical backward"
					}
					positions = append(positions, fmt.Sprintf("%s found at row %d, col %d", dir, i, j))
				}
			}

			// Right diagonal
			if i+wordLen <= len(crossword) && j+wordLen <= len(crossword[i]) {
				rightDiag := getRightDiagonalSlice(crossword, i, j, wordLen)
				if string(rightDiag) == word || reverse(string(rightDiag)) == word {
					dir := "diagonal right"
					if string(rightDiag) == reverse(word) {
						dir = "diagonal right backward"
					}
					positions = append(positions, fmt.Sprintf("%s found at row %d, col %d", dir, i, j))
				}
			}

			// Left diagonal
			if i+wordLen <= len(crossword) && j-wordLen+1 >= 0 {
				leftDiag := getLeftDiagonalSlice(crossword, i, j, wordLen)
				if string(leftDiag) == word || reverse(string(leftDiag)) == word {
					dir := "diagonal left"
					if string(leftDiag) == reverse(word) {
						dir = "diagonal left backward"
					}
					positions = append(positions, fmt.Sprintf("%s found at row %d, col %d", dir, i, j))
				}
			}
		}
	}

	return positions
}

func checkForX(crossword [][]rune, i, j int) bool {
	fmt.Printf("%c   %c\n", crossword[i-1][j-1], crossword[i-1][j+1])
	fmt.Printf("  %c  \n", crossword[i][j])
	fmt.Printf("%c   %c\n", crossword[i+1][j-1], crossword[i+1][j+1])
	if ((crossword[i-1][j-1] == 'S' && crossword[i+1][j+1] == 'M') || (crossword[i-1][j-1] == 'M' && crossword[i+1][j+1] == 'S')) && ((crossword[i+1][j-1] == 'S' && crossword[i-1][j+1] == 'M') || (crossword[i+1][j-1] == 'M' && crossword[i-1][j+1] == 'S')) {
		fmt.Println("found")
		return true
	}
	return false
}

func findMAS(crossword [][]rune) []string {
	var positions []string

	for i := 1; i <= len(crossword)-2; i++ {
		for j := 1; j <= len(crossword[i])-2; j++ {
			if crossword[i][j] == 'A' {
				if checkForX(crossword, i, j) {
					positions = append(positions, fmt.Sprintf("found at row %d, col %d", i, j))
				}
			}
		}
	}

	return positions
}

func part1() {
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read input file
	scanner := bufio.NewScanner(f)
	var crossword [][]rune
	for scanner.Scan() {
		row := []rune(scanner.Text())
		crossword = append(crossword, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print crossword for verification
	for _, row := range crossword {
		fmt.Println(string(row))
	}

	// Find occurrences of XMAS
	positions := findXMAS(crossword)
	for _, pos := range positions {
		fmt.Println(pos)
	}

	fmt.Printf("Total occurrences: %d\n", len(positions))
}

func part2() {
	// Open the file
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read input file
	scanner := bufio.NewScanner(f)
	var crossword [][]rune
	for scanner.Scan() {
		row := []rune(scanner.Text())
		crossword = append(crossword, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Print crossword for verification
	for _, row := range crossword {
		fmt.Println(string(row))
	}

	// Find occurrences of XMAS
	positions := findMAS(crossword)
	for _, pos := range positions {
		fmt.Println(pos)
	}

	fmt.Printf("Total occurrences: %d\n", len(positions))
}

func main() {
	part2()
}
