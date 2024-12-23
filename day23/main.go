package main

import (
	"AOC2024/util"
	"fmt"
	_ "slices"
	"sort"
	"strings"
)

func main() {
	fmt.Printf("Part 1 test data count, %d, \n", part1(*util.GetFileAsLines("day23/testInput1")))
	fmt.Printf("Part 1  data count, %d, \n", part1(*util.GetFileAsLines("day23/input.txt")))
	fmt.Printf("Part 2 test data: %s \n", part2(*util.GetFileAsLines("day23/testInput1")))
	fmt.Printf("Part 2 data: %s \n", part2(*util.GetFileAsLines("day23/input.txt")))
	//keypad4 := Keypad{keys: [][]string{{"X", "^", "A"}, {"<", "V", ">"}}}
	//keypad3 := Keypad{keys: [][]string{{"X", "^", "A"}, {"<", "V", ">"}}}

}

type pair struct {
	x, y string
}

type triple struct {
	x, y, z string
}

func createTriple(p pair, z string) triple {
	s := []string{p.x, p.y, z}
	sort.Strings(s)
	return triple{s[0], s[1], s[2]}
}

func createPair(x string, y string) pair {
	if x > y {
		return pair{x, y}
	} else {
		return pair{y, x}
	}
}

func part2(lines []string) string {
	pairs := make(map[pair]bool)
	computers := make(map[string][]bool)

	for _, line := range lines {
		fields := strings.Split(line, "-")

		pairs[createPair(fields[0], fields[1])] = true
		computers[fields[0]] = append(computers[fields[0]], true)
		computers[fields[1]] = append(computers[fields[1]], true)
	}

	groups := map[string][]string{}
	for p, _ := range pairs {
		newGroup := []string{p.x, p.y}
		sort.Strings(newGroup)
		newGroupKey := strings.Join(newGroup, ",")
		groups[newGroupKey] = newGroup
	}
	return findNextGroupSize(groups, computers, pairs)

}

func findNextGroupSize(groups map[string][]string, computers map[string][]bool, pairs map[pair]bool) string {
	newGroups := map[string][]string{}
	for _, group := range groups {
		for compName, _ := range computers {
			found := true
			for _, existingCompName := range group {
				pairToCheck := createPair(existingCompName, compName)
				_, ok := pairs[pairToCheck]
				if !ok {
					found = false
					break
				}
			}
			if found {
				newGroup := append(group, compName)
				sort.Strings(newGroup)
				newGroupKey := strings.Join(newGroup, ",")
				newGroups[newGroupKey] = newGroup
			}

		}
	}
	if len(newGroups) == 1 {
		for key, _ := range newGroups {
			return key
		}

	}
	return findNextGroupSize(newGroups, computers, pairs)

}
func part1(lines []string) int {
	pairs := make(map[pair]bool)
	triples := make(map[triple]bool)
	computers := make(map[string][]bool)

	for _, line := range lines {
		fields := strings.Split(line, "-")
		pairs[createPair(fields[0], fields[1])] = true
		computers[fields[0]] = append(computers[fields[0]], true)
		computers[fields[1]] = append(computers[fields[1]], true)

	}

	for p, _ := range pairs {
		for compName, _ := range computers {
			pair1 := createPair(p.x, compName)
			pair2 := createPair(p.y, compName)

			_, ok1 := pairs[pair1]
			_, ok2 := pairs[pair2]
			if ok1 && ok2 {
				triples[createTriple(p, compName)] = true
			}
		}
	}
	filtered := make(map[triple]bool)
	for t, _ := range triples {
		if string(t.x[0]) == "t" || string(t.y[0]) == "t" || string(t.z[0]) == "t" {
			filtered[t] = true
		}
	}
	return len(filtered)
}
