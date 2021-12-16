package main

import (
	"reflect"
	"testing"
)

func Test_solveFirst(t *testing.T) {
	type args struct {
		values string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "default",
			args: args{preprocessMessage("8A004A801A8002F478")},
			want: 16,
		},
		{
			name: "case1",
			args: args{preprocessMessage("620080001611562C8802118E34")},
			want: 12,
		},
		{
			name: "case2",
			args: args{preprocessMessage("C0015000016115A2E0802F182340")},
			want: 23,
		},
		{
			name: "case3",
			args: args{preprocessMessage("A0016C880162017C3686B18A3D4780")},
			want: 31,
		},
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
		values string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "sum",
			args: args{preprocessMessage("C200B40A82")},
			want: 3,
		},
		{
			name: "product",
			args: args{preprocessMessage("04005AC33890")},
			want: 54,
		},
		{
			name: "minimum",
			args: args{preprocessMessage("880086C3E88112")},
			want: 7,
		},
		{
			name: "maximum",
			args: args{preprocessMessage("CE00C43D881120")},
			want: 9,
		},
		{
			name: "maximum",
			args: args{preprocessMessage("D8005AC2A8F0")},
			want: 1,
		},
		{
			name: "greater",
			args: args{preprocessMessage("F600BC2D8F")},
			want: 0,
		},
		{
			name: "equal",
			args: args{preprocessMessage("9C005AC2F8F0")},
			want: 0,
		},
		{
			name: "complex",
			args: args{preprocessMessage("9C0141080250320F1802104A08")},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSecond(tt.args.values); got != tt.want {
				t.Errorf("solveSecond() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preprocessMessage(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name       string
		args       args
		wantBinary string
	}{
		{
			name:       "default",
			args:       args{"D2FE28"},
			wantBinary: "110100101111111000101000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preprocessMessage(tt.args.hex); got != tt.wantBinary {
				t.Errorf("preprocessMessage() = %v, want %v", got, tt.wantBinary)
			}
		})
	}
}

func Test_decodeHeader(t *testing.T) {
	type args struct {
		binary string
	}
	tests := []struct {
		name         string
		args         args
		wantVersion  int
		wantID       int
		wantNewStart offset
	}{
		{
			name:         "default",
			args:         args{preprocessMessage("D2FE28")},
			wantVersion:  6,
			wantID:       4,
			wantNewStart: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := decodeHeader(tt.args.binary, 0)
			if got != tt.wantVersion {
				t.Errorf("decodeHeader() got = %v, want %v", got, tt.wantVersion)
			}
			if got1 != tt.wantID {
				t.Errorf("decodeHeader() got1 = %v, want %v", got1, tt.wantID)
			}
			if got2 != tt.wantNewStart {
				t.Errorf("decodeHeader() got2 = %v, want %v", got2, tt.wantID)
			}
		})
	}
}

func Test_decodeLiteral(t *testing.T) {
	type args struct {
		binary string
		start  offset
	}
	tests := []struct {
		name         string
		args         args
		wantLiteral  int
		wantNewStart offset
	}{
		{
			name:         "default",
			args:         args{"110100101111111000101000", 6},
			wantLiteral:  2021,
			wantNewStart: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := decodeLiteral(tt.args.binary, tt.args.start)
			if got != tt.wantLiteral {
				t.Errorf("decodeLiteral() got = %v, want %v", got, tt.wantLiteral)
			}
			if got1 != tt.wantNewStart {
				t.Errorf("decodeLiteral() got1 = %v, want %v", got1, tt.wantNewStart)
			}
		})
	}
}

func Test_decodeLiteralPacket(t *testing.T) {
	type args struct {
		binary string
		start  offset
	}
	tests := []struct {
		name  string
		args  args
		want  packet
		want1 offset
	}{
		{
			name: "",
			args: args{"110100101111111000101000", 0},
			want: packet{
				version:    6,
				id:         4,
				literal:    2021,
				subpackets: nil,
			},
			want1: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := decodeLiteralPacket(tt.args.binary, tt.args.start)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeLiteralPacket() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("decodeLiteralPacket() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_decodeOperatorPacket(t *testing.T) {
	type args struct {
		binary string
		start  offset
	}
	tests := []struct {
		name  string
		args  args
		want  packet
		want1 offset
	}{
		{
			name: "default_length",
			args: args{preprocessMessage("38006F45291200"), 0},
			want: packet{
				version: 1,
				id:      6,
				literal: 0,
				subpackets: []packet{
					{6, 4, 10, nil},
					{2, 4, 20, nil},
				},
			},
			want1: 49,
		},
		{
			name: "default_size",
			args: args{preprocessMessage("EE00D40C823060"), 0},
			want: packet{
				version: 7,
				id:      3,
				literal: 0,
				subpackets: []packet{
					{2, 4, 1, nil},
					{4, 4, 2, nil},
					{1, 4, 3, nil},
				},
			},
			want1: 51,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := decodeOperatorPacket(tt.args.binary, tt.args.start)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeOperatorPacket() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("decodeOperatorPacket() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
