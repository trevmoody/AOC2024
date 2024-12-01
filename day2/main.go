package main

import (
	"AOC2024/util"
	"fmt"
)

func main() {
	lines := *util.GetFileAsLines("day2/input.txt")
	part1(lines)
	part2(lines)
}

func part1(lines []string) int {
	safeCount := 0
	for _, line := range lines {
		if getSafeSimple(util.StringsToInts(line)) {
			safeCount += 1
		}
	}
	fmt.Printf("Part 1 SafeCount %d\n", safeCount)
	return safeCount
}

func part2(lines []string) int {
	safeCount := 0
	for _, line := range lines {
		if getSafeDampened(util.StringsToInts(line)) {
			safeCount += 1
		}
	}
	fmt.Printf("Part 2 SafeCount %d\n", safeCount)

	return safeCount

}

func getSafeSimple(data []int) bool {
	fmt.Printf(" Check Data = %v", data)

	if (data[1] - data[0]) > 0 {
		//ascending
		return isSafeWithFunc(data, func(i int) bool { return i > 3 || i < 1 })
	} else {
		//descending
		return isSafeWithFunc(data, func(i int) bool { return i < -3 || i > -1 })
	}
}

// Function to check if the differences are within the safe range using a lambda
func isSafeWithFunc(data []int, notSafeDiffCheck func(int) bool) bool {
	for i := 0; i < len(data)-1; i++ {
		diff := data[i+1] - data[i]
		if notSafeDiffCheck(diff) {
			fmt.Printf(" not safe\n")
			return false
		}
	}
	fmt.Printf(" safe\n")
	return true
}

func getSafeDampened(data []int) bool {
	fmt.Printf("Dampened Raw Data = %v \n", data)
	if getSafeSimple(data) {
		return true
	} else {
		for i := 0; i < len(data); i++ {
			dateWithItemRemoved := util.RemoveAtIndex(data, i)
			if getSafeSimple(dateWithItemRemoved) {
				fmt.Printf("Dampened safe\n")
				return true
			}

		}
	}
	fmt.Printf("Dampened NOT safe\n")
	return false
}
