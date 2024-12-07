package main

import (
	"AOC2024/util"
	"fmt"
	"strconv"
	"strings"
)

func day7(lines []string, operators []string) int {
	total := 0
	for _, line := range lines {
		total += getValidResult(line, operators)
	}

	return total
}

func getValidResult(line string, operators []string) int {

	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		fmt.Println("Invalid input format")
		return 0
	}
	testValue, _ := strconv.Atoi(parts[0])
	fields := util.StringsToInts(parts[1])

	for _, result := range getResults(fields, operators) {
		if result == testValue {
			return testValue
		}
	}
	return 0
}

func getResults(fields []int, operators []string) []int {
	var updateResultList []int
	if len(fields) == 1 {
		updateResultList = append(updateResultList, fields[0])
	}

	if len(fields) == 2 {
		for _, operator := range operators {
			switch operator {
			case "*":
				{
					updateResultList = append(updateResultList, fields[0]*fields[1])
				}
			case "+":
				{
					updateResultList = append(updateResultList, fields[0]+fields[1])
				}
			case "||":
				{
					concatenated, _ := strconv.Atoi(strconv.Itoa(fields[0]) + strconv.Itoa(fields[1]))
					updateResultList = append(updateResultList, concatenated)
				}
			}

		}

	}

	if len(fields) > 2 {
		for _, operator := range operators {
			var newFields []int
			switch operator {
			case "*":
				{
					newFields = append([]int{fields[0] * fields[1]}, fields[2:]...)
				}
			case "+":
				{
					newFields = append([]int{fields[0] + fields[1]}, fields[2:]...)
				}
			case "||":
				{
					concatenated, _ := strconv.Atoi(strconv.Itoa(fields[0]) + strconv.Itoa(fields[1]))
					newFields = append([]int{concatenated}, fields[2:]...)

				}
			}
			updateResultList = append(updateResultList, getResults(newFields, operators)...)

		}
	}
	return updateResultList

}

func main() {
	lines := *util.GetFileAsLines("day7/input.txt")
	fmt.Printf("Part1 count: %d\n", day7(lines, []string{"*", "+"}))
	fmt.Printf("Part2 count: %d\n", day7(lines, []string{"*", "+", "||"}))
}
