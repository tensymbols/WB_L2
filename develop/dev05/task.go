package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func addBefore(lines []string, res []string, currIx int, N int) []string {
	if N < currIx {
		return append(res, lines[currIx-N:currIx]...)
	} else {
		return append(res, lines[:currIx]...)
	}
}

func addAfter(lines []string, res []string, currIx int, N int) []string {
	if N < len(lines)-currIx-1 {
		return append(res, lines[currIx+1:currIx+1+N]...)
	} else {
		return append(res, lines[currIx+1:]...)
	}
}

type grepSettings struct {
	A, B, C, c int
	i, v, n, F bool
}

func main() {
	gs := grepSettings{}

	flag.IntVar(&gs.A, "A", 0, "after +N lines")
	flag.IntVar(&gs.B, "B", 0, "before +N lines")
	flag.IntVar(&gs.C, "C", 0, "context, print A+B lines")
	flag.IntVar(&gs.c, "c", -1, "count,  lines quantity ")

	flag.BoolVar(&gs.i, "i", false, "ignore case")
	flag.BoolVar(&gs.v, "v", false, "invert, print everything except matching lines")
	flag.BoolVar(&gs.n, "n", false, "line num, print line number")
	flag.BoolVar(&gs.F, "F", false, "explicit match, not a pattern match")
	flag.Parse()
	var file *os.File
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
	pattern := flag.Arg(1)
	if gs.i {
		pattern = strings.ToLower(pattern)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("could not read from file")
	}
	lines := strings.Split(string(data), "\n")
	var res []string

	for i := 0; i < len(lines) && gs.c != 0; i++ {
		curr := ""
		var match bool
		if gs.i {
			match = gs.F && strings.ToLower(lines[i]) == pattern || !gs.F && strings.Contains(strings.ToLower(lines[i]), pattern)
		} else {
			match = gs.F && lines[i] == pattern || !gs.F && strings.Contains(lines[i], pattern)
		}
		if match && !gs.v {
			if gs.n {
				curr += "line " + strconv.Itoa(i) + ":"
			}
			if gs.C > 0 {
				curr += strings.Join(addBefore(lines, []string{}, i, gs.A), "\n")
				curr += lines[i] + "\n"
				curr += strings.Join(addAfter(lines, []string{}, i, gs.B), "\n")
			} else {
				if gs.A > 0 {
					curr += strings.Join(addBefore(lines, []string{}, i, gs.A), "\n")
				}
				curr += lines[i] + "\n"
				if gs.B > 0 {
					curr += strings.Join(addAfter(lines, []string{}, i, gs.B), "\n")
				}
			}

		} else if !match && gs.v {
			if gs.n {
				curr += "line " + strconv.Itoa(i) + ":"
			}
			curr += lines[i] + "\n"
		}
		res = append(res, curr)
		gs.c--
	}
	for _, v := range res {
		fmt.Print(v)
	}
}
