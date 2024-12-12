package main

import (
	"AOC2024/util"
	"fmt"
	"regexp"
	"strconv"
)

type data struct {
	a1 int
	a2 int
	b1 int
	b2 int
	c1 int
	c2 int
}

func main() {
	fmt.Printf("\nPart 1 testdata 1 cost %d\n", part1(*util.GetFileAsLines("day13/testInput1"), 0))
	fmt.Printf("\nPart 1 real data cost %d\n", part1(*util.GetFileAsLines("day13/input.txt"), 0))
	fmt.Printf("\nPart 2 testdata 1 cost %d\n", part1(*util.GetFileAsLines("day13/testInput1"), 10000000000000))
	fmt.Printf("\nPart 2 real data cost %d\n", part1(*util.GetFileAsLines("day13/input.txt"), 10000000000000))
}

func part1(strings []string, offset int) int {
	buttonRe := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`)
	prizeRe := regexp.MustCompile(`X=(\d+), Y=(\d+)`)

	inputData := []data{}

	for i := 0; i < len(strings); i += 4 {
		aMatches := buttonRe.FindStringSubmatch(strings[i])
		bMatches := buttonRe.FindStringSubmatch(strings[i+1])
		pMatches := prizeRe.FindStringSubmatch(strings[i+2])

		inputData = append(inputData, data{
			a1: toInt(aMatches[1]),
			a2: toInt(aMatches[2]),
			b1: toInt(bMatches[1]),
			b2: toInt(bMatches[2]),
			c1: toInt(pMatches[1]) + offset,
			c2: toInt(pMatches[2]) + offset,
		})

	}

	total := 0
	for _, dataRow := range inputData {
		total += getCost(dataRow.a1, dataRow.b1, dataRow.c1, dataRow.a2, dataRow.b2, dataRow.c2)
	}
	return total
}

func toInt(intStr string) int {
	result, _ := strconv.Atoi(intStr)
	return result
}

func getCost(a1 int, b1 int, c1 int, a2 int, b2 int, c2 int) int {
	////
	//Button A: X+94, Y+34
	//Button B: X+22, Y+67
	//Prize: X=8400, Y=5400
	//
	////cramers rule for 2 equations
	//
	//ButtonA*94 + ButtonB*22 = 8400
	//ButtonA*34 + ButtonB *67 = 5400
	//
	////https://www.chilimath.com/lessons/advanced-algebra/cramers-rule-with-two-variables/
	//
	//a1 * x + b1 * y = c1
	//a2 * x + b2 * y = c2
	//
	//// where x and y are the button pressess

	result := 0

	//Coef Matrix
	D := (a1 * b2) - (b1 * a2)
	//X Matrix
	Dx := (c1 * b2) - (b1 * c2)
	//Y Matrix
	Dy := (a1 * c2) - (c1 * a2)

	// Calculate the cross product results

	if Dx%D == 0 && Dy%D == 0 {
		x := Dx / D
		y := Dy / D

		if x >= 0 && y >= 0 {
			result += x*3 + y
		}
	}

	return result
}
