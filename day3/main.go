package main

import (
	"AOC2024/util"
	"fmt"
	"regexp"
	"strconv"
)

func part1(lines []string) int {

	total := 0
	for _, line := range lines {
		total += part1Line(line)
	}
	fmt.Printf("Total for all lines summed: %d", total)

	return total
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		total += part2Line(line)
	}
	fmt.Printf("Total for all lines summed: %d\n", total)

	return total
}

func part2Line(line string) int {
	total := 0

	matchesBetween := findMatchesBetween(line)
	for _, include := range matchesBetween {
		total += part1Line(include)
	}

	return total

}

func main() {
	lines := *util.GetFileAsLines("day3/input.txt")

	fmt.Printf("Part 1 Total %d\n", part1(lines))
	fmt.Printf("Part 2 Total %d\n", part2(lines))

}

func part1Line(line string) int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(line, -1)

	total := 0

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])

		total += x * y
	}

	fmt.Printf("total %d for line: \n", total)

	return total

}

func findMatchesBetween(text string) []string {
	pattern := regexp.MustCompile(fmt.Sprintf("%s(.*?)%s", regexp.QuoteMeta("do()"), regexp.QuoteMeta("don't()")))
	matches := pattern.FindAllStringSubmatch("do()"+text, -1) // cant make stupid regex match start of line.

	var results []string
	for _, match := range matches {
		results = append(results, match[1])
	}
	return results
}
