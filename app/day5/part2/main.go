package main

import (
	"advent/app/app/day5/util"
	"fmt"
	"slices"
)

func main() {
	rules, updates := util.SplitData()
	result := getResult(rules, updates)
	fmt.Println(result)
}

func getResult(rules, updates [][]string) int {
	result := 0
	for _, update := range updates {
		if !util.IsUpdateValid(rules, update) {
			reorderInvalidUpdate(rules, update)
			result += util.GetMiddlePage(update)
		}
	}
	return result
}

func reorderInvalidUpdate(rules [][]string, update []string) []string {
	for !util.IsUpdateValid(rules, update) {
		for _, rule := range rules {
			rule1 := slices.Index(update, rule[0])
			rule2 := slices.Index(update, rule[1])
			if rule1 != -1 && rule2 != -1 && rule1 > rule2 {
				tempPage := update[rule1]
				update[rule1] = update[rule2]
				update[rule2] = tempPage
			}
		}
	}
	return update
}
