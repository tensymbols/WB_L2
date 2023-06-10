package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Cut(str, d *string, f *int, s *bool) string {
	if *s && !strings.Contains(*str, *d) {
		return ""
	}

	strArr := strings.Split(*str, *d)

	if *f > len(strArr) {
		return *str + "\n"
	} else {
		return strArr[*f-1] + "\n"
	}
}

func main() {

	var f int
	var d string
	var s bool

	flag.IntVar(&f, "f", 0, "choose fields(columns)")
	flag.StringVar(&d, "d", "\t", "use different delimiter")
	flag.BoolVar(&s, "s", false, "only lines with delimiter")

	flag.Parse()

	if f == 0 {
		log.Fatalln("wrong input")
	}

	for {
		buf, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			break
		}

		str := strings.Trim(buf, "\n")
		res := Cut(&str, &d, &f, &s)
		fmt.Print(res)
	}

}
