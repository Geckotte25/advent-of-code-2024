package main

import (
	"advent/app/app/day2/util"
	"fmt"
	"math"
)

func main() {
	reports := util.SplitData()
	fmt.Print(getSafeCount(reports))
}

func getSafeCount(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		if isSafe(report) {
			safeReports++
		}
	}
	return safeReports
}

func isSafe(report []int) bool {
	if report[0] > report[1] {
		return isDecreasing(report)
	} else if report[0] < report[1] {
		return isIncreasing(report)
	}
	return false
}

func isDecreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		difference := math.Abs(float64(report[i] - report[i+1]))
		if difference > 3 || report[i] <= report[i+1] {
			return false
		}
	}
	return true
}

func isIncreasing(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		difference := math.Abs(float64(report[i] - report[i+1]))
		if difference > 3 || report[i] >= report[i+1] {
			return false
		}
	}
	return true
}
