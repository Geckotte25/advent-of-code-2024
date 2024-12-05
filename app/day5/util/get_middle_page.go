package util

import "strconv"

func GetMiddlePage(update []string) int {
	middlePageString := update[len(update)/2]
	middlePageValue, err := strconv.Atoi(middlePageString)
	if err != nil {
		panic(err)
	}
	return middlePageValue
}
