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

	doc, err := goquery.NewDocument(target)
	if err != nil {
		fmt.Print("url scraping failed")
	}
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		if !strings.Contains(url, "/") {
			fmt.Println(url)
		}
	})
}

func writeLocalFile(filename, data string) error {
	b := []byte(data)
	err := ioutil.WriteFile(filename, b, 0666)
	if err != nil {
		return err
	}
	return nil
}
