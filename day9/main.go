package main

import (
	"AOC2024/util"
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Part1 test input checksum: %d\n", part1(*util.GetFileAsLines("day9/testinput.txt")))
	fmt.Printf("Part1 test input 2 checksum : %d\n", part1(*util.GetFileAsLines("day9/testinput2.txt")))
	fmt.Printf("Part1 checksum : %d\n", part1(*util.GetFileAsLines("day9/input.txt")))

	fmt.Printf("Part2 test input 2 checksum : %d\n", part2((*util.GetFileAsLines("day9/testinput2.txt"))[0]))
	fmt.Printf("Part2 input checksum : %d\n", part2((*util.GetFileAsLines("day9/input.txt"))[0]))
}

func part2(line string) int {
	expanded := expand(line)
	//fmt.Printf("Got line:\n %v \n\n", expanded)

	prevId := expanded[len(expanded)-1]
	wordEnd := len(expanded) - 1

	for i := len(expanded) - 2; i >= 0; i-- {
		currentId := expanded[i]
		if currentId == prevId {
			prevId = currentId
			continue
		}
		//boundary..
		if prevId != "." {
			word := expanded[i+1 : wordEnd+1]

			// now to find the gap.

			var prev string
			gapLength := 0
			for j := 0; j < i+len(word); j++ {

				if expanded[j] == "." {
					if prev != "." {
						gapLength = 0
					}
					gapLength += 1
					if gapLength == len(word) {
						// move the word.
						//	fmt.Printf("moving word: %v , i = %d, j = %d \n", word, i, j)

						for k := 0; k < gapLength; k++ {
							expanded[j-gapLength+k+1] = word[k]
							expanded[i+k+1] = "."
						}

						break
					}
				}

				prev = expanded[j]
			}
			//fmt.Printf("no gap, i = %d\n", i)
		}
		wordEnd = i
		prevId = currentId

	}

	return checksum(expanded)
}

func part1(lines []string) int {
	return part1Line(lines[0])
}

func part1Line(line string) int {
	result := expand(line)
	//fmt.Printf("Got line:\n %v \n\n", result)

	firstPotentialEmptySpace := 0
	for i := len(result) - 1; i > 0; i-- {
		switch result[i] {
		case ".":
			continue
		default:
			for j := firstPotentialEmptySpace; j < i; j++ {
				switch result[j] {
				case ".":
					result[j] = result[i]
					result[i] = "."

				default:
					continue
				}

			}
		}
	}

	//fmt.Printf("Got line:\n %v \n\n", result)

	return checksum(result)

}

func expand(line string) []string {
	fileId := 0
	result := make([]string, 0)

	for i, value := range line {
		valueInt, _ := strconv.Atoi(string(value))
		//first is file
		if i%2 == 0 {
			for j := 0; j < valueInt; j++ {
				result = append(result, strconv.Itoa(fileId))
			}
			fileId += 1
		} else {
			for j := 0; j < valueInt; j++ {
				result = append(result, ".")
			}
		}
	}
	return result
}

func checksum(input []string) int {
	result := 0
	for i := 0; i < len(input); i++ {
		if input[i] == "." {
			continue
		}
		valInt, _ := strconv.Atoi(string(input[i]))
		result += valInt * i
	}

	return result
}
