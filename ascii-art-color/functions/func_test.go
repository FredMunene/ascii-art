package functions

import (
	"testing"
)

type TestCases struct {
	Name     string
	Input    string
	Expected string
}

func TestCheckBannerName(t *testing.T) {
	testCases := []TestCases{
		{Name: "provided filename is shadow", Input: "shadow", Expected: "shadow.txt"},
		{Name: "provided filename is standard", Input: "standard", Expected: "standard.txt"},
		{Name: "provided filename is thinkertoy", Input: "thinkertoy", Expected: "thinkertoy.txt"},
		{Name: "provided filename is not shadow, standard,thinkertoy", Input: "othername", Expected: ""},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := CheckBannerName(tc.Input)
			if result != tc.Expected {
				t.Errorf("expected %v, got %v", tc.Expected, result)
			}
		})
	}
}

func TestReadInput(t *testing.T) {
	testCases := []TestCases{
		{Name: "alphanumerics", Input: "0123456789abcdefghijklmnopqrstvuvwxyz", Expected: "0123456789abcdefghijklmnopqrstvuvwxyz"},
		// {Name: "provided filename is standard",Input:"standard",Expected:"standard.txt"},
		// {Name: "provided filename is thinkertoy",Input:"thinkertoy",Expected:"thinkertoy.txt"},
		{Name: "foreign language, i.e. Chinese characters", Input: "会意; 會意", Expected: "Provide characters within the ASCII range (26-126)"},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := ReadInput(tc.Input)
			if result != tc.Expected {
				t.Fatalf("expected %v, got %v", tc.Expected, result)
			}
		})
	}
}
