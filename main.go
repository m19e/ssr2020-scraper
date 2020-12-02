package main

import (
	"fmt"
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
	url := os.Getenv("TARGET_URL")

	doc, err := goquery.NewDocument(url)
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
