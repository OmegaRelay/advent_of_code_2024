// Word search for xmas
// This word search allows words to be horizontal, vertical, diagonal, written
// backwards, or even overlapping other words.
//
// Solution: make a 2D slice where the top level value indicates row, and the
// second level value indicates column. Then iterate through the rows and
// columns, if x is found, check the surrounding for m then again for a and so forth
//
// Part 1: complete
// Part 2: wip
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type wordSearch [][]string

func main() {
	if len(os.Args) < 2 {
		log.Fatal("error no input")
	}

	inputFile := os.Args[1]

	if inputFile == "" {
		log.Fatal("error no input")
	}

	ws, err := getWordSearchFromFile(inputFile)
	if err != nil {
		log.Fatalf("could not read input file: %s", err.Error())
	}

	startTime := time.Now().UnixMicro()
	ans := part1(ws)
	ansTime := time.Now().UnixMicro() - startTime
	fmt.Printf("Part 1: \n\tans: %d\n\tin %dus\n\n", ans, ansTime)

	startTime = time.Now().UnixMicro()
	ans = part2(ws)
	ansTime = time.Now().UnixMicro() - startTime
	fmt.Printf("Part 2: \n\tans: %d\n\tin %dus\n\n", ans, ansTime)
}

func part1(ws wordSearch) int {
	found := 0
	for row, line := range ws {
		for col := range line {
			found += ws.findXmas(row, col)
		}
	}
	return found
}

func part2(ws wordSearch) int {
	matchedRows := []int{}
	matchedCols := []int{}

	found := 0
	for row, line := range ws {
		for col := range line {
			for i, matchedRow := range matchedRows {
				if row == matchedRow && col == matchedCols[i] {
					continue
				}
			}

			found += ws.findCrossedMas(row, col, matchedRows, matchedCols)
		}
	}
	return found
}

// Finds all occurrences of the string "XMAS" in the word search at the row,col given and returns
// the number found
func (ws wordSearch) findXmas(row, col int) int {
	found := 0
	for xFactor := -1; xFactor <= 1; xFactor++ {
		for yFactor := -1; yFactor <= 1; yFactor++ {
			if ws.findStringInDir("XMAS", row, col, xFactor, yFactor) {
				found++
			}
		}
	}
	return found
}

func (ws wordSearch) findCrossedMas(row int, col int, matchedRows, matchedCols []int) int {
	const kMasStr = "MAS"

	factors := []int{-1, 1}

	found := 0
	for _, xFactor := range factors {
		for _, yFactor := range factors {
			if ws.findStringInDir(kMasStr, row, col, xFactor, yFactor) {
				newCol := col + ((len(kMasStr) - 1) * xFactor)
				newRow := row + ((len(kMasStr) - 1) * yFactor)
				if ws.findStringInDir(kMasStr, row, newCol, xFactor*-1, yFactor) {
					matchedRows = append(matchedRows, row)
					matchedCols = append(matchedCols, newCol)
					found++
				}

				if ws.findStringInDir(kMasStr, newRow, col, xFactor, yFactor*-1) {
					matchedRows = append(matchedRows, newRow)
					matchedCols = append(matchedCols, col)
					found++
				}
			}
		}
	}
	return found
}

// Match a string at a row,col point in the wordsearch in a particular direction
func (ws wordSearch) findStringInDir(str string, row, col, xFactor, yFactor int) bool {
	chars := strings.Split(str, "")

	// Sanitize
	if row >= len(ws) || row < 0 || col >= len(ws[row]) || col < 0 {
		return false
	}

	nrChar := len(chars) - 1
	if (row+(nrChar*yFactor) >= len(ws)) || (col+(nrChar*xFactor)) >= len(ws[row]) ||
		(col+(nrChar*xFactor)) < 0 || (row+(nrChar*yFactor)) < 0 {
		return false
	}

	for i, char := range chars {
		if ws[row+(i*yFactor)][col+(i*xFactor)] != char {
			return false
		}
	}
	return true
}

func getWordSearchFromFile(filePath string) (wordSearch, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	ws := wordSearch{}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			break
		}

		letters := strings.Split(line, "")
		ws = append(ws, letters)
	}
	return ws, nil
}
