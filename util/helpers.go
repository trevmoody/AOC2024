package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ConvertToCharSlices(strings []string) [][]string {
	var charSlices [][]string
	for _, str := range strings {
		var chars []string
		for _, ch := range str {
			chars = append(chars, string(ch))
		}
		charSlices = append(charSlices, chars)
	}
	return charSlices
}

func GetFileAsLines(fileName string) *[]string {

	currentDir, _ := os.Getwd()
	fmt.Printf("Current DIR: %s\n", currentDir)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic("Error opening file")
		return nil
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	//for i, line := range lines {
	//	fmt.Printf("line %d: %s\n", i, line)
	//}
	//
	//fmt.Printf("line count: %d\n", len(lines))j

	return &lines
}

func SplitSliceOddEvenIndex[T any](slice []T) ([]T, []T) {
	var evenIndices []T
	var oddIndices []T

	for i, value := range slice {
		if i%2 == 0 {
			evenIndices = append(evenIndices, value)
		} else {
			oddIndices = append(oddIndices, value)
		}
	}

	return evenIndices, oddIndices
}

func StringsToInts(line string) []int {
	fields := strings.Fields(line)
	var retList []int
	for _, f := range fields {
		i, _ := strconv.Atoi(strings.TrimSpace(f))
		retList = append(retList, i)
	}
	return retList
}

func StringsToIntsCSV(line string) []int {
	fields := strings.Split(line, ",")
	var retList []int
	for _, f := range fields {
		i, _ := strconv.Atoi(strings.TrimSpace(f))
		retList = append(retList, i)
	}
	return retList
}

func RemoveAtIndex(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	copied := make([]int, len(slice))
	copy(copied, slice)
	newSlice := append(copied[:index], copied[index+1:]...)
	return newSlice
}

func Split(lines []string, splitter string) ([]string, []string) {
	var slice1, slice2 []string
	var split bool
	for _, str := range lines {
		if str == splitter {
			split = true
			continue
		}
		if split {
			slice2 = append(slice2, str)
		} else {
			slice1 = append(slice1, str)
		}
	}
	return slice1, slice2
}

//
//func ReverseInts(s []int) []int {
//	slices.Reverse(s)
//	return s
//
//}
