package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Options структура для хранения параметров командной строки
type Options struct {
	FilePath string
	Column   int
	Numeric  bool
	Reverse  bool
	Unique   bool
}

func main() {
	options := parseFlags()

	if options.FilePath == "" {
		fmt.Println("Необходимо указать путь к файлу c несортированными строками (-file)")
		return
	}

	lines, err := readLines(options.FilePath)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	sortLines(lines, options.Column, options.Numeric, options.Reverse, options.Unique)

	err = writeLines(options.FilePath, lines)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
	}
}

func parseFlags() Options {
	var options Options

	flag.StringVar(&options.FilePath, "file", "", "Путь к файлу c несортированными строками")
	flag.IntVar(&options.Column, "k", 0, "Номер колонки для сортировки")
	flag.BoolVar(&options.Numeric, "n", false, "Сортировать по числовому значению")
	flag.BoolVar(&options.Reverse, "r", false, "Сортировать в обратном порядке")
	flag.BoolVar(&options.Unique, "u", false, "He выводить повторяющиеся строки")

	flag.Parse()

	return options
}

func sortLines(lines []string, column int, numeric bool, reverse bool, unique bool) {
	sort.SliceStable(lines, func(i, j int) bool {
		return compare(lines[i], lines[j], column, numeric)
	})

	if reverse {
		reverseLines(lines)
	}

	if unique {
		lines = removeDuplicates(lines)
	}
}

func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func writeLines(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	return writer.Flush()
}

func compare(line1, line2 string, column int, numeric bool) bool {
	fields1 := splitFields(line1)
	fields2 := splitFields(line2)

	if column > 0 && column <= len(fields1) && column <= len(fields2) {
		field1 := fields1[column-1]
		field2 := fields2[column-1]

		if numeric {
			num1, err1 := strconv.ParseFloat(field1, 64)
			num2, err2 := strconv.ParseFloat(field2, 64)

			if err1 == nil && err2 == nil {
				return num1 < num2
			}
		} else {
			return field1 < field2
		}
	}

	return line1 < line2
}

func splitFields(line string) []string {
	return strings.Fields(line)
}

func reverseLines(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

func removeDuplicates(lines []string) []string {
	seen := make(map[string]struct{})
	var result []string

	for _, line := range lines {
		if _, ok := seen[line]; !ok {
			seen[line] = struct{}{}
			result = append(result, line)
		}
	}

	return result
}
