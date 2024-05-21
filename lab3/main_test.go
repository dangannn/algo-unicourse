package main

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"2 . 5 + 2.2 * 10", "24.5"},
		{"V4", "2"},
		{"2^3", "8"},
		{"2 + (5-V16)^(8*9/5) - 1", "2"},
		{"3 - 1", "2"},
		{"- 1", "-1"},
		{"-8-(66.4+4*(-6*(5+6^(-7.5/8))))", "50.0739841838953"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.input)
		t.Run(testname, func(t *testing.T) {
			ans := Calculate(tt.input)
			if ans != tt.want {
				t.Errorf("got %d, want %s", ans, tt.want)
			}
		})
	}
}
