package main

import (
	"testing"
)

func Test_solveFirst(t *testing.T) {
	type args struct {
		values diagram
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"simplest", args{readInput("./input_test_simplest.txt", 2)}, 4600},
		{"default", args{readInput("./input_test.txt", 2)}, 12521},
		{"solution", args{readInput("./input.txt", 2)}, 14148},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveFirst(tt.args.values); got != tt.want {
				t.Errorf("solveFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

// simple test to be able to run:
// go test -bench=Benchmark_solver -run=X -cpuprofile cpu.prof
// go tool pprof -http :8000 main.test cpu.prof
func Benchmark_solver(b *testing.B) {
	for n := 0; n < b.N; n++ {
		solveFirst(readInput("./input_test_part2.txt", 4))
	}
}

func Test_solveSecond(t *testing.T) {
	type args struct {
		values diagram
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test_part2.txt", 4)}, 44169},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSecond(tt.args.values); got != tt.want {
				t.Errorf("solveSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}
