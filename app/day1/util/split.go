package util

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SplitData() ([]int, []int) {
	leftValues := []int{}
	rightValues := []int{}

	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())

		leftValue, err := strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}

		rightValue, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		leftValues = append(leftValues, leftValue)
		rightValues = append(rightValues, rightValue)

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}

	sort.Ints(leftValues)
	sort.Ints(rightValues)

	return leftValues, rightValues
}
