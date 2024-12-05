package main

import (
	"advent/app/app/day5/util"
	"fmt"
)

func main() {
	rules, updates := util.SplitData()
	result := getResult(rules, updates)
	fmt.Println(result)
}

func getResult(rules, updates [][]string) int {
	result := 0
	for _, update := range updates {
		if util.IsUpdateValid(rules, update) {
			result += util.GetMiddlePage(update)
		}
	}
	return result
}
