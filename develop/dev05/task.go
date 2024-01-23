package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// Объявление флагов
	after := flag.Int("A", 0, "Print N lines after each matching line")
	before := flag.Int("B", 0, "Print N lines before each matching line")
	context := flag.Int("C", 0, "Print ±N lines around each matching line")
	count := flag.Bool("c", false, "Print only the count of matching lines")
	ignoreCase := flag.Bool("i", false, "Case-insensitive matching")
	invert := flag.Bool("v", false, "Invert the sense of matching")
	fixed := flag.Bool("F", false, "Fixed string matching (exact match)")
	lineNum := flag.Bool("n", false, "Print line numbers")

	// Обработка флагов командной строки
	flag.Parse()

	// Получение паттерна для поиска
	pattern := flag.Arg(0)
	if pattern == "" {
		fmt.Println("Usage: grep [OPTIONS] PATTERN [FILE]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Компиляция регулярного выражения, если не используется фиксированный поиск
	var regex *regexp.Regexp
	if !*fixed {
		if *ignoreCase {
			pattern = "(?i)" + pattern
		}
		regex = regexp.MustCompile(pattern)
	}

	// Открытие файла или использование стандартного ввода
	var input *os.File
	if flag.NArg() > 1 {
		fileName := flag.Arg(1)
		var err error
		input, err = os.Open(fileName)
		if err != nil {
			fmt.Printf("Error opening file %s: %s\n", fileName, err)
			os.Exit(1)
		}
		defer input.Close()
	} else {
		input = os.Stdin
	}

	// Чтение входных данных построчно
	scanner := bufio.NewScanner(input)
	lineNumber := 0
	matchingCount := 0

	var beforeBuffer []string
	addToBuffer := true

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		// Проверка на соответствие паттерну
		var fix, match bool
		if *fixed {
			fix = strings.Contains(line, pattern)
		} else {
			match = regex.MatchString(line)
		}

		// Обработка совпадения
		if (*invert && !match) || (fix || match) {
			// Вывод номера строки, если флаг установлен
			if *lineNum && match {
				fmt.Printf("%d: ", lineNumber)
			}

			if !*count {
				fmt.Println(line)
			}
		}

		if match {
			matchingCount++
			addToBuffer = false

			if *after > 0 {
				afterLines(scanner, *after, lineNumber)
			}

			if *context > 0 {
				contextLines(scanner, beforeBuffer, *context, lineNumber)
			}
		}
		// Пока добавление в буфер открыто
		if addToBuffer {
			beforeBuffer = append(beforeBuffer, line)
		}
	}
	// Вывод строк до совпадения
	if *before > 0 {
		beforeLines(beforeBuffer, *before)
	}

	// Вывод количества совпадений
	if *count {
		fmt.Printf("Total matching lines: %d\n", matchingCount)
	}
}

// Вывод N строк после совпадения
func afterLines(scanner *bufio.Scanner, n, lineNumber int) {
	for i := 0; i < n; i++ {
		if scanner.Scan() {
			lineNumber++
			line := scanner.Text()
			fmt.Printf("%d: %s\n", lineNumber, line)
		}
	}
}

// Вывод N строк до совпадения
func beforeLines(buffer []string, n int) {
	len := len(buffer)
	lineNumber := len
	// Вывод нужного количества строк до совпадения
	for i := lineNumber - 1; i >= len-n; i-- {
		fmt.Printf("%d: %s\n", lineNumber, buffer[i])
		lineNumber--
	}
}

// Вывод ±N строк вокруг совпадения
func contextLines(scanner *bufio.Scanner, beforeBuffer []string, n, lineNumber int) {
	beforeLines(beforeBuffer, n)
	afterLines(scanner, n, lineNumber)
}
