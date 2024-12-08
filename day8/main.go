package main

import (
	"AOC2024/util"
	"fmt"
)

func main() {
	fmt.Printf("Part1 count of antinodes: %d\n", part1(*util.GetFileAsLines("day8/input.txt")))
	//fmt.Printf("Part2 test input count of antinodes: %d\n", part2(*util.GetFileAsLines("day8/testinput2.txt")))
	fmt.Printf("Part2 count of antinodes: %d\n", part2(*util.GetFileAsLines("day8/input.txt")))
}

func part1(lines []string) int {

	maxHorizontal := len(lines[0]) - 1
	maxVertical := len(lines) - 1

	antennaPoints := getAntennaPoints(lines)
	antinodePoints := make(map[util.Point]struct{})

	for _, pointsForAntenna := range antennaPoints {
		for _, point1 := range pointsForAntenna {
			for _, point2 := range pointsForAntenna {
				if point1 == point2 {
					continue
				}
				scale := 2
				newAntinodePoint := util.Point{Vertical: point1.Vertical + scale*(point2.Vertical-point1.Vertical), Horizontal: point1.Horizontal + scale*(point2.Horizontal-point1.Horizontal)}

				if newAntinodePoint.Vertical < 0 ||
					newAntinodePoint.Vertical > maxVertical ||
					newAntinodePoint.Horizontal < 0 ||
					newAntinodePoint.Horizontal > maxHorizontal {
					continue
				}

				antinodePoints[newAntinodePoint] = struct{}{}
			}
		}
	}

	return len(antinodePoints)
}

func part2(lines []string) int {

	maxHorizontal := len(lines[0]) - 1
	maxVertical := len(lines) - 1

	antennaPoints := getAntennaPoints(lines)
	antinodePoints := make(map[util.Point]struct{})

	for _, pointsForAntenna := range antennaPoints {
		for _, point1 := range pointsForAntenna {
			for _, point2 := range pointsForAntenna {
				if point1 == point2 {
					continue
				}
				scale := 1
				for {
					newAntinodePoint := util.Point{Vertical: point1.Vertical + scale*(point2.Vertical-point1.Vertical), Horizontal: point1.Horizontal + scale*(point2.Horizontal-point1.Horizontal)}
					if newAntinodePoint.Vertical < 0 || newAntinodePoint.Vertical > maxVertical || newAntinodePoint.Horizontal < 0 || newAntinodePoint.Horizontal > maxHorizontal {
						break
					}
					antinodePoints[newAntinodePoint] = struct{}{}
					scale += 1
				}
			}
		}
	}

	return len(antinodePoints)
}

func getAntennaPoints(lines []string) map[string][]util.Point {

	points := make(map[string][]util.Point)

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			antenna := string(lines[i][j])
			if antenna != "." {
				currentPoints, ok := points[antenna]
				if ok {
					points[antenna] = append(currentPoints, util.Point{Vertical: i, Horizontal: j})
				} else {
					points[antenna] = []util.Point{{Vertical: i, Horizontal: j}}
				}
			}
		}
	}

	return points

}
