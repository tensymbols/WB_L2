package main

import (
	"fmt"
	"sort"
	"strings"
)

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
func Contains[T comparable](arr []T, val T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func FindAnagrams(words []string) map[string][]string {
	dict := map[string][]string{}
	res := map[string][]string{}
	for _, w := range words {
		w = strings.ToLower(w)
		wr := []rune(w)
		sort.Slice(wr, func(i, j int) bool {
			return wr[i] < wr[j]
		})
		ws := string(wr)
		_, ok := dict[ws]
		if !ok {
			dict[ws] = []string{w}
		} else if !Contains(dict[ws], w) {
			dict[ws] = append(dict[ws], w)
		}
	}
	for _, v := range dict {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		if len(v) < 2 {
			continue
		}
		prev := ""
		var currSlice []string
		for _, val := range v {
			if val != prev {
				currSlice = append(currSlice, val)
			}
			prev = val
		}
		if len(currSlice) > 1 {
			res[currSlice[0]] = currSlice
		}
	}
	return res
}

func main() {
	words := []string{
		"ток", "кот", "Кот", "лес", "Сел", "слиток", "столик",
	}
	anagrams := FindAnagrams(words)
	for k, v := range anagrams {
		fmt.Println(k, v)
	}
}
