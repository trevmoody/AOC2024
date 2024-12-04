package main

import (
	"AOC2024/util"
	"fmt"
)

func part2(data [][]string) int {
	total := 0

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			if data[i][j] == "A" &&
				i+1 < len(data) &&
				j+1 < len(data[i]) &&
				data[i+1][j+1] == "S" && //down forward
				i+1 < len(data) &&
				data[i+1][j-1] == "M" && // down back
				j-1 >= 0 &&
				data[i-1][j-1] == "M" && // up back
				i-1 >= 0 &&
				j+1 < len(data[i]) &&
				data[i-1][j+1] == "S" { // up forward
				total += 1
			}

			if data[i][j] == "A" &&
				i+1 < len(data) &&
				j+1 < len(data[i]) &&
				data[i+1][j+1] == "S" && //down forward
				i+1 < len(data) &&
				data[i+1][j-1] == "S" && // down back
				j-1 >= 0 &&
				data[i-1][j-1] == "M" && // up back
				i-1 >= 0 &&
				j+1 < len(data[i]) &&
				data[i-1][j+1] == "M" { // up forward
				total += 1
			}
			if data[i][j] == "A" &&
				i+1 < len(data) &&
				j+1 < len(data[i]) &&
				data[i+1][j+1] == "M" && //down forward
				i+1 < len(data) &&
				data[i+1][j-1] == "S" && // down back
				j-1 >= 0 &&
				data[i-1][j-1] == "S" && // up back
				i-1 >= 0 &&
				j+1 < len(data[i]) &&
				data[i-1][j+1] == "M" { // up forward
				total += 1
			}
			if data[i][j] == "A" &&
				i+1 < len(data) &&
				j+1 < len(data[i]) &&
				data[i+1][j+1] == "M" && //down forward
				i+1 < len(data) &&
				data[i+1][j-1] == "M" && // down back
				j-1 >= 0 &&
				data[i-1][j-1] == "S" && // up back
				i-1 >= 0 &&
				j+1 < len(data[i]) &&
				data[i-1][j+1] == "S" { // up forward
				total += 1
			}
		}
	}

	return total
}

func part1(data [][]string) int {
	total := 0

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {

			if data[i][j] == "X" {
				// CHECKMASFORWARDS
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

	data := convertToCharSlices(lines)

	//for i, runes := range data {
	//	fmt.Printf("String %d: %v\n", i, runes)
	//}

	fmt.Printf("Part1 count: %d", part1(data))
	fmt.Printf("Part2 count: %d", part2(data))

}

func convertToCharSlices(strings []string) [][]string {
	var charSlices [][]string
	for _, str := range strings {
		var chars []string
		for _, ch := range str {
			chars = append(chars, string(ch))
		}
		charSlices = append(charSlices, chars)
	}
	return charSlices
}
