package main

import (
	"AOC2024/util"
	"fmt"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

type OrderingRule struct {
	beforePageList []int
	afterPageList  []int
}

func part1(orderingRulesStrings []string, numbersToPrintString []string) int {
	total := 0

	orderingRules := parseOrderingRules(orderingRulesStrings)

	for _, update := range numbersToPrintString {
		updateAsInts := util.StringsToIntsCSV(update)
		if isUpdateInRightOrder(updateAsInts, orderingRules) {
			total += updateAsInts[(len(updateAsInts)-1)/2] // get middle number
		}

	}

	return total
}

func parseOrderingRules(orderingRulesStrings []string) map[int]OrderingRule {
	orderingRules := make(map[int]OrderingRule) //before page to list of after page. // trev maybe set here

	for _, orderingRuleString := range orderingRulesStrings {
		parts := strings.SplitN(orderingRuleString, "|", 2)

		beforePage, _ := strconv.Atoi(parts[0])
		afterPage, _ := strconv.Atoi(parts[1])

		orderingRule, ok := orderingRules[beforePage]
		if ok {
			orderingRules[beforePage] = OrderingRule{orderingRule.beforePageList, append(orderingRule.afterPageList, afterPage)}
		} else {
			orderingRules[beforePage] = OrderingRule{[]int{}, []int{afterPage}}
		}

		orderingRule, ok = orderingRules[afterPage]
		if ok {
			orderingRules[afterPage] = OrderingRule{append(orderingRule.beforePageList, beforePage), orderingRule.afterPageList}
		} else {
			orderingRules[afterPage] = OrderingRule{[]int{beforePage}, []int{}}
		}

	}
	return orderingRules
}

func isUpdateInRightOrder(update []int, orderingRules map[int]OrderingRule) bool {
	// sort the slice, and check if its changed.
	copyUpdate := make([]int, len(update))
	copy(copyUpdate, update)
	fixOrder(copyUpdate, orderingRules)
	return reflect.DeepEqual(update, copyUpdate)
}

func part2(orderingRulesStrings []string, numbersToPrintString []string) int {
	total := 0

	orderingRules := parseOrderingRules(orderingRulesStrings)

	for _, update := range numbersToPrintString {
		updateAsInts := util.StringsToIntsCSV(update)
		if !isUpdateInRightOrder(updateAsInts, orderingRules) {
			fixOrder(updateAsInts, orderingRules)
			total += updateAsInts[(len(updateAsInts)-1)/2] // get middle number
		}
	}
	return total
}

func fixOrder(outOfOrderUpdate []int, orderingRules map[int]OrderingRule) {
	slices.SortFunc(outOfOrderUpdate, func(i, j int) int {
		rule, oki := orderingRules[i]
		if oki {
			if slices.Contains(rule.beforePageList, j) {
				return 1
			}
			return -1
		}
		return 0
	})
}

func main() {
	lines := *util.GetFileAsLines("day5/input.txt")
	orderingRulesStrings, numbersToPrintStrings := util.Split(lines, "")

	fmt.Printf("Part1 count: %d\n", part1(orderingRulesStrings, numbersToPrintStrings))
	fmt.Printf("Part2 count: %d\n", part2(orderingRulesStrings, numbersToPrintStrings))

}
