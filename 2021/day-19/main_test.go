package main

import (
	"testing"
)

func Test_solveFirst(t *testing.T) {
	type args struct {
		values    []scannerData
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"small", args{readInput("./input_small_test.txt"), 3}, 3},
		{"default", args{readInput("./input_test.txt"), 12}, 79},
		{"solution", args{readInput("./input.txt"), 12}, 365},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveFirst(tt.args.values, tt.args.threshold); got != tt.want {
				t.Errorf("solveFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveSecond(t *testing.T) {
	type args struct {
		values []scannerData
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test.txt")}, 3621},
		{"solution", args{readInput("./input.txt")}, 11060},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// make sure the scanners are positioned
			solveFirst(tt.args.values, 12)
			if got := solveSecond(tt.args.values); got != tt.want {
				t.Errorf("solveSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}
