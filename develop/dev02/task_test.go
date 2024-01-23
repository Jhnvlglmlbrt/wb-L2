package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{`qwe\4\5`, `qwe45`, false},
		{`qwe\45`, `qwe44444`, false},
		{`qwe\\5`, `qwe\\\\\`, false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("Unexpected error status, got %v, wantErr: %v", err, tc.wantErr)
				return
			}

			if result != tc.expected {
				t.Errorf("Unexpected result, got %v, want %v", result, tc.expected)
			}
		})
	}
}
