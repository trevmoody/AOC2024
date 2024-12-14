package main

import (
	"AOC2024/util"
	"fmt"
	"image"
	"regexp"
	"strconv"
	_ "unicode/utf8"
)

func main() {
	fmt.Printf("\nPart 1 testdata 1 cost %d\n", part1(*util.GetFileAsLines("day14/testInput1"), 100, 11, 7))
	fmt.Printf("\nPart 1 data cost %d\n", part1(*util.GetFileAsLines("day14/input.txt"), 100, 101, 103))
	fmt.Printf("\nPart 1 data cost %d\n", part2(*util.GetFileAsLines("day14/input.txt"), 101, 103))
}

func part2(strings []string, width int, height int) int {

	parsedInput := parse(strings)

	steps := 0
	for {
		steps += 1
		positions := map[image.Point]bool{}
		for _, input := range parsedInput {
			startPoint := image.Pt(input[0], input[1])

			velocity := image.Pt(input[2], input[3])
			distance := velocity.Mul(steps)

			endPoint := startPoint.Add(distance)

			collapsedWith := endPoint.X % width
			if collapsedWith < 0 {
				collapsedWith += width
			}

			collapsedHeight := endPoint.Y % height
			if collapsedHeight < 0 {
				collapsedHeight += height
			}
			endPointCollapsed := image.Pt(collapsedWith, collapsedHeight)
			positions[endPointCollapsed] = true

		}
		if len(positions) == len(parsedInput) {

			var grid [][]bool
			grid = make([][]bool, height)
			for y := 0; y < height; y++ {
				grid[y] = make([]bool, width)
			}
			for position, _ := range positions {
				grid[position.Y][position.X] = true
			}

			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					if grid[y][x] == true {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println()
			}

			return steps
		}

		if steps%10 == 0 {
			fmt.Printf("Steps: %d\n", steps)
		}
	}

}

func part1(strings []string, steps int, width int, height int) int {

	parsedInput := parse(strings)

	quadrants := map[string]int{
		"NW": 0,
		"NE": 0,
		"SE": 0,
		"SW": 0,
	}

	for _, input := range parsedInput {
		startPoint := image.Pt(input[0], input[1])

		velocity := image.Pt(input[2], input[3])
		distance := velocity.Mul(steps)

		endPoint := startPoint.Add(distance)

		collapsedWith := endPoint.X % width
		if collapsedWith < 0 {
			collapsedWith += width
		}

		collapsedHeight := endPoint.Y % height
		if collapsedHeight < 0 {
			collapsedHeight += height
		}
		endPointCollapsed := image.Pt(collapsedWith, collapsedHeight)

		quadrant := getQuadrant(endPointCollapsed, width, height)
		fmt.Printf("end %s, endCollapsed %s, quadrant %s \n", endPoint, endPointCollapsed, quadrant)
		quadrants[quadrant] = quadrants[quadrant] + 1

	}

	total := 1
	for q, count := range quadrants {
		if q != "NONE" {
			total *= count
		}
	}
	return total
}

func getQuadrant(point image.Point, width int, height int) string {

	//7 high  0 - 2 W 4 - 6
	//11 wide
	var NorS string
	var EorW string
	if point.Y == (height-1)/2 || point.X == (width-1)/2 {
		return "NONE"
	}

	if point.Y < (height-1)/2 {
		NorS = "N"
	} else {
		NorS = "S"
	}
	if point.X < (width-1)/2 {
		EorW = "W"
	} else {
		EorW = "E"
	}

	return NorS + EorW

}

func parse(strings []string) [][]int {

	var parsedResults [][]int

	re := regexp.MustCompile(`[-]?\d+`)
	for _, input := range strings {
		matches := re.FindAllString(input, -1)
		if len(matches) != 4 {

			fmt.Printf("invalid input format")
		}

		ints := make([]int, 4)
		for i, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				fmt.Printf("invalid integer: %v", err)
			}
			ints[i] = num
		}

		parsedResults = append(parsedResults, ints)

	}

	return parsedResults
}
