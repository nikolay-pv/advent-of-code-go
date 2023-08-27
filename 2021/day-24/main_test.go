package main

import (
	"reflect"
	"testing"
)

func Test_compile(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"add", args{"add z w"}, "z += w"},
		{"add", args{"add z 2"}, "z += 2"},
		{"add", args{"add z -20"}, "z += -20"},

		{"mul", args{"mul z w"}, "z *= w"},
		{"mul", args{"mul z 2"}, "z *= 2"},
		{"mul", args{"mul z -20"}, "z *= -20"},

		{"div", args{"div z w"}, "z /= w"},
		{"div", args{"div z 2"}, "z /= 2"},
		{"div", args{"div z -20"}, "z /= -20"},

		{"eql", args{"eql z w"}, "z = btoi(z == w)"},
		{"eql", args{"eql z 2"}, "z = btoi(z == 2)"},
		{"eql", args{"eql z -20"}, "z = btoi(z == -20)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compile(tt.args.line); got != tt.want {
				t.Errorf("compile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decrement(t *testing.T) {
	type args struct {
		input *[14]int
	}
	tests := []struct {
		name string
		args args
		want args
	}{
		{"simple", args{&[14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}}, args{&[14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8}}},
		{"middle", args{&[14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 1}}, args{&[14]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 8, 9}}},
		{"most significant", args{&[14]int{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}, args{&[14]int{1, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decrement(tt.args.input)
			if !reflect.DeepEqual(tt.args.input, tt.want.input) {
				t.Errorf("slice after decrement() = %v, want %v", tt.args.input, tt.want.input)
			}
		})
	}
}
