package main

import (
	"AOC2024/util"
	"fmt"
	"strings"
	"time"
)

type PointDirection struct {
	point     util.Point
	direction string
}

func main() {

	start := time.Now()

	result1 := part1(*util.GetFileAsLines("day6/input.txt"))
	duration := time.Since(start)

	start = time.Now()
	result2 := part2(*util.GetFileAsLines("day6/input.txt"))
	duration2 := time.Since(start)

	fmt.Printf("Running Part 1, Time %v: Visited Count: %d \n", duration, result1)
	fmt.Printf("Running Part 2,  Time %v: Positions with Loop Count: %d \n", duration2, result2)

}

func part2(lines []string) int {
	grid, startHorizontal, startVertical, startDirection := parseGrid(lines)
	//fmt.Printf("StartX %d, Start Y %d, Grid: %v\n", startHorizontal, startVertical, grid)

	grid[startVertical][startHorizontal] = "."

	loopCount := 0

	for vertical := 0; vertical < len(grid); vertical++ {
		fmt.Printf("Checking vertical %d\n", vertical)
		for horizontal := 0; horizontal < len(grid[0]); horizontal++ {
			if grid[vertical][horizontal] == "." {
				copied := copy2DSlice(grid) // for each place, add a piece and the see it its a loop
				copied[vertical][horizontal] = "#"
				if checkForLoopInPath(copied, startHorizontal, startVertical, startDirection) {
					loopCount += 1
				}
			}

		}
	}

	return loopCount
}

func checkForLoopInPath(grid [][]string, horizontal int, vertical int, direction string) bool {
	pointDirectionsSeen := make(map[PointDirection]bool)
	pd := PointDirection{point: util.Point{Horizontal: horizontal, Vertical: vertical}, direction: direction}
	pointDirectionsSeen[pd] = true

	return moveAndCheck(grid, pd, pointDirectionsSeen)

}

func moveAndCheck(grid [][]string, pd PointDirection, seen map[PointDirection]bool) bool {

	nextPoint := pd.point.Move(pd.direction)

	if nextPoint.Horizontal < 0 || nextPoint.Horizontal > len(grid[0])-1 || nextPoint.Vertical < 0 || nextPoint.Vertical > len(grid)-1 {
		return false // exit the area, not in a loop
	}

	nextPointValue := grid[nextPoint.Vertical][nextPoint.Horizontal]

	switch nextPointValue {

	case ".":
		{
			newPD := PointDirection{point: nextPoint, direction: pd.direction}
			_, ok := seen[newPD]
			if ok {
				return true
			}
			seen[newPD] = true

			return moveAndCheck(grid, newPD, seen)

		}

	case "#":
		{
			return moveAndCheck(grid, PointDirection{pd.point, getNextDirection(pd.direction)}, seen)
		}

	default:
		panic("aaargh somehting wrong")
	}

}

func part1(lines []string) int {

	grid, startHorizontal, startVertical, startDirection := parseGrid(lines)
	fmt.Printf("StartX %d, Start Y %d, Grid: %v\n", startHorizontal, startVertical, grid)

	grid[startVertical][startHorizontal] = "."
	visited := make(map[util.Point]bool)

	moveAndIncrement(grid, 1, util.Point{Horizontal: startHorizontal, Vertical: startVertical}, startDirection, visited)

	return len(visited)

}

func parseGrid(lines []string) ([][]string, int, int, string) {
	var grid [][]string

	var startHorizontal int
	var startVertical int

	for i, line := range lines {

		parsedLine := strings.Split(line, "")
		grid = append(grid, parsedLine)

		index := strings.Index(line, "^")
		if index != -1 {
			startVertical = i
			startHorizontal = index
			continue
		}
		index = strings.Index(line, ">")
		if index != -1 {
			startVertical = i
			startHorizontal = index
			continue
		}
		index = strings.Index(line, "V")
		if index != -1 {
			startVertical = i
			startHorizontal = index
			continue
		}
		index = strings.Index(line, "<")
		if index != -1 {
			startVertical = i
			startHorizontal = index
			continue
		}
	}

	startDirection := grid[startVertical][startHorizontal]
	return grid, startHorizontal, startVertical, startDirection
}

func moveAndIncrement(grid [][]string, currentCount int, currentPoint util.Point, currentDirection string, visited map[util.Point]bool) int {
	visited[currentPoint] = true

	nextPoint := currentPoint.Move(currentDirection)

	if nextPoint.Horizontal < 0 || nextPoint.Horizontal > len(grid[0])-1 || nextPoint.Vertical < 0 || nextPoint.Vertical > len(grid)-1 {
		return currentCount
	}

	nextPointData := grid[nextPoint.Vertical][nextPoint.Horizontal]

	switch nextPointData {

	case ".":
		{
			// move there and increment
			fmt.Printf("Moving to Next Point horizontal %d, vertical %d\n", nextPoint.Horizontal, nextPoint.Vertical)
			return moveAndIncrement(grid, currentCount+1, nextPoint, currentDirection, visited)

		}

	case "#":
		{
			nextDirection := getNextDirection(currentDirection)
			// turn and dont increment
			fmt.Printf("TURNING\n")
			return moveAndIncrement(grid, currentCount, currentPoint, nextDirection, visited)
		}

	default:
		panic("aaargh something wrong")
	}

}

func getNextDirection(currentDirection string) string {
	var nextDirection string
	switch currentDirection {
	case "^":
		nextDirection = ">"
	case ">":
		nextDirection = "V"
	case "V":
		nextDirection = "<"
	case "<":
		nextDirection = "^"

	}
	return nextDirection
}

func copy2DSlice(original [][]string) [][]string {
	// Create a new slice with the same size as the original
	copied := make([][]string, len(original))
	for i := range original {
		// Create a new inner slice and copy the elements from the original inner slice
		copied[i] = make([]string, len(original[i]))
		copy(copied[i], original[i])
	}
	return copied
}
