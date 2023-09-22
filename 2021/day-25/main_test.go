package main

import (
	"testing"
)

func Test_solveFirst(t *testing.T) {
	type args struct {
		values [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test.txt")}, 58},
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
		values [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test.txt")}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSecond(tt.args.values); got != tt.want {
				t.Errorf("solveSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveOnce(t *testing.T) {
	type args struct {
		values [][]int
	}
	tests := []struct {
		name      string
		args      args
		want      args
		wantMoves int
	}{
		{"default", args{readInput("./input_test.txt")}, args{readInput("./input_test_first_step.txt")}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveOnce(tt.args.values); got != tt.wantMoves {
				t.Errorf("moveOnce() = %v, want %v", got, tt.wantMoves)
			}
			rows := len(tt.want.values)
			cols := len(tt.want.values[0])
			for row := 0; row < rows; row++ {
				for col := 0; col < cols; col++ {
					got := tt.args.values[row][col]
					want := tt.want.values[row][col]
					if want != got {
						t.Errorf("values[%v][%v] = %v, want %v", row, col, got, want)

					}
				}
			}
		})
	}
}
