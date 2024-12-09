package main

import (
	"AOC2024/util"
	"testing"
)

func Test_part1(t *testing.T) {

	testdata := *util.GetFileAsLines("testinput.txt")

	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Part 1", args{testdata}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.lines); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
