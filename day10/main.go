package main

import (
	"AOC2024/util"
	"fmt"
	"image"
)

func main() {
	fmt.Printf("Part1 test2 numPaths: %d\n", part1(*util.GetFileAsLines("day10/testinput.txt")))
	fmt.Printf("Part1 numPaths: %d\n", part1(*util.GetFileAsLines("day10/input.txt")))
}

func part1(lines []string) int {

	grid := util.ConvertToIntSlices(lines)

	return getNumberOfPaths(grid, 0, 9)
}

func getNumberOfPaths(grid [][]int, startHeight int, finishHeight int) int {
	total := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			currentHeight := grid[i][j]
			if currentHeight == startHeight {
				pathsForTrailHead := getNumberOfPathsFromStartPosition(grid, image.Pt(j, i), 9)
				fmt.Printf("Found Trail Head at %d,%d: Paths = %d\n", i, j, pathsForTrailHead)
				total += pathsForTrailHead
			}

		}
	}
	return total
}

func getNumberOfPathsFromStartPosition(grid [][]int, startPosition image.Point, finishHeight int) int {
	directions := []string{"N", "E", "S", "W"}

	total := 0
	found := make(map[image.Point]bool)
	var nextPoint image.Point
	for _, direction := range directions {
		switch direction {
		case "N":
			nextPoint = startPosition.Add(image.Pt(0, -1))
		case "E":
			nextPoint = startPosition.Add(image.Pt(1, 0))
		case "S":
			nextPoint = startPosition.Add(image.Pt(0, 1))
		case "W":
			nextPoint = startPosition.Add(image.Pt(-1, 0))
		}
		total += getNumberOfPathsFromPosition(grid, nextPoint, startPosition, finishHeight, []image.Point{startPosition}, found)
	}
	return total
}

func getNumberOfPathsFromPosition(grid [][]int, currentPoint image.Point, prevPoint image.Point, finishHeight int, pointsVisited []image.Point, found map[image.Point]bool) int {
	if currentPoint.X > len(grid[0])-1 || currentPoint.Y > len(grid)-1 || currentPoint.X < 0 || currentPoint.Y < 0 {
		return 0
	}

	currentHeight := grid[currentPoint.Y][currentPoint.X]
	prevHeight := grid[prevPoint.Y][prevPoint.X]

	if currentHeight != prevHeight+1 {
		return 0
	}
	if currentHeight == finishHeight {
		//
		//fmt.Printf("Found Trail End, Points %v\n", append(pointsVisited, currentPoint))
		found[currentPoint] = true
		return 1
	}
	total := 0
	directions := []string{"N", "E", "S", "W"}

	var nextPoint image.Point
	for _, direction := range directions {
		switch direction {
		case "N":
			nextPoint = currentPoint.Add(image.Pt(0, -1))
		case "E":
			nextPoint = currentPoint.Add(image.Pt(1, 0))
		case "S":
			nextPoint = currentPoint.Add(image.Pt(0, 1))
		case "W":
			nextPoint = currentPoint.Add(image.Pt(-1, 0))
		}
		total += getNumberOfPathsFromPosition(grid, nextPoint, currentPoint, finishHeight, append(pointsVisited, currentPoint), found)
	}
	return total
}
