package main

import "testing"

func Test_getCommonLetters(t *testing.T) {
	type args struct {
		boxIds []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "abcdef",
			args: args{
				boxIds: []string{"abcde", "fghij", "klmno", "pqrst", "fguij", "axcye", "wvxyz"},
			},
			want: "fgij",
		},
		{
			name: "abcde",
			args: args{
				boxIds: []string{"abcde", "fghxj", "klmno", "pqrst", "fguij", "abcye", "wvxyz"},
			},
			want: "abce",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCommonLetters(tt.args.boxIds); got != tt.want {
				t.Errorf("getCommonLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
