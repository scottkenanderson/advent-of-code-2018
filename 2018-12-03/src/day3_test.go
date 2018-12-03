package main

import "testing"

func Test_findOverlaps(t *testing.T) {
	type args struct {
		claims []claim
		sizeX  int
		sizeY  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "abcdef", args: args{claims: []claim{
			claim{claimID: "#1", x: 1, y: 3, sizeX: 4, sizeY: 4},
			claim{claimID: "#3", x: 3, y: 1, sizeX: 4, sizeY: 4},
			claim{claimID: "#1", x: 5, y: 5, sizeX: 2, sizeY: 2},
		}, sizeX: 8, sizeY: 8}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOverlaps(tt.args.claims, tt.args.sizeX, tt.args.sizeY); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
