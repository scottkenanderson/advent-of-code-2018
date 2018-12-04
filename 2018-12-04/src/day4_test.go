package main

import "testing"

func Test_firstStar(t *testing.T) {
	type args struct {
		schedule []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "abcdef", args: args{schedule: []string{
			"[1518-11-01 00:00] Guard #10 begins shift",
			"[1518-11-01 00:05] falls asleep",
			"[1518-11-01 00:25] wakes up",
			"[1518-11-01 00:30] falls asleep",
			"[1518-11-01 00:55] wakes up",
			"[1518-11-01 23:58] Guard #99 begins shift",
			"[1518-11-02 00:40] falls asleep",
			"[1518-11-02 00:50] wakes up",
			"[1518-11-03 00:05] Guard #10 begins shift",
			"[1518-11-03 00:24] falls asleep",
			"[1518-11-03 00:29] wakes up",
			"[1518-11-04 00:02] Guard #99 begins shift",
			"[1518-11-04 00:36] falls asleep",
			"[1518-11-04 00:46] wakes up",
			"[1518-11-05 00:03] Guard #99 begins shift",
			"[1518-11-05 00:45] falls asleep",
			"[1518-11-05 00:55] wakes up",
		}}, want: 10 * 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstStar(tt.args.schedule); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_secondStar(t *testing.T) {
	type args struct {
		schedule []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "abcdef", args: args{schedule: []string{
			"[1518-11-01 00:00] Guard #10 begins shift",
			"[1518-11-01 00:05] falls asleep",
			"[1518-11-01 00:25] wakes up",
			"[1518-11-01 00:30] falls asleep",
			"[1518-11-01 00:55] wakes up",
			"[1518-11-01 23:58] Guard #99 begins shift",
			"[1518-11-02 00:40] falls asleep",
			"[1518-11-02 00:50] wakes up",
			"[1518-11-03 00:05] Guard #10 begins shift",
			"[1518-11-03 00:24] falls asleep",
			"[1518-11-03 00:29] wakes up",
			"[1518-11-04 00:02] Guard #99 begins shift",
			"[1518-11-04 00:36] falls asleep",
			"[1518-11-04 00:46] wakes up",
			"[1518-11-05 00:03] Guard #99 begins shift",
			"[1518-11-05 00:45] falls asleep",
			"[1518-11-05 00:55] wakes up",
		}}, want: 99 * 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondStar(tt.args.schedule); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
