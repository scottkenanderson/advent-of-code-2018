package main

import "testing"

func Test_firstStar(t *testing.T) {
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
		}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstStar(tt.args.claims); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := firstStar(tt.args.claims); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secondStar(t *testing.T) {
	type args struct {
		claims []claim
		allClaimIds map[string]bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "abcdef", args: args{claims: []claim{
			claim{claimID: "#1", x: 1, y: 3, sizeX: 4, sizeY: 4},
			claim{claimID: "#2", x: 3, y: 1, sizeX: 4, sizeY: 4},
			claim{claimID: "#3", x: 5, y: 5, sizeX: 2, sizeY: 2},
		}, allClaimIds: map[string]bool{"#1": true, "#2": true, "#3": true}}, want: "#3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondStar(tt.args.claims, tt.args.allClaimIds); got != tt.want {
				t.Errorf("secondStar() = %v, want %v", got, tt.want)
			}
		})
	}
}
