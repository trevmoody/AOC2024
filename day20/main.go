package main

import (
	"AOC2024/util"
	"fmt"
	"image"
	"math"
	_ "slices"
	"sort"
)

func main() {
	fmt.Printf("Part 1 test data count, %d, \n", day20(*util.GetFileAsLines("day20/testInput1"), 2, 2, true))
	fmt.Printf("Part 1  data count, %d, \n", day20(*util.GetFileAsLines("day20/input.txt"), 2, 100, false))
	fmt.Printf("Part 2 test data count, %d, \n", day20(*util.GetFileAsLines("day20/testInput1"), 20, 50, true))
	fmt.Printf("Part 2  data count, %d, \n", day20(*util.GetFileAsLines("day20/input.txt"), 20, 100, false))
}

func day20(lines []string, maxCheatLength int, timeThreshold int, printSummary bool) int {

	start, end, points := parse(lines)
	fmt.Println(start, end)

	pd, ok := points[end]
	if ok {
		fmt.Printf("end point data distance: %v\n", pd.distance)
	}
	// for each point on the track, we  eed to find any other points that have a
	// manhattan distance is less than the cheat length,
	// and then check if the time saved is greater than the threshold

	cheatMap := make(map[cheat]int)

	for currPoint, currPointData := range points {
		for nextPoint, nextPointData := range points {
			if nextPointData.distance-currPointData.distance > timeThreshold {
				// now check the manhattan distance
				cheatLength :=
					int(math.Abs(float64(currPoint.X)-float64(nextPoint.X))) +
						int(math.Abs(float64(currPoint.Y)-float64(nextPoint.Y)))

				save := nextPointData.distance - currPointData.distance - cheatLength

				if cheatLength <= maxCheatLength && save >= timeThreshold {
					cheatMap[cheat{currPoint, nextPoint}] = save
				}
			}
		}
	}

	if printSummary {
		printSummaryToConsole(cheatMap)
	}

	return len(cheatMap)
}

func printSummaryToConsole(cheatMap map[cheat]int) {
	valueCount := make(map[int]int)
	for _, value := range cheatMap {
		valueCount[value]++
	}

	values := make([]int, 0, len(valueCount))
	for value := range valueCount {
		values = append(values, value)
	}

	sort.Ints(values)

	for _, value := range values {
		fmt.Printf("Value: %d, Count: %d\n", value, valueCount[value])
	}
}

type pointData struct {
	distance  int
	prevPoint image.Point
	nextPoint image.Point
}

type cheat struct {
	startPoint image.Point
	endPoint   image.Point
}

func parse(lines []string) (image.Point, image.Point, map[image.Point]pointData) {

	var start, end image.Point

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if string(lines[i][j]) == "S" {
				start = image.Point{Y: i, X: j}

			}
			if string(lines[i][j]) == "E" {
				end = image.Point{Y: i, X: j}
			}
		}
	}

	points := make(map[image.Point]pointData)
	processPoint(start, 0, start, points, lines)

	return start, end, points

}

func processPoint(
	currentPoint image.Point,
	currentDist int,
	prevPoint image.Point,
	points map[image.Point]pointData,
	lines []string,
) {
	rect := image.Rect(
		0,
		0,
		len(lines[0]),
		len(lines),
	)
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
		if nextPoint.In(rect) && nextPoint != prevPoint {

			switch string(lines[nextPoint.Y][nextPoint.X]) {
			case "#":
				continue
			case "E":
				points[currentPoint] = pointData{currentDist, prevPoint, nextPoint}
				points[nextPoint] = pointData{currentDist + 1, currentPoint, nextPoint}
				return
			case ".":
				// store current point data
				points[currentPoint] = pointData{currentDist, prevPoint, nextPoint}
				processPoint(nextPoint, currentDist+1, currentPoint, points, lines)
				return
			}
		}
	}
}
