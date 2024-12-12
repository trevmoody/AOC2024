package main

import (
	"AOC2024/util"
	"fmt"
	"image"
)

func main() {
	i, i2 := part1(*util.GetFileAsLines("day12/input.txt"))
	fmt.Printf("\nPart 1 testdata 1 cost %d, Part 2 %d\n", i, i2)
	//fmt.Printf("\nPart 1 testdata 2 cost %d, Part 2 %d\n", part1(*util.GetFileAsLines("day12/testInput2")))
	//fmt.Printf("\nPart 1 testdata 3 cost %d, PArt 2 %d\n", part1(*util.GetFileAsLines("day12/testInput3")))
	//fmt.Printf("\nPart 1 cost %d\n, Part 2 %d\n", part1(*util.GetFileAsLines("day12/input.txt")))
}

func part1(lines []string) (int, int) {

	grid := util.ConvertToCharSlices(lines)
	pointsSeen := make(map[image.Point]bool)
	outsideCorners := make(map[corner]bool)
	totalPart1 := 0
	totalPart2 := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			name, area, perimeter, corners := calcFromPoint(grid, image.Point{
				X: j,
				Y: i,
			}, pointsSeen, outsideCorners)

			if area != 0 {
				fmt.Printf("Area %s, area %d, perimeter %d, corners %d\n", name, area, perimeter, corners)
			}
			totalPart1 += area * perimeter
			totalPart2 += area * corners
		}
	}

	return totalPart1, totalPart2
}

func calcFromPoint(grid [][]string, point image.Point, seen map[image.Point]bool, outsideCorners map[corner]bool) (string, int, int, int) {

	// have we been here before.
	_, ok := seen[point]
	if ok {
		return "", 0, 0, 0
	}
	seen[point] = true

	totalArea := 0
	totalPerimeter := 0
	totalCorners := 0

	directions := []string{"N", "E", "S", "W"}
	edgesFound := map[string]bool{}
	var nextPoint image.Point
	for _, direction := range directions {
		switch direction {
		case "N":
			nextPoint = point.Add(image.Pt(0, -1))
		case "E":
			nextPoint = point.Add(image.Pt(1, 0))
		case "S":
			nextPoint = point.Add(image.Pt(0, 1))
		case "W":
			nextPoint = point.Add(image.Pt(-1, 0))
		}

		// area
		if nextPoint.Y < 0 ||
			nextPoint.Y > len(grid)-1 ||
			nextPoint.X < 0 || nextPoint.X > len(grid[0])-1 ||
			grid[nextPoint.Y][nextPoint.X] != grid[point.Y][point.X] {
			totalPerimeter += 1
			edgesFound[direction] = true
		} else {
			_, nextArea, nextPerimeter, corners := calcFromPoint(grid, nextPoint, seen, outsideCorners)
			totalArea += nextArea
			totalPerimeter += nextPerimeter
			totalCorners += corners
		}
	}

	totalArea += 1
	totalCorners += calcInsideCorners(edgesFound, grid, point, outsideCorners)

	return grid[point.Y][point.X], totalArea, totalPerimeter, totalCorners
}

var cornerDirections = []string{"NE", "NW", "SE", "SW"}

type corner struct {
	point image.Point
}

func calcInsideCorners(
	edges map[string]bool,
	grid [][]string,
	point image.Point,
	outsideCorners map[corner]bool,
) int {
	cornerCount := 0

	for _, direction := range cornerDirections {
		d1 := string(direction[0])
		d2 := string(direction[1])

		_, ok1 := edges[d1]
		_, ok2 := edges[d2]
		if ok1 && ok2 {
			cornerCount += 1
		} else if (ok1 && !ok2) || (!ok1 && ok2) { //NS
			switch direction {
			case "SE":
				if point.Y < len(grid)-1 && point.X < len(grid[0])-1 && grid[point.Y+1][point.X+1] == grid[point.Y][point.X] {
					_, ok := outsideCorners[corner{point}]
					if !ok {
						cornerCount += 1
					}

					outsideCorners[corner{point}] = true

				}
			case "NE":
				if point.Y > 0 && point.X < len(grid[0])-1 && grid[point.Y-1][point.X+1] == grid[point.Y][point.X] {
					_, ok := outsideCorners[corner{image.Point{point.X, point.Y - 1}}]
					if !ok {
						cornerCount += 1
					}

					outsideCorners[corner{image.Point{X: point.X, Y: point.Y - 1}}] = true
				}
			case "SW":
				if point.Y < len(grid)-1 && point.X > 0 && grid[point.Y+1][point.X-1] == grid[point.Y][point.X] {
					_, ok := outsideCorners[corner{image.Point{point.X - 1, point.Y}}]
					if !ok {
						cornerCount += 1
					}

					outsideCorners[corner{image.Point{X: point.X - 1, Y: point.Y}}] = true
				}
			case "NW":
				if point.Y > 0 && point.X > 0 && grid[point.Y-1][point.X-1] == grid[point.Y][point.X] {
					_, ok := outsideCorners[corner{image.Point{point.X - 1, point.Y - 1}}]
					if !ok {
						cornerCount += 1
					}

					outsideCorners[corner{image.Point{X: point.X - 1, Y: point.Y - 1}}] = true
				}

			}
		}
	}
	return cornerCount
}
