package main

import (
	"reflect"
	"testing"
)

func TestMapStringsToNumbers(t *testing.T) {
	tests := []struct {
		name string
		line string
		want []string
	}{
		{
			name: "Test 1",
			line: "1234567890",
			want: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"},
		},
		{
			name: "Test 2",
			line: "onetwothreefourfivesixseveneightninezero",
			want: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"},
		},
		{
			name: "Test 3",
			line: "one67nine",
			want: []string{"1", "6", "7", "9"},
		},
		{
			name: "Test 4",
			line: "two1nine",
			want: []string{"2", "1", "9"},
		},
		{
			name: "Test 5",
			line: "eightwothree",
			want: []string{"8", "2", "3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapStringsToNumbers(tt.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapStringsToNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
