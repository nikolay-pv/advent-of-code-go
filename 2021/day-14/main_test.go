package main

import (
	"testing"
)

func Test_solveFirst(t *testing.T) {
	type args struct {
		in Input
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// {"default", args{readInput("./input_test.txt")}, 1588},
		{"submitted", args{readInput("./input.txt")}, 2321},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveFirst(tt.args.in.formula, tt.args.in.rules); got != tt.want {
				t.Errorf("solveFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveSecond(t *testing.T) {
	type args struct {
		in Input
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"default", args{readInput("./input_test.txt")}, 2188189693529},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSecond(tt.args.in.formula, tt.args.in.rules); got != tt.want {
				t.Errorf("solveSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}
