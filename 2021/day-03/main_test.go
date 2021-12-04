package main

import (
	"testing"
)

func Test_solve_first(t *testing.T) {
	type args struct {
		input Input
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{read_input("./input_test.txt")}, 198},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_first(tt.args.input.values, tt.args.input.width); got != tt.want {
				t.Errorf("solve_first() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solve_second(t *testing.T) {
	type args struct {
		input Input
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{read_input("./input_test.txt")}, 230},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve_second(tt.args.input.values, tt.args.input.width); got != tt.want {
				t.Errorf("solve_second() = %v, want %v", got, tt.want)
			}
		})
	}
}
