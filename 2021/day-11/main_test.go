package main

import (
	"testing"
)

func Test_solveFirst(t *testing.T) {
	type args struct {
		values [][]int
		steps  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default_0", args{readInput("./input_test.txt"), 1}, 0},
		{"default_2", args{readInput("./input_test.txt"), 2}, 35},
		{"default_10", args{readInput("./input_test.txt"), 10}, 204},
		{"default", args{readInput("./input_test.txt"), 100}, 1656},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveFirst(tt.args.values, tt.args.steps); got != tt.want {
				t.Errorf("solveFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveSecond(t *testing.T) {
	type args struct {
		values [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test.txt")}, 195},
		// {"tooLow", args{readInput("./input.txt")}, 248},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSecond(tt.args.values); got != tt.want {
				t.Errorf("solveSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}
