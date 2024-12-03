package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func SplitData() [][]int {
	reports := [][]int{}
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reportString := strings.Fields(scanner.Text())
		report := []int{}
		for _, levelString := range reportString {
			level, err := strconv.Atoi(levelString)
			if err != nil {
				panic(err)
			}
			report = append(report, level)
		}
		reports = append(reports, report)
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
	return reports
}
