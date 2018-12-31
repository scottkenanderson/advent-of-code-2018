package main

import "testing"

func Test_firstStar(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "abcdef", args: args{input: []string{
			"Step C must be finished before step A can begin.",
			"Step C must be finished before step F can begin.",
			"Step A must be finished before step B can begin.",
			"Step A must be finished before step D can begin.",
			"Step B must be finished before step E can begin.",
			"Step D must be finished before step E can begin.",
			"Step F must be finished before step E can begin.",
		}}, want: "CABDFE"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstStar(tt.args.input); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secondStar(t *testing.T) {
	type args struct {
		schedule      []string
		offset, elves int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "abcdef", args: args{schedule: []string{
			"Step C must be finished before step A can begin.",
			"Step C must be finished before step F can begin.",
			"Step A must be finished before step B can begin.",
			"Step A must be finished before step D can begin.",
			"Step B must be finished before step E can begin.",
			"Step D must be finished before step E can begin.",
			"Step F must be finished before step E can begin.",
		}, offset: 0, elves: 1}, want: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondStar(tt.args.schedule, tt.args.offset, tt.args.elves); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
