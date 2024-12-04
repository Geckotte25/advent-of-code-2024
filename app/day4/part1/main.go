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
			if j >= 3 {
				totalXMAS += findLeft(i, j, lines)
			}
			if j <= len(line)-4 {
				totalXMAS += findRight(i, j, lines)
			}
			if i >= 3 {
				totalXMAS += findTop(i, j, lines)
				if j >= 3 {
					totalXMAS += findDiagonalTopLeft(i, j, lines)
				}
				if j <= len(line)-4 {
					totalXMAS += findDiagonalTopRight(i, j, lines)
				}
			}
			if i <= len(lines)-4 {
				totalXMAS += findBottom(i, j, lines)
				if j >= 3 {

					totalXMAS += findDiagonalBottomLeft(i, j, lines)
				}
				if j <= len(line)-4 {
					totalXMAS += findDiagonalBottomRight(i, j, lines)
				}
			}
		}
	}
	return totalXMAS
}

func findDiagonalTopLeft(lineIndex, xIndex int, lines []string) int {
	total := 0
	if lines[lineIndex][xIndex] == 'X' && lines[lineIndex-1][xIndex-1] == 'M' && lines[lineIndex-2][xIndex-2] == 'A' && lines[lineIndex-3][xIndex-3] == 'S' {
		total++
	}
	return total
}

func findDiagonalTopRight(lineIndex, xIndex int, lines []string) int {
	total := 0
	if lines[lineIndex][xIndex] == 'X' && lines[lineIndex-1][xIndex+1] == 'M' && lines[lineIndex-2][xIndex+2] == 'A' && lines[lineIndex-3][xIndex+3] == 'S' {
		total++
	}
	return total
}

func findDiagonalBottomLeft(lineIndex, xIndex int, lines []string) int {
	total := 0
	if lines[lineIndex][xIndex] == 'X' && lines[lineIndex+1][xIndex-1] == 'M' && lines[lineIndex+2][xIndex-2] == 'A' && lines[lineIndex+3][xIndex-3] == 'S' {
		total++
	}
	return total
}

func findDiagonalBottomRight(lineIndex, xIndex int, lines []string) int {
	total := 0
	if lines[lineIndex][xIndex] == 'X' && lines[lineIndex+1][xIndex+1] == 'M' && lines[lineIndex+2][xIndex+2] == 'A' && lines[lineIndex+3][xIndex+3] == 'S' {
		total++
	}
	return total
}

func findTop(lineIndex, xIndex int, lines []string) int {
	total := 0
	if lines[lineIndex][xIndex] == 'X' && lines[lineIndex-1][xIndex] == 'M' && lines[lineIndex-2][xIndex] == 'A' && lines[lineIndex-3][xIndex] == 'S' {
		total++
	}
	return total
}

func findBottom(lineIndex, xIndex int, lines []string) int {
	total := 0
	if lines[lineIndex][xIndex] == 'X' && lines[lineIndex+1][xIndex] == 'M' && lines[lineIndex+2][xIndex] == 'A' && lines[lineIndex+3][xIndex] == 'S' {
		total++
	}
	return total
}

func findLeft(lineIndex, xIndex int, lines []string) int {
	total := 0
	if lines[lineIndex][xIndex] == 'X' && lines[lineIndex][xIndex-1] == 'M' && lines[lineIndex][xIndex-2] == 'A' && lines[lineIndex][xIndex-3] == 'S' {
		total++
	}
	return total
}

func findRight(lineIndex, xIndex int, lines []string) int {
	total := 0
	if lines[lineIndex][xIndex] == 'X' && lines[lineIndex][xIndex+1] == 'M' && lines[lineIndex][xIndex+2] == 'A' && lines[lineIndex][xIndex+3] == 'S' {
		total++
	}
	return total
}
