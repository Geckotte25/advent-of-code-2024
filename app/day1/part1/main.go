package main

import (
	"advent/app/app/day1/util"
	"fmt"
	"math"
)

func main() {
	leftValues, rightValues := util.SplitData()
	fmt.Print(getDistance(leftValues, rightValues))
}

func getDistance(leftValues []int, rightValues []int) int {
	distance := 0
	for index, value := range leftValues {
		distance = int(math.Abs(float64(value-rightValues[index]))) + distance
	}
	return distance
}
