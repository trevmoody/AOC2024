package main

import (
	"AOC2024/util"
	"testing"
)

func Test_part2(t *testing.T) {
	data := util.ConvertToCharSlices(*util.GetFileAsLines("input.txt"))
	testdata := util.ConvertToCharSlices(*util.GetFileAsLines("testinput2.txt"))

	type args struct {
		data [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Part 2 Test Data", args{testdata}, 9},
		{"Part 2 Real Data", args{data}, 2034},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.data); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {

	lines := *util.GetFileAsLines("input.txt")
	data := util.ConvertToCharSlices(lines)

	type args struct {
		data [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Part 1", args{data}, 2662},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.data); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
