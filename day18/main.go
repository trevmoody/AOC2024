package main

import (
	"AOC2024/util"
	"container/heap"
	"fmt"
	"image"
	"strconv"
	"strings"
)

type Item struct {
	point    image.Point
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
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
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

var directions = []image.Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func main() {
	fmt.Printf("\nPart 1 testdata 1 distance, %d\n", part1(*util.GetFileAsLines("day18/testInput1"), 6, 12))
	fmt.Printf("\nPart 1 min distance, %d\n", part1(*util.GetFileAsLines("day18/input.txt"), 70, 1024))

	fmt.Printf("\nPart 2 testdata 1 distance, %s\n", part2(*util.GetFileAsLines("day18/testInput1"), 6, 12))
	fmt.Printf("\nPart 2 coord %s\n", part2(*util.GetFileAsLines("day18/input.txt"), 70, 1024))
}

func part1(inputs []string, length int, take int) int {
	grid := buildSquareGrid(length)

	for i := 0; i < take; i++ {
		split := strings.Split(inputs[i], ",")
		i, _ := strconv.Atoi(strings.TrimSpace(split[0]))
		j, _ := strconv.Atoi(strings.TrimSpace(split[1]))
		grid[j][i] = "#"
	}

	start := image.Point{0, 0}
	end := image.Point{length, length}

	shortestPath := dijkstra(grid, start, end)

	return shortestPath
}

func part2(inputs []string, length int, take int) string {
	grid := buildSquareGrid(length)

	start := image.Point{0, 0}
	end := image.Point{length, length}

	shortestPath := 0

	for shortestPath != -1 {
		fmt.Printf("TAKe %d\n", take)
		for i := 0; i < take; i++ {
			split := strings.Split(inputs[i], ",")
			x, _ := strconv.Atoi(strings.TrimSpace(split[0]))
			y, _ := strconv.Atoi(strings.TrimSpace(split[1]))
			grid[y][x] = "#"
		}

		shortestPath = dijkstra(grid, start, end)
		take++

	}

	return inputs[take-2]
}

func dijkstra(grid [][]string, start, end image.Point) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{point: start, priority: 0})

	distances := map[image.Point]int{start: 0}
	visited := map[image.Point]bool{}

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)
		if current.point == end {
			return current.priority
		}
		if visited[current.point] {
			continue
		}
		visited[current.point] = true

		for _, direction := range directions {
			nextPoint := current.point.Add(direction)
			if nextPoint.X < 0 || nextPoint.X >= len(grid) || nextPoint.Y < 0 || nextPoint.Y >= len(grid[0]) {
				continue
			}
			if grid[nextPoint.Y][nextPoint.X] == "#" {
				continue
			}
			newDist := current.priority + 1
			if oldDist, ok := distances[nextPoint]; !ok || newDist < oldDist {
				distances[nextPoint] = newDist
				heap.Push(&pq, &Item{point: nextPoint, priority: newDist})
			}
		}
	}

	return -1 // return -1 if there is no path
}

func printGrid(grid [][]string) {
	for i := range grid {
		fmt.Println(strings.Join(grid[i], ""))
	}
}

func buildSquareGrid(length int) [][]string {
	grid := make([][]string, length+1)
	for i := range grid {
		grid[i] = make([]string, length+1)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	return grid
}
