package main

import (
	"AOC2024/util"
	"fmt"
	"image"
	_ "unicode/utf8"
)

func main() {
	//fmt.Printf("\nPart 1 testdata 1 cost %d\n", part1(*util.GetFileAsLines("day15/testInput1")))
	//fmt.Printf("\nPart 1 testdata 2 cost %d\n", part1(*util.GetFileAsLines("day15/testInput2")))
	//fmt.Printf("\nPart 1 cost %d\n", part1(*util.GetFileAsLines("day15/input.txt")))

	fmt.Printf("\nPart 2 testdata 1 cost %d\n", part2(*util.GetFileAsLines("day15/testInput3")))
	fmt.Printf("\nPart 2 test data 2 cost %d\n", part2(*util.GetFileAsLines("day15/testInput2")))
	fmt.Printf("\nPart 2 test data 2 cost %d\n", part2(*util.GetFileAsLines("day15/input.txt")))
}

func part2(strings []string) int {
	grid, instructions := parse2(strings)

	currentPoint := startPoint(grid, instructions)

	printGrid(grid)

	for i, char := range instructions {

		var move image.Point
		var vertical bool

		fmt.Printf("%d: Moving %s\n", i, string(char))

		switch string(char) {
		case ">":
			move = image.Pt(1, 0)
			vertical = false
		case "v":
			move = image.Pt(0, 1)
			vertical = true
		case "<":
			move = image.Pt(-1, 0)
			vertical = false
		case "^":
			move = image.Pt(0, -1)
			vertical = true
		}
		moved := checkAndMoveIfCan(grid, currentPoint, move, vertical)
		if moved {
			currentPoint = currentPoint.Add(move)
		}

	}
	printGrid(grid)
	return sumBoxes(grid)
}

func checkAndMoveIfCan(grid [][]string, startPoint image.Point, move image.Point, vertical bool) bool {
	if canMove(grid, startPoint, startPoint.Add(move), move, vertical, false) {
		canMove(grid, startPoint, startPoint.Add(move), move, vertical, true)
		return true
	}
	return false
}

func canMove(grid [][]string, currentPoint image.Point, nextPoint image.Point, move image.Point, vertical bool, processMove bool) bool {
	objectAtNextPoint := grid[nextPoint.Y][nextPoint.X]
	objectAtCurrentPoint := grid[currentPoint.Y][currentPoint.X]

	result := false
	switch objectAtNextPoint {
	case ".":
		result = true

	case "#":
		result = false
	case "[":
		if vertical {
			result = canMove(grid, nextPoint, nextPoint.Add(move), move, vertical, processMove) &&
				canMove(grid, nextPoint.Add(image.Pt(1, 0)), nextPoint.Add(move).Add(image.Pt(1, 0)), move, vertical, processMove)
		} else {
			result = canMove(grid, nextPoint, nextPoint.Add(move), move, vertical, processMove)
		}

	case "]":
		if vertical {
			result = canMove(grid, nextPoint, nextPoint.Add(move), move, vertical, processMove) &&
				canMove(grid, nextPoint.Add(image.Pt(-1, 0)), nextPoint.Add(move).Add(image.Pt(-1, 0)), move, vertical, processMove)
		} else {
			result = canMove(grid, nextPoint, nextPoint.Add(move), move, vertical, processMove)
		}

	}
	if result && processMove {
		grid[nextPoint.Y][nextPoint.X] = objectAtCurrentPoint
		grid[currentPoint.Y][currentPoint.X] = "."
	}
	return result
}

func part1(strings []string) int {

	grid, instructions := parse(strings)

	currentPoint := startPoint(grid, instructions)

	printGrid(grid)
	for i, char := range instructions {

		var robotNextPoint image.Point
		var move image.Point

		fmt.Printf("%d: Moving %s\n", i, string(char))

		switch string(char) {
		case ">":
			move = image.Pt(1, 0)
		case "v":
			move = image.Pt(0, 1)
		case "<":
			move = image.Pt(-1, 0)
		case "^":
			move = image.Pt(0, -1)
		}

		robotNextPoint = currentPoint.Add(move)

		dataAtPoint := grid[robotNextPoint.Y][robotNextPoint.X]
		switch dataAtPoint {
		case ".":
			grid[currentPoint.Y][currentPoint.X] = "."
			grid[robotNextPoint.Y][robotNextPoint.X] = "@"
			currentPoint = robotNextPoint
		case "#":
			continue // do nothing
		case "O":
			// whats next to it. need to know first free .
			pointToCheck := robotNextPoint
			canMove := true
			for canMove {
				pointToCheck = pointToCheck.Add(move)
				switch grid[pointToCheck.Y][pointToCheck.X] {
				case ".":
					grid[pointToCheck.Y][pointToCheck.X] = "O"
					grid[robotNextPoint.Y][robotNextPoint.X] = "@"
					grid[currentPoint.Y][currentPoint.X] = "."
					currentPoint = robotNextPoint
					canMove = false
					break
				case "#":
					canMove = false
					break
				case "O":
				}

			}
		default:
			fmt.Printf("This should never happen")

		}
		//		printGrid(grid)

		//reader := bufio.NewReader(os.Stdin)
		//
		//_, _ = reader.ReadString('\n')

	}

	return sumBoxes(grid)
}

func printGrid(grid [][]string) {
	for _, strings := range grid {
		fmt.Printf("%s\n", strings)
	}
	fmt.Printf("\n")
}

func sumBoxes(grid [][]string) int {

	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == "O" || grid[i][j] == "[" {
				total += 100*i + j
			}
		}
	}
	return total
}

func startPoint(grid [][]string, instructions string) image.Point {
	var startPoint image.Point

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == "@" {
				return image.Pt(j, i)
			}
		}
	}

	return startPoint
}

func parse(strings []string) ([][]string, string) {

	var charSlices [][]string
	var instruction string

	for i := 0; i < len(strings); i++ {
		if strings[i] == "" {
			for j := i + 1; j < len(strings); j++ {
				instruction = instruction + strings[j]
			}
			break
		}
		var chars []string
		for _, ch := range strings[i] {
			chars = append(chars, string(ch))
		}
		charSlices = append(charSlices, chars)
	}

	return charSlices, instruction
}

func parse2(strings []string) ([][]string, string) {

	var charSlices [][]string
	var instruction string

	for i := 0; i < len(strings); i++ {
		if strings[i] == "" {
			for j := i + 1; j < len(strings); j++ {
				instruction = instruction + strings[j]
			}
			break
		}
		var chars []string
		for _, ch := range strings[i] {
			switch string(ch) {
			case "#":
				chars = append(chars, "#")
				chars = append(chars, "#")
			case "O":
				chars = append(chars, "[")
				chars = append(chars, "]")
			case ".":
				chars = append(chars, ".")
				chars = append(chars, ".")
			case "@":
				chars = append(chars, "@")
				chars = append(chars, ".")
			}

		}
		charSlices = append(charSlices, chars)
	}

	return charSlices, instruction
}
