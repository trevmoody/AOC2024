package main

import (
	"AOC2024/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("\nPart 1 testdata stoneCount: %d\n", part1("125 17", 25))
	fmt.Printf("\nPart 2 testdata stoneCount: %d\n", part2("125 17", 25))
	fmt.Printf("Part 2 Stone Count 25 iterations: %d\n", part2((*util.GetFileAsLines("day11/input.txt"))[0], 25))
	fmt.Printf("Part2 Stone Count 75 iterations : %d\n", part2((*util.GetFileAsLines("day11/input.txt"))[0], 75))
	//fmt.Printf("Part2 Stone Count: %d\n", part1((*util.GetFileAsLines("day11/input.txt"))[0], 75))

}

type MapKey struct {
	stoneFace int
	blinks    int
}

func stoneCount(stoneFace int, blinks int, data map[MapKey]int) int {
	mapKey := MapKey{stoneFace, blinks}
	result, ok := data[mapKey]
	if ok {
		return result
	}

	if blinks == 0 {
		data[mapKey] = 1
		return 1
	}

	inputAsStr := strconv.Itoa(stoneFace)
	if len(inputAsStr)%2 == 0 {

		leftStr := string(inputAsStr[0 : len(inputAsStr)/2])
		leftInt, _ := strconv.Atoi(leftStr)

		rightStr := string(inputAsStr[len(inputAsStr)/2:])
		rightInt, _ := strconv.Atoi(rightStr)

		result = stoneCount(leftInt, blinks-1, data) + stoneCount(rightInt, blinks-1, data)
		data[mapKey] = result
		return result
	}

	if stoneFace == 0 {
		result = stoneCount(1, blinks-1, data)
		data[mapKey] = result
		return result
	}

	result = stoneCount(stoneFace*2024, blinks-1, data)
	data[MapKey{stoneFace: stoneFace, blinks: blinks}] = result
	return result
}

func part2(line string, blinkCount int) int {
	var inputList []int
	for _, str := range strings.Split(line, " ") {
		parsedInt, _ := strconv.Atoi(str)
		inputList = append(inputList, parsedInt)
	}

	count := 0
	for _, stoneFace := range inputList {
		count += stoneCount(stoneFace, blinkCount, make(map[MapKey]int))
	}

	return count
}

func part1(line string, blinkCount int) int {

	var inputList []int
	for _, str := range strings.Split(line, " ") {
		parsedInt, _ := strconv.Atoi(str)
		inputList = append(inputList, parsedInt)
	}

	blink := 0
	for blink < blinkCount {
		blink += 1
		fmt.Printf("blink %d, length %d\n", blink, len(inputList))

		var newList []int
		for i := 0; i < len(inputList); i++ {
			newList = append(newList, process(inputList[i])...)
		}
		inputList = newList
		//fmt.Printf("%v", inputList)
	}

	return len(inputList)
}

func process(input int) []int {
	if input == 0 {
		return []int{1}
	}

	inputAsStr := strconv.Itoa(input)
	if len(inputAsStr)%2 == 0 {

		leftStr := string(inputAsStr[0 : len(inputAsStr)/2])
		leftInt, _ := strconv.Atoi(leftStr)

		rightStr := string(inputAsStr[len(inputAsStr)/2:])
		rightInt, _ := strconv.Atoi(rightStr)

		return []int{leftInt, rightInt}
	}
	return []int{input * 2024}
}
