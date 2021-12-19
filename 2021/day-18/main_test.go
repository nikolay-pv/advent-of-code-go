package main

import (
	"testing"
)

func Test_solveFirst(t *testing.T) {
	type args struct {
		nodes []*node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"small", args{nodes: []*node{makeNodes("[[1,2],[[3,4],5]]")}}, 143},
		{"case1", args{nodes: []*node{makeNodes("[[[[0,7],4],[[7,8],[6,0]]],[8,1]]")}}, 1384},
		{"case2", args{nodes: []*node{makeNodes("[[[[1,1],[2,2]],[3,3]],[4,4]]")}}, 445},
		{"case3", args{nodes: []*node{makeNodes("[[[[3,0],[5,3]],[4,4]],[5,5]]")}}, 791},
		{"case4", args{nodes: []*node{makeNodes("[[[[5,0],[7,4]],[5,5]],[6,6]]")}}, 1137},
		{"case5", args{nodes: []*node{makeNodes("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")}}, 3488},
		{"default", args{readInput("./input_test.txt")}, 4140},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveFirst(tt.args.nodes); got != tt.want {
				t.Errorf("solveFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_solveSecond(t *testing.T) {
	type args struct {
		nodes []*node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"default", args{readInput("./input_test.txt")}, 3993},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSecond(tt.args.nodes); got != tt.want {
				t.Errorf("solveSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_magnitude(t *testing.T) {
	type fields struct {
		root *node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "default",
			fields: fields{
				root: &node{
					number: -1,
					left:   &node{9, nil, nil, nil},
					right:  &node{1, nil, nil, nil},
					parent: nil,
				},
			},
			want: 29,
		},
		{
			name:   "nested",
			fields: fields{root: makeNodes("[[9,1],[1,9]]")},
			want:   129,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tt.fields.root
			if got := n.magnitude(); got != tt.want {
				t.Errorf("node.magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reduce(t *testing.T) {
	type args struct {
		parent *node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{parent: makeNodes("[[[[[9,8],1],2],3],4]")},
			want: "[[[[0,9],2],3],4]",
		},
		{
			name: "bugHunt2",
			args: args{parent: makeNodes("[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]")},
			want: "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tt.args.parent
			reduce(n)
			if got := n.print(); got != tt.want {
				t.Errorf("node.print() = %s, want %s", got, tt.want)
			}
		})
	}
}
