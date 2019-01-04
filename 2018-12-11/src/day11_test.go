package main

import "testing"

func Test_calculatePowerLevel(t *testing.T) {
	type args struct {
		input fuelCell
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "abcdef", args: args{input: fuelCell{
			gridSerialNumber: 8,
			position:         position{x: 3, y: 5},
		}}, want: 4},
		{name: "abcdef", args: args{input: fuelCell{
			gridSerialNumber: 57,
			position:         position{x: 122, y: 79},
		}}, want: -5},
		{name: "abcdef", args: args{input: fuelCell{
			gridSerialNumber: 39,
			position:         position{x: 217, y: 196},
		}}, want: 0},
		{name: "abcdef", args: args{input: fuelCell{
			gridSerialNumber: 71,
			position:         position{x: 101, y: 153},
		}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.input.powerLevel(); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_findSquarePower(t *testing.T) {
	type args struct {
		startPos                   position
		gridSize, gridSerialNumber int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "abcdef", args: args{
			startPos:         position{x: 90, y: 269},
			gridSerialNumber: 18,
			gridSize:         16}, want: 113},
		{name: "abcdef", args: args{
			startPos:         position{x: 232, y: 251},
			gridSerialNumber: 42,
			gridSize:         12}, want: 119},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSquarePower(tt.args.startPos, tt.args.gridSize, tt.args.gridSerialNumber); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstStar(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name  string
		args  args
		want  position
		power int
	}{
		{name: "abcdef", args: args{input: 18}, want: position{x: 33, y: 45}, power: 29},
		{name: "abcdef", args: args{input: 42}, want: position{x: 21, y: 61}, power: 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, power := firstStar(tt.args.input); got != tt.want || power != tt.power {
				t.Errorf("findOverlaps() = %v %v, want %v %v", got, power, tt.want, tt.power)
			}
		})
	}
}

func Test_secondStar(t *testing.T) {
	type args struct {
		input, gridSize int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "abcdef", args: args{input: 18, gridSize: 300}, want: "90,269,16"},
		{name: "abcdef", args: args{input: 42, gridSize: 300}, want: "232,251,12"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondStar(tt.args.input, tt.args.gridSize); got != tt.want {
				t.Errorf("findOverlaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
