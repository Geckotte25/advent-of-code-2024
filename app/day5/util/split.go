package util

import (
	"bufio"
	"os"
	"strings"
)

func SplitData() ([][]string, [][]string) {
	isRule := true
	rules := [][]string{}
	updates := [][]string{}

	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			isRule = false
			continue
		}

		if isRule {
			rules = append(rules, strings.Split(scanner.Text(), "|"))
		} else {
			updates = append(updates, strings.Split(scanner.Text(), ","))
		}
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
	return rules, updates
}
