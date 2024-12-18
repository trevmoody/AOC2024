package main

import (
	"AOC2024/util"
	"container/heap"
	"fmt"
	"math"
	_ "unicode/utf8"
)

func main() {
	i := part1(*util.GetFileAsLines("day16/testInput1"))
	fmt.Printf("\nPart 1 testdata 1 distance, %d seats %d\n", i.distanceFromStart, len(i.pointsSeenOnShortestRoute))

	i2 := part1(*util.GetFileAsLines("day16/testInput2"))
	fmt.Printf("\nPart 1 testdata 2 distance, %d seats %d\n", i2.distanceFromStart, len(i2.pointsSeenOnShortestRoute))

	i3 := part1(*util.GetFileAsLines("day16/input.txt"))
	fmt.Printf("\nPart 1 testdata 2 distance, %d seats %d\n", i3.distanceFromStart, len(i3.pointsSeenOnShortestRoute))
}

type Direction int

const (
	N Direction = iota
	S
	E
	W
)

type Turn int

const (
	FORWARD Turn = iota
	LEFT
	RIGHT
)

var directionsToTry = []Turn{FORWARD, LEFT, RIGHT}

type PointDirection struct {
	X         int
	Y         int
	direction Direction
}
type Point struct {
	X int
	Y int
}

type PointData struct {
	pointDirection            PointDirection
	distanceFromStart         int
	pointsSeenOnShortestRoute map[Point]bool
}

func (p PointDirection) toPoint() Point {
	return Point{p.X, p.Y}
}

type Item struct {
	value    PointData // The value of the item; arbitrary.
	priority int       // The priority of the item in the queue.
	index    int       // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value PointData, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func part1(strings []string) PointData {
	points, startPoint, endPoint1, endpoint2 := parse(strings)
	fmt.Printf("Start pointDirection %v\n", startPoint)
	fmt.Printf("End pointDirection %v\n", endPoint1)

	sptSet := make(map[PointDirection]bool)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for _, pointData := range points {
		heap.Push(&pq, &Item{
			value:    pointData,
			priority: pointData.distanceFromStart,
		})
	}

	for len(sptSet) != len(points) {
		processClosestPoint(points, sptSet, &pq)
	}

	if points[endPoint1].distanceFromStart < points[endpoint2].distanceFromStart {
		return points[endPoint1]
	} else {
		return points[endpoint2]
	}
}

func processClosestPoint(allPoints map[PointDirection]PointData, sptSet map[PointDirection]bool, pq *PriorityQueue) {
	closestPointData := heap.Pop(pq).(*Item).value

	closestPointData.pointsSeenOnShortestRoute[closestPointData.pointDirection.toPoint()] = true

	sptSet[closestPointData.pointDirection] = true
	checkNeighboursAndUpdateDistance(allPoints, closestPointData, pq)
}

func copyMap(original map[Point]bool) map[Point]bool {
	copied := make(map[Point]bool)
	for key, value := range original {
		copied[key] = value
	}
	return copied
}

func checkNeighboursAndUpdateDistance(points map[PointDirection]PointData, currentPoint PointData, pq *PriorityQueue) {
	for _, turn := range directionsToTry {
		nextPoint, distance := getNextPoint(currentPoint, turn)
		_, ok := points[nextPoint]
		if ok {
			newDistance := currentPoint.distanceFromStart + distance
			if newDistance < points[nextPoint].distanceFromStart {
				points[nextPoint] = PointData{nextPoint, newDistance, copyMap(currentPoint.pointsSeenOnShortestRoute)}
				heap.Push(pq, &Item{
					value:    points[nextPoint],
					priority: newDistance,
				})
			} else if newDistance == points[nextPoint].distanceFromStart {
				combinedPoints := combineMaps(points[nextPoint].pointsSeenOnShortestRoute, currentPoint.pointsSeenOnShortestRoute)

				points[nextPoint] = PointData{nextPoint, newDistance, combinedPoints}
				heap.Push(pq, &Item{
					value:    points[nextPoint],
					priority: newDistance,
				})
			}
		}
	}
}

func combineMaps(route map[Point]bool, route2 map[Point]bool) map[Point]bool {
	for key, value := range route2 {
		route[key] = value
	}
	return route

}

func parse(strings []string) (map[PointDirection]PointData, PointDirection, PointDirection, PointDirection) {
	points := make(map[PointDirection]PointData)
	var startPoint PointDirection
	var endPoint1 PointDirection
	var endPoint2 PointDirection

	for i := 0; i < len(strings); i++ {
		for j := 0; j < len(strings[0]); j++ {
			char := strings[i][j]

			switch string(char) {
			case "S":
				startPoint = PointDirection{j, i, E}
				points[startPoint] = PointData{startPoint, 0, map[Point]bool{}}
				points[PointDirection{j, i, N}] = PointData{PointDirection{j, i, N}, math.MaxInt64, map[Point]bool{}}
			case "E":
				endPoint1 = PointDirection{j, i, N}
				endPoint2 = PointDirection{j, i, E}
				points[endPoint1] = PointData{endPoint1, math.MaxInt64, map[Point]bool{}}
				points[endPoint2] = PointData{endPoint2, math.MaxInt64, map[Point]bool{}}
			case ".":
				point1 := PointDirection{j, i, N}
				point2 := PointDirection{j, i, S}
				point3 := PointDirection{j, i, E}
				point4 := PointDirection{j, i, W}
				points[point1] = PointData{point1, math.MaxInt64, map[Point]bool{}}
				points[point2] = PointData{point2, math.MaxInt64, map[Point]bool{}}
				points[point3] = PointData{point3, math.MaxInt64, map[Point]bool{}}
				points[point4] = PointData{point4, math.MaxInt64, map[Point]bool{}}
			}
		}
	}

	return points, startPoint, endPoint1, endPoint2
}

func getNextPoint(startPoint PointData, turn Turn) (PointDirection, int) {
	switch {
	case startPoint.pointDirection.direction == N && turn == FORWARD:
		return PointDirection{startPoint.pointDirection.X, startPoint.pointDirection.Y - 1, startPoint.pointDirection.direction}, 1
	case startPoint.pointDirection.direction == N && turn == LEFT:
		return PointDirection{startPoint.pointDirection.X, startPoint.pointDirection.Y, W}, 1000
	case startPoint.pointDirection.direction == N && turn == RIGHT:
		return PointDirection{startPoint.pointDirection.X, startPoint.pointDirection.Y, E}, 1000

	case startPoint.pointDirection.direction == S && turn == FORWARD:
		return PointDirection{startPoint.pointDirection.X, startPoint.pointDirection.Y + 1, startPoint.pointDirection.direction}, 1
	case startPoint.pointDirection.direction == S && turn == LEFT:
		return PointDirection{startPoint.pointDirection.X, startPoint.pointDirection.Y, E}, 1000
	case startPoint.pointDirection.direction == S && turn == RIGHT:
		return PointDirection{startPoint.pointDirection.X, startPoint.pointDirection.Y, W}, 1000

	case startPoint.pointDirection.direction == E && turn == FORWARD:
		return PointDirection{startPoint.pointDirection.X + 1, startPoint.pointDirection.Y, startPoint.pointDirection.direction}, 1
	case startPoint.pointDirection.direction == E && turn == LEFT:
		return PointDirection{startPoint.pointDirection.X, startPoint.pointDirection.Y, N}, 1000
	case startPoint.pointDirection.direction == E && turn == RIGHT:
		return PointDirection{startPoint.pointDirection.X, startPoint.pointDirection.Y, S}, 1000

	case startPoint.pointDirection.direction == W && turn == FORWARD:
		return PointDirection{startPoint.pointDirection.X - 1, startPoint.pointDirection.Y, startPoint.pointDirection.direction}, 1

	case startPoint.pointDirection.direction == W && turn == LEFT:
		return PointDirection{startPoint.pointDirection.X, startPoint.pointDirection.Y, S}, 1000
	case startPoint.pointDirection.direction == W && turn == RIGHT:
		return PointDirection{startPoint.pointDirection.X, startPoint.pointDirection.Y, N}, 1000
	}
	panic("This should not happen")

}
