package main

import (
	"AOC2024/util"
	"testing"
)

func Test_getSafeSimple(t *testing.T) {
	type args struct {
		data []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"simple 1", args{[]int{7, 6, 4, 2, 1}}, true},
		{"simple 2", args{[]int{1, 2, 7, 8, 9}}, false},
		{"simple 3", args{[]int{9, 7, 6, 2, 1}}, false},
		{"simple 4", args{[]int{1, 3, 2, 4, 5}}, false},
		{"simple 5", args{[]int{8, 6, 4, 4, 1}}, false},
		{"simple 6", args{[]int{1, 3, 6, 7, 9}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSafeSimple(tt.args.data); got != tt.want {
				t.Errorf("getSafeSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_getSafeDampened(t *testing.T) {
	type args struct {
		data []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"simple 1", args{[]int{7, 6, 4, 2, 1}}, true},
		{"simple 2", args{[]int{1, 2, 7, 8, 9}}, false},
		{"simple 3", args{[]int{9, 7, 6, 2, 1}}, false},
		{"simple 4", args{[]int{1, 3, 2, 4, 5}}, true},
		{"simple 5", args{[]int{8, 6, 4, 4, 1}}, true},
		{"simple 6", args{[]int{1, 3, 6, 7, 9}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSafeDampened(tt.args.data); got != tt.want {
				t.Errorf("getSafeSimple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	type args struct {
		lines []string
	}
	lines := *util.GetFileAsLines("testinput.txt")
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Input", args: args{lines}, want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.lines); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_part2(t *testing.T) {
	type args struct {
		lines []string
	}
	lines := *util.GetFileAsLines("testinput.txt")
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Input", args: args{lines}, want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.lines); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
