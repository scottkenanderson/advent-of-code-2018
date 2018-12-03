package main

import "testing"

func Test_parseBoxID(t *testing.T) {
	type args struct {
		boxID string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 bool
	}{
		{name: "abcdef", args: args{boxID: "abcdef"}, want: false, want1: false},
		{name: "bababc", args: args{boxID: "bababc"}, want: true, want1: true},
		{name: "abbcde", args: args{boxID: "abbcde"}, want: true, want1: false},
		{name: "abcccd", args: args{boxID: "abcccd"}, want: false, want1: true},
		{name: "aabcdd", args: args{boxID: "aabcdd"}, want: true, want1: false},
		{name: "abcdee", args: args{boxID: "abcdee"}, want: true, want1: false},
		{name: "ababab", args: args{boxID: "ababab"}, want: false, want1: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseBoxID(tt.args.boxID)
			if got != tt.want {
				t.Errorf("parseBoxID() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseBoxID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
