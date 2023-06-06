package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
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
	flag.IntVar(&k, "k", 1, "column index")
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
	file.Close()

	var words [][]string
	for _, v := range strings.Split(string(data), "\n") {

		tokens := strings.Fields(v)
		if len(tokens) < k {
			log.Fatal("invalid column index or there is an empty string")
		}
		words = append(words, tokens)
	}

	if n {
		sort.Slice(words, func(i, j int) bool {
			iv, err := strconv.Atoi(words[i][k-1])
			if err != nil {
				log.Fatal("non-numeric value was met")
			}
			jv, err := strconv.Atoi(words[j][k-1])
			if err != nil {
				log.Fatal("non-numeric value was met")
			}
			if r {
				return iv > jv
			}
			return iv < jv
		})
	} else {
		sort.Slice(words, func(i, j int) bool {
			if r {
				return words[i][k-1] > words[j][k-1]
			}
			return words[i][k-1] < words[j][k-1]
		})
	}
	fmt.Println(file, err)
	file, err = os.OpenFile(flag.Arg(0), os.O_RDWR, 0660)
	if err != nil {
		log.Fatal("could not open file for writing")
	}
	//	prev:=words[0]
	for i := range words {
		_, _ = file.WriteString(strings.Join(words[i], " "))
		if i < len(words)-1 {
			file.Write([]byte("\n"))
		}
	}
	fmt.Println(words)
	file.Close()
}
