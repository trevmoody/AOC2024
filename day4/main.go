package main

import (
	"AOC2024/util"
	"fmt"
)

func part2(data [][]string) int {
	// we can skip the outside row / cols as they will not be able to make the X.
	total := 0
	for i := 1; i < len(data)-1; i++ {
		for j := 1; j < len(data[i])-1; j++ {
			if checkMAS(i, j, data) {
				total += 1
			}
		}
	}
	return total
}

func checkMAS(i int, j int, data [][]string) bool {
	return data[i][j] == "A" && // fins the center A, then check the for corners for each combination
		(checkCornersOfX(i, j, data, "S", "M", "M", "S") ||
			checkCornersOfX(i, j, data, "S", "S", "M", "M") ||
			checkCornersOfX(i, j, data, "M", "S", "S", "M") ||
			checkCornersOfX(i, j, data, "M", "M", "S", "S"))

}

func checkCornersOfX(i int, j int, data [][]string, df string, db string, ub string, uf string) bool {
	return data[i+1][j+1] == df && //down forward
		data[i+1][j-1] == db && // down back
		data[i-1][j-1] == ub && // up back
		data[i-1][j+1] == uf // up forward
}

func part1(data [][]string) int {
	total := 0

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {

			if data[i][j] == "X" {
				// forwards
				if j+3 < len(data[i]) && data[i][j+1] == "M" && data[i][j+2] == "A" && data[i][j+3] == "S" {
					total += 1
				}
				// For and down
				if j+3 < len(data[i]) && i+3 < len(data) && data[i+1][j+1] == "M" && data[i+2][j+2] == "A" && data[i+3][j+3] == "S" {
					total += 1
				}
				//dowm
				if i+3 < len(data) && data[i+1][j] == "M" && data[i+2][j] == "A" && data[i+3][j] == "S" {
					total += 1
				}
				//down back
				if j-3 >= 0 && i+3 < len(data) && data[i+1][j-1] == "M" && data[i+2][j-2] == "A" && data[i+3][j-3] == "S" {
					total += 1
				}
				// back
				if j-3 >= 0 && data[i][j-1] == "M" && data[i][j-2] == "A" && data[i][j-3] == "S" {
					total += 1
				}
				// back up
				if i-3 >= 0 && j-3 >= 0 && data[i-1][j-1] == "M" && data[i-2][j-2] == "A" && data[i-3][j-3] == "S" {
					total += 1
				}
				//  up
				if i-3 >= 0 && data[i-1][j] == "M" && data[i-2][j] == "A" && data[i-3][j] == "S" {
					total += 1
				}
				//  up forward
				if i-3 >= 0 && j+3 < len(data[i]) && data[i-1][j+1] == "M" && data[i-2][j+2] == "A" && data[i-3][j+3] == "S" {
					total += 1
				}

			}
		}
	}

	return total
}

func main() {
	lines := *util.GetFileAsLines("day4/input.txt")
	data := util.ConvertToCharSlices(lines)

	//for i, runes := range data {
	//	fmt.Printf("String %d: %v\n", i, runes)
	//}

	fmt.Printf("Part1 count: %d\n", part1(data))
	fmt.Printf("Part2 count: %d\n", part2(data))

}
