package main

import (
	"testing"
)

func Test_solveFirst(t *testing.T) {
	type args struct {
		input Input
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test.txt")}, 198},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveFirst(tt.args.input.values, tt.args.input.width); got != tt.want {
				t.Errorf("solveFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveSecond(t *testing.T) {
	type args struct {
		input Input
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test.txt")}, 230},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSecond(tt.args.input.values, tt.args.input.width); got != tt.want {
				t.Errorf("solveSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}
