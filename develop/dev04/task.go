/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func unorderedEqual(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	exists := make(map[rune]bool)
	for _, v := range str1 {
		exists[v] = true
	}
	for _, v := range str2 {
		if !exists[v] {
			return false
		}
	}

	return true
}

func Anagram(data []string) map[string][]string {
	res := make(map[string][]string)
	exists := make(map[string]bool)

	len := len(data)
	for i, v := range data {
		exists[v] = true
		for j := i + 1; j < len; j++ {
			if exists[data[j]] {
				continue
			}

			if unorderedEqual(v, data[j]) {
				res[v] = append(res[v], data[j])
				exists[data[j]] = true
			}
		}
		sort.Strings(res[v])
	}

	return res
}

func PrepareData(data []string) []string {
	res := make([]string, len(data))

	for i, v := range data {
		res[i] = strings.ToLower(v)
	}

	return res
}

func main() {
	data := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "СтоЛИК", "ПЯТКА", "раки", "КАИР", "ирак", "РАки"}
	data = PrepareData(data)
	res := Anagram(data)

	fmt.Println(res)
}
