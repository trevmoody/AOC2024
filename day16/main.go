package main

import (
	"AOC2024/util"
	"fmt"
	"math"
	"sort"
	_ "unicode/utf8"
)

func main() {
	fmt.Printf("\nPart 1 testdata 1 cost %d\n", part1(*util.GetFileAsLines("day16/testInput1")))
	fmt.Printf("\nPart 1 testdata 2 cost %d\n", part1(*util.GetFileAsLines("day16/testInput2")))
	fmt.Printf("\nPart 1 cost %d\n", part1(*util.GetFileAsLines("day16/input.txt")))
	//fmt.Printf("\nPart 1 cost %d\n", part1(*util.GetFileAsLines("day16/input.txt")))
}

type Point struct {
	X         int
	Y         int
	direction string
}

type PointData struct {
	point             Point
	distanceFromStart int
}

func part1(strings []string) int {

	points, startPoint, endPoint1, endpoint2 := parse(strings)
	fmt.Printf("Start point %v\n", startPoint)
	fmt.Printf("End point %v\n", endPoint1)

	sptSet := make(map[Point]bool)

	for len(sptSet) != len(points) {
		processClosestPoint(points, sptSet, endPoint1)

	}

	if points[endPoint1].distanceFromStart < points[endpoint2].distanceFromStart {
		return points[endPoint1].distanceFromStart
	} else {
		return points[endpoint2].distanceFromStart
	}

}

func processClosestPoint(allPoints map[Point]PointData, sptSet map[Point]bool, endPoint Point) {

	closestPointData := findClosestPoint(allPoints, sptSet)
	//fmt.Printf("\nClosest point %v\n", closestPointData)

	sptSet[closestPointData.point] = true
	checkNeighboursAndUpdateDistance(allPoints, closestPointData, endPoint)

}

var directionsToTry = []string{"FORWARD", "Left", "Right"}

func checkNeighboursAndUpdateDistance(points map[Point]PointData, currentPoint PointData, endPoint Point) {
	for _, turn := range directionsToTry {
		nextPoint, distance := getNextPoint(currentPoint, turn)
		//is this a valid point
		_, ok := points[nextPoint]
		if ok {
			if points[nextPoint].point == endPoint {
				fmt.Printf("At Endpoint")
			}
			newDistance := currentPoint.distanceFromStart + distance
			if newDistance <= points[nextPoint].distanceFromStart {
				//fmt.Printf("Distance for point : %v : %d\n", nextPoint, newDistance)
				points[nextPoint] = PointData{nextPoint, newDistance}
			}
		}

	}

}

func findClosestPoint(points map[Point]PointData, sptSet map[Point]bool) PointData {
	var values []PointData
	for _, value := range points {
		_, ok := sptSet[value.point]
		if ok {
			continue
		}

		values = append(values, value)
	}

	sort.Slice(values, func(i, j int) bool {
		if values[i].distanceFromStart == 0 {
			return true
		}
		if values[j].distanceFromStart == 0 {
			return false
		}
		return values[i].distanceFromStart < values[j].distanceFromStart
	})

	return points[values[0].point]
}

func parse(strings []string) (map[Point]PointData, Point, Point, Point) {
	points := make(map[Point]PointData)
	var startPoint Point
	var endPoint1 Point
	var endPoint2 Point

	for i := 0; i < len(strings); i++ {
		for j := 0; j < len(strings[0]); j++ {
			char := strings[i][j]

			switch string(char) {
			case "S":
				startPoint = Point{j, i, "E"}
				points[startPoint] = PointData{startPoint, 0}
				points[Point{j, i, "N"}] = PointData{Point{j, i, "N"}, math.MaxInt64}
			case "E":
				endPoint1 = Point{j, i, "N"}
				endPoint2 = Point{j, i, "E"}
				points[endPoint1] = PointData{endPoint1, math.MaxInt64}
				points[endPoint2] = PointData{endPoint2, math.MaxInt64}
			case ".":
				point1 := Point{j, i, "N"}
				point2 := Point{j, i, "S"}
				point3 := Point{j, i, "E"}
				point4 := Point{j, i, "W"}
				points[point1] = PointData{point1, math.MaxInt64}
				points[point2] = PointData{point2, math.MaxInt64}
				points[point3] = PointData{point3, math.MaxInt64}
				points[point4] = PointData{point4, math.MaxInt64}
			}
		}
	}

	return points, startPoint, endPoint1, endPoint2
}

func getNextPoint(startPoint PointData, turn string) (Point, int) {
	var nextPoint Point
	currentDirection := startPoint.point.direction
	switch {
	case currentDirection == "N" && turn == "FORWARD":
		return Point{startPoint.point.X, startPoint.point.Y - 1, currentDirection}, 1
	case currentDirection == "N" && turn == "Left":
		return Point{startPoint.point.X, startPoint.point.Y, "W"}, 1000
	case currentDirection == "N" && turn == "Right":
		return Point{startPoint.point.X, startPoint.point.Y, "E"}, 1000

	case currentDirection == "S" && turn == "FORWARD":
		return Point{startPoint.point.X, startPoint.point.Y + 1, currentDirection}, 1
	case currentDirection == "S" && turn == "Left":
		return Point{startPoint.point.X, startPoint.point.Y, "E"}, 1000
	case currentDirection == "S" && turn == "Right":
		return Point{startPoint.point.X, startPoint.point.Y, "W"}, 1000

	case currentDirection == "E" && turn == "FORWARD":
		return Point{startPoint.point.X + 1, startPoint.point.Y, currentDirection}, 1
	case currentDirection == "E" && turn == "Left":
		return Point{startPoint.point.X, startPoint.point.Y, "N"}, 1000
	case currentDirection == "E" && turn == "Right":
		return Point{startPoint.point.X, startPoint.point.Y, "S"}, 1000

	case currentDirection == "W" && turn == "FORWARD":
		return Point{startPoint.point.X - 1, startPoint.point.Y, currentDirection}, 1

	case currentDirection == "W" && turn == "Left":
		return Point{startPoint.point.X, startPoint.point.Y, "S"}, 1000
	case currentDirection == "W" && turn == "Right":
		return Point{startPoint.point.X, startPoint.point.Y, "N"}, 1000
	}

	panic("This should not happen")
	return nextPoint, 0
}
