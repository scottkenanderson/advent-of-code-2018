package main

import "testing"

func Test_firstStar(t *testing.T) {
	type args struct {
		polymer string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "abcdef", args: args{polymer: "aa"}, want: 2},
		{name: "abcdef", args: args{polymer: "aA"}, want: 0},
		{name: "abcdef", args: args{polymer: "Aa"}, want: 0},
		{name: "abcdef", args: args{polymer: "abBA"}, want: 0},
		{name: "abcdef", args: args{polymer: "abAB"}, want: 4},
		{name: "abcdef", args: args{polymer: "aAbB"}, want: 0},
		{name: "abcdef", args: args{polymer: "dabAcCaCBAcCcaDA"}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstStar(tt.args.polymer); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secondStar(t *testing.T) {
	type args struct {
		polymer string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "abcdef", args: args{polymer: "aba"}, want: 1},
		{name: "abcdef", args: args{polymer: "aAb"}, want: 0},
		{name: "abcdef", args: args{polymer: "Aa"}, want: 0},
		{name: "abcdef", args: args{polymer: "abcBA"}, want: 0},
		{name: "abcdef", args: args{polymer: "dabAcCaCBAcCcaDA"}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondStar(tt.args.polymer); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
