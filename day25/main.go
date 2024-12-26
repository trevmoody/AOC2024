package main

import (
	"AOC2024/util"
	"fmt"
	_ "slices"
)

func main() {
	fmt.Printf("Part 1 test data count, %d, \n", part1(*util.GetFileAsLines("day25/testInput1")))
	fmt.Printf("Part 1 data count, %d, \n", part1(*util.GetFileAsLines("day25/input.txt")))
}

func part1(lines []string) int {
	var keys, locks [][]int

	keys, locks = parse(lines)

	fmt.Printf("locks: %v\n", locks)
	fmt.Printf("keys: %v\n", keys)

	total := 0
	for i := 0; i < len(locks); i++ {
		for k := 0; k < len(keys); k++ {
			if compare(locks[i], keys[k]) {
				total++
			}
		}

	}

	return total
}

func compare(lock []int, key []int) bool {
	for i := 0; i < len(lock); i++ {
		if lock[i]+key[i] > 5 {
			return false
		}
	}
	return true
}

func parse(lines []string) ([][]int, [][]int) {
	keys := make([][]int, 0)
	locks := make([][]int, 0)

	var parsing bool = false
	var lock = false
	var first int = 0
	var last int = 0
	for i, line := range lines {

		if parsing == false && line == "#####" {
			first = i
			lock = true
			parsing = true
			continue
		}

		if parsing == false && line == "....." {
			first = i
			lock = false
			parsing = true
			continue
		}
		if i == len(lines)-1 {
			last = i + 1
			if lock {
				locks = append(locks, parseL(lines, first, last))
			} else {
				//continue
				keys = append(keys, parseK(lines, first, last))
			}
			parsing = false
		}

		if line == "" {
			last = i
			if lock {
				locks = append(locks, parseL(lines, first, last))
			} else {
				//continue
				keys = append(keys, parseK(lines, first, last))
			}
			parsing = false
		}
	}

	return keys, locks

}

func parseL(lines []string, first int, last int) []int {
	var lock []int
	for col := 0; col < len(lines[first]); col++ {
		//work out the height to #
		for row := first + 1; row <= last; row++ {
			if string(lines[row][col]) == "#" {
				continue
			}
			if string(lines[row][col]) == "." {
				lock = append(lock, row-(first+1))
				break
			}
			if row == last {
				panic("trev")
			}
		}

	}
	return lock
}

func parseK(lines []string, first int, last int) []int {
	var key []int
	for col := 0; col < len(lines[first]); col++ {
		//work out the height to #
		for row := last - 1; row >= first; row-- {
			if string(lines[row][col]) == "#" {
				continue
			}
			if string(lines[row][col]) == "." {
				key = append(key, last-1-row-1)
				break
			}
			if row == last {
				panic("trev")
			}
		}

	}
	return key
}
