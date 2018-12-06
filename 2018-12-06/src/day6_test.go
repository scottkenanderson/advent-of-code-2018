package main

import (
	"testing"
)

func Test_firstStar(t *testing.T) {
	type args struct {
		coOrdinates coOrdinates
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{coOrdinates{coOrdinates: []coOrdinate{
			coOrdinate{x: 1, y: 1},
			coOrdinate{x: 1, y: 6},
			coOrdinate{x: 8, y: 3},
			coOrdinate{x: 3, y: 4},
			coOrdinate{x: 5, y: 5},
			coOrdinate{x: 8, y: 9},
		}}}, want: 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstStar(tt.args.coOrdinates); got != tt.want {
				t.Errorf("firstStar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secondStar(t *testing.T) {
	type args struct {
		coOrdinates coOrdinates
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{coOrdinates{coOrdinates: []coOrdinate{
			coOrdinate{x: 1, y: 1},
			coOrdinate{x: 1, y: 6},
			coOrdinate{x: 8, y: 3},
			coOrdinate{x: 3, y: 4},
			coOrdinate{x: 5, y: 5},
			coOrdinate{x: 8, y: 9},
		}}}, want: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondStar(tt.args.coOrdinates, 32); got != tt.want {
				t.Errorf("firstStar() = %v, want %v", got, tt.want)
			}
		})
	}
}
