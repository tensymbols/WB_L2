package main

import (
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var k int        // column index
	var n, r, u bool // by numeric value, reverse order, no duplicates
	var file *os.File
	flag.IntVar(&k, "k", 0, "column index")
	flag.BoolVar(&n, "n", false, "sort by numeric value")
	flag.BoolVar(&r, "r", false, "sort in reverse order")
	flag.BoolVar(&u, "u", false, "do not print duplicates")
	flag.Parse()
	if fName := flag.Arg(0); fName != "" {
		var err error
		file, err = os.Open(fName)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("filename is empty")
	}
	data, err := io.ReadAll(file)

	if err != nil {
		log.Fatal("could not read file contents")
	}
	var words [][]string
	for _, v := range strings.Split(string(data), "\n") {
		words = append(words, strings.Fields(v))
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i][k] < words[i][k]
	})
}
