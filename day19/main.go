package main

import (
	"AOC2024/util"
	"fmt"
	"strings"
)

func main() {
	testCount, testTotalCount := day19(*util.GetFileAsLines("day19/testInput1"))
	fmt.Printf("Part 1 count, %d, totalCount %d\n", testCount, testTotalCount)

	count, totalCount := day19(*util.GetFileAsLines("day19/input.txt"))
	fmt.Printf("Part 1 count, %d, totalCount %d\n", count, totalCount)
}

func day19(inputs []string) (int, int) {
	availablePatterns, designs := parse(inputs)

	//fmt.Printf("availablePatterns: %v\n", availablePatterns)
	//fmt.Printf("designs: %v\n", designs)

	count := 0
	totalCount := 0

	for _, design := range designs {
		//fmt.Printf("design %v\n", design)
		result := designPossible(design, availablePatterns, make(map[string]int))
		if result != 0 {
			count += 1
		}
		totalCount += result

		//fmt.Printf("design %v is possible: %d\n", design, result)
	}

	return count, totalCount
}

func designPossible(design []string, patterns [][]string, cache map[string]int) int {

	designString := strings.Join(design, "")
	count, ok := cache[designString]
	if ok {
		return count
	}

	//fmt.Printf("checking design possible%v\n", design)
	if len(design) == 0 {
		return 1
	}

	possibleCount := 0
	for _, pattern := range patterns {
		if len(pattern) > len(design) {
			continue
		}
		match := true
		for i := 0; i < len(pattern) && i < len(design) && match == true; i++ {
			if pattern[i] != design[i] {
				match = false
			} else {
				continue
			}
			// if we make it here we have a match
		}
		if match {
			// here we have matched a pattern,
			patternPossible := designPossible(design[len(pattern):], patterns, cache)
			possibleCount += patternPossible
		}
	}

	cache[designString] += possibleCount
	return possibleCount
}

func parse(inputs []string) ([][]string, [][]string) {

	var availablePatterns [][]string
	var designs [][]string

	for _, split := range strings.Split(inputs[0], ",") {
		availablePatterns = append(availablePatterns, stringToSlice(strings.Trim(split, " ")))
	}

	for i := 2; i < len(inputs); i++ {
		designs = append(designs, stringToSlice(strings.Trim(inputs[i], " ")))
	}

	return availablePatterns, designs

}
func stringToSlice(input string) []string {

	var result []string

	for _, char := range input {
		result = append(result, string(char))
	}

	return result
}
