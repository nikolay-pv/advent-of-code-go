package main

import (
	"reflect"
	"testing"
)

func Test_solveFirst(t *testing.T) {
	type args struct {
		in input
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test.txt")}, 35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveFirst(tt.args.in.a, tt.args.in.img); got != tt.want {
				t.Errorf("solveFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveSecond(t *testing.T) {
	type args struct {
		in input
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test.txt")}, 3351},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSecond(tt.args.in.a, tt.args.in.img); got != tt.want {
				t.Errorf("solveSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_point_neighboring(t *testing.T) {
	tests := []struct {
		name string
		p    point
		want []point
	}{
		{
			name: "default",
			p:    [2]int{5, 10},
			want: []point{{4, 9}, {4, 10}, {4, 11}, {5, 9}, {5, 10}, {5, 11}, {6, 9}, {6, 10}, {6, 11}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.neighboring(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("point.neighboring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_image_binary(t *testing.T) {
	type args struct {
		around point
	}
	tests := []struct {
		name string
		i    image
		args args
		want int
	}{
		{
			name: "default",
			i:    readInput("./input_test.txt").img,
			args: args{
				around: point{2, 2},
			},
			want: 34,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.i.binary(tt.args.around); got != tt.want {
				t.Errorf("image.binary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_imgAlgorithm_outputForBinary(t *testing.T) {
	type args struct {
		binary int
	}
	tests := []struct {
		name string
		img  imgAlgorithm
		args args
		want int
	}{
		{
			name: "default",
			img:  readInput("./input_test.txt").a,
			args: args{
				binary: 34,
			},
			want: 1,
		},
		{
			name: "default",
			img:  readInput("./input_test.txt").a,
			args: args{
				binary: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.img.outputForBinary(tt.args.binary); got != tt.want {
				t.Errorf("imgAlgorithm.outputForBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
