package main

import (
	"testing"
)

func Test_solveFirst(t *testing.T) {
	type args struct {
		values [2]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test.txt")}, 739785},
		//too low 502130
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
		values [2]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test.txt")}, 444356092776315},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSecond(tt.args.values); got != tt.want {
				t.Errorf("solveSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}
