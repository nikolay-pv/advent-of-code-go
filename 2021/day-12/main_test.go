package main

import (
	"testing"
)

func Test_solveFirst(t *testing.T) {
	type args struct {
		values Graph
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"small", args{readInput("./input_test.txt")}, 10},
		{"medium", args{readInput("./input_test2.txt")}, 19},
		{"large", args{readInput("./input_test3.txt")}, 226},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveFirst(tt.args.values); got != tt.want {
				t.Errorf("solveFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveSecond(t *testing.T) {
	type args struct {
		values Graph
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"small", args{readInput("./input_test.txt")}, 36},
		{"medium", args{readInput("./input_test2.txt")}, 103},
		{"large", args{readInput("./input_test3.txt")}, 3509},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSecond(tt.args.values); got != tt.want {
				t.Errorf("solveSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}
