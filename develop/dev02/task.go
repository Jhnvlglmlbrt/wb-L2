package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(s string) (string, error) {
	var b strings.Builder
	isEscapedChar := false

	for i, char := range s {
		if unicode.IsDigit(char) && !isEscapedChar {
			m, err := strconv.Atoi(string(char))
			if err != nil || i == 0 {
				return "", errors.New("некорректная строка")
			}
			r := strings.Repeat(string(s[i-1]), m-1)
			b.WriteString(r)
		} else {
			isEscapedChar = string(char) == `\` && string(s[i-1]) != `\`
			if !isEscapedChar { // if // !false = true - > b.Write...
				b.WriteRune(char)
			}
		}
	}

	return b.String(), nil
}

func main() {
	str := `qwe\\5`
	r, err := Unpack(str)
	if err != nil {
		fmt.Printf("bad unpack for %s: got error %v", str, err)
		return
	}
	fmt.Println(r)
}
