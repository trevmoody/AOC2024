package main

import (
	"AOC2024/util"
	"container/heap"
	"fmt"
	"math"
	_ "unicode/utf8"
)

func main() {
	fmt.Printf("\nPart 1 testdata 1 cost %d\n", part1(*util.GetFileAsLines("day16/testInput1")))
	fmt.Printf("\nPart 1 testdata 2 cost %d\n", part1(*util.GetFileAsLines("day16/testInput2")))
	fmt.Printf("\nPart 1 cost %d\n", part1(*util.GetFileAsLines("day16/input.txt")))
	//fmt.Printf("\nPart 1 cost %d\n", part1(*util.GetFileAsLines("day16/input.txt")))
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

type Point struct {
	X         int
	Y         int
	direction Direction
}

type PointData struct {
	point             Point
	distanceFromStart int
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

func part1(strings []string) int {
	points, startPoint, endPoint1, endpoint2 := parse(strings)
	fmt.Printf("Start point %v\n", startPoint)
	fmt.Printf("End point %v\n", endPoint1)

	sptSet := make(map[Point]bool)
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
		return points[endPoint1].distanceFromStart
	} else {
		return points[endpoint2].distanceFromStart
	}
}

func processClosestPoint(allPoints map[Point]PointData, sptSet map[Point]bool, pq *PriorityQueue) {
	closestPointData := heap.Pop(pq).(*Item).value
	sptSet[closestPointData.point] = true
	checkNeighboursAndUpdateDistance(allPoints, closestPointData, pq)
}

var directionsToTry = []Turn{FORWARD, LEFT, RIGHT}

func checkNeighboursAndUpdateDistance(points map[Point]PointData, currentPoint PointData, pq *PriorityQueue) {
	for _, turn := range directionsToTry {
		nextPoint, distance := getNextPoint(currentPoint, turn)
		_, ok := points[nextPoint]
		if ok {
			newDistance := currentPoint.distanceFromStart + distance
			if newDistance <= points[nextPoint].distanceFromStart {
				points[nextPoint] = PointData{nextPoint, newDistance}
				heap.Push(pq, &Item{
					value:    points[nextPoint],
					priority: newDistance,
				})
			}
		}
	}
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
				startPoint = Point{j, i, E}
				points[startPoint] = PointData{startPoint, 0}
				points[Point{j, i, N}] = PointData{Point{j, i, N}, math.MaxInt64}
			case "E":
				endPoint1 = Point{j, i, N}
				endPoint2 = Point{j, i, E}
				points[endPoint1] = PointData{endPoint1, math.MaxInt64}
				points[endPoint2] = PointData{endPoint2, math.MaxInt64}
			case ".":
				point1 := Point{j, i, N}
				point2 := Point{j, i, S}
				point3 := Point{j, i, E}
				point4 := Point{j, i, W}
				points[point1] = PointData{point1, math.MaxInt64}
				points[point2] = PointData{point2, math.MaxInt64}
				points[point3] = PointData{point3, math.MaxInt64}
				points[point4] = PointData{point4, math.MaxInt64}
			}
		}
	}

	return points, startPoint, endPoint1, endPoint2
}

func getNextPoint(startPoint PointData, turn Turn) (Point, int) {
	switch {
	case startPoint.point.direction == N && turn == FORWARD:
		return Point{startPoint.point.X, startPoint.point.Y - 1, startPoint.point.direction}, 1
	case startPoint.point.direction == N && turn == LEFT:
		return Point{startPoint.point.X, startPoint.point.Y, W}, 1000
	case startPoint.point.direction == N && turn == RIGHT:
		return Point{startPoint.point.X, startPoint.point.Y, E}, 1000

	case startPoint.point.direction == S && turn == FORWARD:
		return Point{startPoint.point.X, startPoint.point.Y + 1, startPoint.point.direction}, 1
	case startPoint.point.direction == S && turn == LEFT:
		return Point{startPoint.point.X, startPoint.point.Y, E}, 1000
	case startPoint.point.direction == S && turn == RIGHT:
		return Point{startPoint.point.X, startPoint.point.Y, W}, 1000

	case startPoint.point.direction == E && turn == FORWARD:
		return Point{startPoint.point.X + 1, startPoint.point.Y, startPoint.point.direction}, 1
	case startPoint.point.direction == E && turn == LEFT:
		return Point{startPoint.point.X, startPoint.point.Y, N}, 1000
	case startPoint.point.direction == E && turn == RIGHT:
		return Point{startPoint.point.X, startPoint.point.Y, S}, 1000

	case startPoint.point.direction == W && turn == FORWARD:
		return Point{startPoint.point.X - 1, startPoint.point.Y, startPoint.point.direction}, 1

	case startPoint.point.direction == W && turn == LEFT:
		return Point{startPoint.point.X, startPoint.point.Y, S}, 1000
	case startPoint.point.direction == W && turn == RIGHT:
		return Point{startPoint.point.X, startPoint.point.Y, N}, 1000
	}
	panic("This should not happen")

}
