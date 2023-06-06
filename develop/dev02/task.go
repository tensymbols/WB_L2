package main

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func UnpackString(s string) (string, error) {
	var ret []rune
	var curr rune = -1
	for _, v := range s {
		if unicode.IsDigit(v) {
			if curr >= 0 {
				n, _ := strconv.Atoi(string(v))
				for i := 0; i < n-1; i++ {
					ret = append(ret, curr)
				}
				curr = -1
			} else {
				return string(ret), fmt.Errorf("invalid input")
			}
		} else {
			curr = v
			ret = append(ret, v)
		}

	}
	return string(ret), nil
}

func main() {
	fmt.Println(UnpackString("a4bc2d5e"))
	fmt.Println(UnpackString("abcd"))
	fmt.Println(UnpackString("45"))
}
