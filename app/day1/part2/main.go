package main

import (
	"advent/app/app/day1/util"
	"fmt"
)

func main() {
	leftValues, rightValues := util.SplitData()
	fmt.Print(getSimilarityScore(leftValues, rightValues))
}

func getSimilarityScore(leftValues []int, rightValues []int) int {
	similarityScore := 0
	for _, leftValue := range leftValues {
		similarity := 0
		for _, rightValue := range rightValues {
			if leftValue == rightValue {
				similarity++
			}
		}
		similarityScore = similarityScore + (leftValue * similarity)
	}
	return similarityScore
}
