package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fields := flag.String("f", "", "Select fields (columns)")
	delimiter := flag.String("d", "\t", "Specify a field delimiter")
	separated := flag.Bool("s", false, "Only output lines containing delimiter")

	flag.Parse()

	selectedFields, err := parseFields(*fields)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing : %v\n", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Проверка наличия разделителя в строке
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		cut(line, selectedFields, *delimiter)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		return
	}
}

func cut(line string, selectedFields []int, delimiter string) {
	// Разбивка строки на поля
	fields := strings.Split(line, delimiter)

	// Если не указаны поля, выводим всю строку
	if selectedFields[0] == 0 {
		fmt.Println(line)
		return
	}

	// Вывод выбранных полей
	for _, fieldNumber := range selectedFields {
		if fieldNumber > 0 && fieldNumber <= len(fields) {
			fmt.Print(fields[fieldNumber-1])
		}
	}

	fmt.Println()
}

func parseFields(fieldsStr string) ([]int, error) {
	var fields []int
	fieldsStr = strings.TrimSpace(fieldsStr)
	if fieldsStr == "-d" {
		return nil, errors.New("either enter value for -f, or don't use the flag")
	}

	fieldsSlice := strings.Split(fieldsStr, ",")
	for _, fieldStr := range fieldsSlice {
		fieldNum, err := strconv.Atoi(fieldStr)
		if err != nil {
			return nil, fmt.Errorf("error converting %s to integer: %v", fieldStr, err)
		}
		fields = append(fields, fieldNum)
	}

	return fields, nil
}
