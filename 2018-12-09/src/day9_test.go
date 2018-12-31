package main

import "testing"

func Test_firstStar(t *testing.T) {
	type args struct {
		players, lastMarble int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {name: "abcdef", args: args{players: 9, lastMarble: 25}, want: 32},
		// {name: "abcdef", args: args{players: 10, lastMarble: 1618}, want: 8317},
		// {name: "abcdef", args: args{players: 13, lastMarble: 7999}, want: 146373},
		// {name: "abcdef", args: args{players: 17, lastMarble: 1104}, want: 2764},
		// {name: "abcdef", args: args{players: 21, lastMarble: 6111}, want: 54718},
		// {name: "abcdef", args: args{players: 30, lastMarble: 5807}, want: 37305},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstStar(tt.args.players, tt.args.lastMarble); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secondStar(t *testing.T) {
	type args struct {
		players, lastMarble int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "abcdef", args: args{players: 9, lastMarble: 25}, want: 32},
		{name: "abcdef", args: args{players: 10, lastMarble: 1618}, want: 8317},
		{name: "abcdef", args: args{players: 13, lastMarble: 7999}, want: 146373},
		{name: "abcdef", args: args{players: 17, lastMarble: 1104}, want: 2764},
		{name: "abcdef", args: args{players: 21, lastMarble: 6111}, want: 54718},
		{name: "abcdef", args: args{players: 30, lastMarble: 5807}, want: 37305},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondStar(tt.args.players, tt.args.lastMarble); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
