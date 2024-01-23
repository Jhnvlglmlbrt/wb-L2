package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		line1, line2 string
		column       int
		numeric      bool
		expected     bool
	}{
		// Сравнение строк по первому столбцу
		{"abc 4 3", "def 3 ff 1", 1, false, true},
		{"def 3 ff 1", "stu 2 ladjh asd", 1, false, true},
		{"stu 2 ladjh asd", "xyz 1 as 2.1", 1, false, true},

		// Сравнение строк по второму столбцу как числа
		{"2.1", "1", 2, true, false},
		{"1.5", "2.1", 2, true, true},
		{"1.5", "1", 2, true, false},
	}

	for _, tt := range tests {
		result := compare(tt.line1, tt.line2, tt.column, tt.numeric)
		fmt.Println(result)
		if result != tt.expected {
			t.Errorf("compare(%q, %q, %d, %v) = %v; want %v", tt.line1, tt.line2, tt.column, tt.numeric, result, tt.expected)
		}
	}
}

func TestReverseLines(t *testing.T) {
	tests := []struct {
		input, expected []string
	}{
		// Обратный порядок строк
		{[]string{"abc 4 3", "def 3 ff 1", "stu 2 ladjh asd", "xyz 1 as 2.1"}, []string{"xyz 1 as 2.1", "stu 2 ladjh asd", "def 3 ff 1", "abc 4 3"}},
		{[]string{"1", "2", "3", "4", "5"}, []string{"5", "4", "3", "2", "1"}},
	}

	for _, tt := range tests {
		lines := make([]string, len(tt.input))
		copy(lines, tt.input)
		reverseLines(lines)
		if !reflect.DeepEqual(lines, tt.expected) {
			t.Errorf("reverseLines(%v) = %v; want %v", tt.input, lines, tt.expected)
		}
	}
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input, expected []string
	}{
		// Удаление дубликатов из строк
		{[]string{"abc 4 3", "def 3 ff 1", "stu 2 ladjh asd", "xyz 1 as 2.1", "abc 4 3"}, []string{"abc 4 3", "def 3 ff 1", "stu 2 ladjh asd", "xyz 1 as 2.1"}},
		{[]string{"1", "2", "3", "4", "5", "1", "2", "3"}, []string{"1", "2", "3", "4", "5"}},
	}

	for _, tt := range tests {
		result := removeDuplicates(tt.input)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("removeDuplicates(%v) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
