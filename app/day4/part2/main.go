package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	total := findXMAS()
	fmt.Println(total)
}

func findXMAS() int {
	totalXMAS := 0
	lines := []string{}
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
	for i, line := range lines {
		for j := range line {
			if i >= 1 && i <= len(line)-2 && j >= 1 && j <= len(line)-2 {
				totalXMAS += findDiagonal(i, j, lines)
			}
		}
	}
	return totalXMAS
}

func findDiagonal(lineIndex, xIndex int, lines []string) int {
	diagonal1 := false
	diagonal2 := false
	if lines[lineIndex][xIndex] == 'A' && (lines[lineIndex-1][xIndex-1] == 'M' && lines[lineIndex+1][xIndex+1] == 'S' || lines[lineIndex-1][xIndex-1] == 'S' && lines[lineIndex+1][xIndex+1] == 'M') {
		diagonal1 = true
	}
	if lines[lineIndex][xIndex] == 'A' && (lines[lineIndex-1][xIndex+1] == 'M' && lines[lineIndex+1][xIndex-1] == 'S' || lines[lineIndex-1][xIndex+1] == 'S' && lines[lineIndex+1][xIndex-1] == 'M') {
		diagonal2 = true
	}
	if diagonal1 && diagonal2 {
		return 1
	}
	return 0
}
