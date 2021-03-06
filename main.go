package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	target := os.Getenv("TARGET_URL")
	// doc, err := goquery.NewDocument(target)

	reader := readLocalFile("root.html")
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		fmt.Print("url scraping failed")
	}

	branchs := []string{}

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		if !strings.Contains(url, "/") {
			branchs = append(branchs, url)
		}
	})

	for _, b := range branchs[:1] {
		fmt.Println(fmt.Sprintf("%s%s", target, b))
		d, err := goquery.NewDocument(fmt.Sprintf("%s%s", target, b))
		if err != nil {
			log.Fatal(err)
		}
		h, err := d.Html()
		fmt.Println(h)
		d.Find("a").Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			fmt.Println(url)
		})
	}
}

func writeLocalFile(filename, data string) error {
	b := []byte(data)
	err := ioutil.WriteFile(filename, b, 0666)
	if err != nil {
		return err
	}
	return nil
}

func readLocalFile(filename string) *strings.Reader {
	infos, _ := ioutil.ReadFile(filename)
	return strings.NewReader(string(infos))
}
