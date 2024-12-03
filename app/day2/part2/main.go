package main

import (
	"advent/app/app/day2/util"
	"fmt"
	"math"
)

func main() {
	reports := util.SplitData()
	fmt.Println(getSafeCount(reports))
}

func getSafeCount(reports [][]int) int {
	safeReports := 0
	for _, report := range reports {
		if isSafe(report, false) {
			safeReports++
		}
	}
	return safeReports
}

func isSafe(report []int, isRemoved bool) bool {
	if report[0] > report[1] {
		_, decreasing := isDecreasing(report)
		if decreasing {
			return true
		}
	} else if report[0] < report[1] {
		_, increasing := isIncreasing(report)
		if increasing {
			return true
		}
	}

	if !isRemoved {
		for index := range report {
			reportwithRemovedLevel := make([]int, len(report))
			copy(reportwithRemovedLevel, report)
			reportwithRemovedLevel = append(reportwithRemovedLevel[:index], reportwithRemovedLevel[index+1:]...)
			if isSafe(reportwithRemovedLevel, true) {
				return true
			}
		}
	}

	return false
}

func isDecreasing(report []int) (int, bool) {
	for i := 0; i < len(report)-1; i++ {
		difference := math.Abs(float64(report[i] - report[i+1]))
		if difference > 3 || report[i] <= report[i+1] {
			return i, false
		}
	}
	return 0, true
}

func isIncreasing(report []int) (int, bool) {
	for i := 0; i < len(report)-1; i++ {
		difference := math.Abs(float64(report[i] - report[i+1]))
		if difference > 3 || report[i] >= report[i+1] {
			return i, false
		}
	}
	return 0, true
}
