package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sjatsh/grab"
	"log"
	"net/http"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	m := flag.Bool("m", false, "mirror - download site")
	flag.Parse()

	url := flag.Arg(0)

	dir := "download/"
	if !*m {
		resp, err := grab.Get(dir, url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Download saved to", resp.Filename)
	} else {
		links := parseLink(url)

		for _, val := range links {
			resp, err := grab.Get(dir, url+val)
			if err != nil {
				fmt.Println("Error download page")
			} else {
				fmt.Println("Download saved to", resp.Filename)
			}
		}

	}

}

func parseLink(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var links []string

	doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		link, _ := linkTag.Attr("href")
		links = append(links, link)
	})
	return links
}
