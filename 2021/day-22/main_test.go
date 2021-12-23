package main

import (
	"testing"
)

func Test_solveFirst(t *testing.T) {
	type args struct {
		values []cube
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test.txt")}, 39},
		{"defaultBig", args{readInput("./input_test2.txt")}, 590784},
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
		values []cube
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"defaultBig", args{readInput("./input_test3.txt")}, 2758514936282235},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSecond(tt.args.values); got != tt.want {
				t.Errorf("solveSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}
