package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput() ([]int, []int) {
	file, err := os.Open("day1/input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer file.Close()

	var list1, list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])
			if err1 == nil && err2 == nil {
				list1 = append(list1, num1)
				list2 = append(list2, num2)
			}
		}
	}

	return list1, list2
}

func part2(list1 []int, list2 []int) int {

	counts := countFrequency(list2)

	similarityScore := 0
	for _, num := range list1 {
		similarityScore += num * counts[num]
	}

	return similarityScore
}

func countFrequency(list []int) map[int]int {
	counts := make(map[int]int)
	for _, num := range list {
		counts[num]++
	}
	return counts
}

func part1(list1 []int, list2 []int) int {

	slices.Sort(list1)
	slices.Sort(list2)

	sumOfDiff := 0
	for i, value1 := range list1 {
		sumOfDiff += int(math.Abs(float64(value1 - list2[i])))
	}

	return sumOfDiff
}
func main() {
	list1, list2 := readInput()

	fmt.Printf("Part 1 Total %d\n", part1(list1, list2))
	fmt.Printf("Part 2 Score %d\n", part2(list1, list2))
}
