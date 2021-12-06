package main

import (
	"testing"
)

func Test_simulateFishCount(t *testing.T) {
	type args struct {
		age      int
		daysLeft int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{3, 11}, 2},
		{"edge", args{3, 10}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simulateFishCount(tt.args.age, tt.args.daysLeft); got != tt.want {
				t.Errorf("simulateFishCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveFirst(t *testing.T) {
	type args struct {
		values []int
		days   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1 day", args{readInput("./input_test.txt"), 1}, 5},
		{"5 days", args{readInput("./input_test.txt"), 5}, 10},
		{"10 days", args{readInput("./input_test.txt"), 10}, 12},
		{"18 days", args{readInput("./input_test.txt"), 18}, 26},
		{"80 days", args{readInput("./input_test.txt"), 80}, 5934},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveFirst(tt.args.values, tt.args.days); got != tt.want {
				t.Errorf("solveFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveSecond(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test.txt")}, 26984457539},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSecond(tt.args.values); got != tt.want {
				t.Errorf("solveSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}
