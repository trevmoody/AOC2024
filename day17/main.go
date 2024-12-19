package main

import (
	"AOC2024/util"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	_ "unicode/utf8"
)

func main2() {
	//test := part1(*util.GetFileAsLines("day17/testInput1"))
	//fmt.Printf("\nPart 1 testdata 1 cost %d\n", test)

	//test2 := part2(*util.GetFileAsLines("day17/testInput2"))
	//fmt.Printf("\nPart 1 testdata 1 cost %d\n", test2)

	//fmt.Printf("\nPart 2 min A %d\n", part2(*util.GetFileAsLines("day17/testInput2")))

	fmt.Printf("\nPart 2 %d\n", part2(*util.GetFileAsLines("day17/input.txt")))
}

func part2(strings []string) any {
	_, program, ints := parse(strings)
	for i := 0; i < 1000000000000000; i++ {
		_, output := runProgran(Register{i, 0, 0}, program)
		if equalSlices(ints, output) {
			return i
		}
		if i%1000000 == 0 {
			fmt.Printf("Iteration %d\n", i)
		}
	}
	return 0
}

func equalSlices(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func part1(strings []string) int {
	register, program, _ := parse(strings)

	_, output := runProgran(register, program)
	for i, data := range output {
		if i != 0 {
			fmt.Printf(",%d", data)
		} else {
			fmt.Printf("%d", data)
		}
	}

	return 0
}

func runProgran(register Register, program []ProgramPiece) (Register, []int) {

	totalOutput := []int{}

	instructionPointer := 0
	for {
		piece := program[instructionPointer]

		switch piece.instruction {
		case 0:
			denominator := int(math.Pow(float64(2), float64(Operand(register, piece.operand))))
			result := register.A / denominator // check the truncation
			register.A = result
		case 1:
			//operandConverted := Operand(register, piece.operand)
			operandConverted := piece.operand
			result := register.B ^ operandConverted
			register.B = result
		case 2:
			register.B = Operand(register, piece.operand) % 8
		case 3:
			if register.A != 0 {
				instructionPointer = piece.operand / 2
				continue
			}
		case 4:
			register.B = register.B ^ register.C
		case 5:
			output := Operand(register, piece.operand) % 8
			totalOutput = append(totalOutput, output)
			//if first {
			//	fmt.Printf("%d", output)
			//	first = false
			//} else {
			//	fmt.Printf(",%d", output)
			//}
		case 6:
			denominator := int(math.Pow(float64(2), float64(Operand(register, piece.operand))))
			result := register.A / denominator // check the truncation
			register.B = result
		case 7:
			denominator := int(math.Pow(float64(2), float64(Operand(register, piece.operand))))
			result := register.A / denominator // check the truncation
			register.C = result
		}

		instructionPointer++
		if instructionPointer >= len(program) {
			return register, totalOutput
		}
	}
}

func Operand(register Register, operand int) int {
	switch operand {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return register.A
	case 5:
		return register.B
	case 6:
		return register.C
	}
	panic("Invalid operand")

}

type Register struct {
	A int
	B int
	C int
}

type ProgramPiece struct {
	instruction int
	operand     int
}

func toInt(intStr string) int {
	result, _ := strconv.Atoi(intStr)
	return result
}

func parse(lines []string) (Register, []ProgramPiece, []int) {
	re := regexp.MustCompile(`Register [A-Z]: (\d+)`)
	match0 := re.FindStringSubmatch(lines[0])
	match1 := re.FindStringSubmatch(lines[1])
	match2 := re.FindStringSubmatch(lines[2])

	ints := util.StringsToIntsCSV(strings.TrimPrefix(lines[4], "Program: "))

	var program []ProgramPiece

	for i := 0; i < len(ints); i += 2 {
		program = append(program, ProgramPiece{ints[i], ints[i+1]})
	}

	register := Register{toInt(match0[1]), toInt(match1[1]), toInt(match2[1])}

	return register, program, ints

}
