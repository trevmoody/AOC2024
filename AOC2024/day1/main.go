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

	// count frequency
	counts := make(map[int]int)
	for _, num := range list2 {
		counts[num]++
	}

	similarityScore := 0
	for _, num := range list1 {
		similarityScore += num * counts[num]
	}

	return similarityScore
}

func part1(list1 []int, list2 []int) int {

	slices.Sort(list1)
	slices.Sort(list2)

	sumOfDiff := 0

	for i := 0; i < len(list1); i++ {
		sumOfDiff += int(math.Abs(float64(list1[i] - list2[i])))
	}

	return sumOfDiff
}
func main() {
	list1, list2 := readInput()

	result1 := part1(list1, list2)
	fmt.Printf("Part 1 Total %d\n", result1)

	result2 := part2(list1, list2)
	fmt.Printf("Part 2 Score %d\n", result2)
}
