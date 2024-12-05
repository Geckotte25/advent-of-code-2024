package util

import "slices"

func IsUpdateValid(rules [][]string, update []string) bool {
	illegalValue := []string{}
	for _, page := range update {
		for _, rule := range rules {
			if rule[0] == page && slices.Contains(illegalValue, rule[1]) {
				return false
			}
		}
		illegalValue = append(illegalValue, page)
	}

	illegalValue = []string{}
	for i := len(update) - 1; i >= 0; i-- {
		page := update[i]
		for _, rule := range rules {
			if rule[1] == page && slices.Contains(illegalValue, rule[0]) {
				return false
			}
		}
		illegalValue = append(illegalValue, page)
	}
	return true
}
