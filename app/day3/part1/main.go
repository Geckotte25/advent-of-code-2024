package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := getUncorruptedData()
	result := getResult(data)
	fmt.Println(result)
}

func getUncorruptedData() []string {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}
	corruptedMsg := string(data)
	r, _ := regexp.Compile("mul[(]{1}[0-9]{1,3},[0-9]{1,3}[)]{1}")
	return r.FindAllString(corruptedMsg, -1)
}

func getResult(data []string) int {
	result := 0
	for _, element := range data {
		res := strings.ReplaceAll(element, "mul(", "")
		res = strings.ReplaceAll(res, ")", "")
		values := strings.Split(res, ",")
		value1, err := strconv.Atoi(values[0])
		if err != nil {
			panic(err)
		}
		value2, err := strconv.Atoi(values[1])
		if err != nil {
			panic(err)
		}
		result += value1 * value2
	}
	return result
}
