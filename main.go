package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

type Quote struct {
	Quote  string `json: "quote"`
	Author string `json: "author"`
}

func main() {
	quotes := []Quote{}

	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response code", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error", err.Error())
	})

	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		quote := div.Find(".text").Text()
		author := div.Find(".author").Text()

		q := Quote{
			Quote:  quote,
			Author: author,
		}
		quotes = append(quotes, q)
	})

	c.Visit("http://quotes.toscrape.com")

	WriteJsonFile(quotes)

}

func WriteJsonFile(quotes []Quote) {
	file, err := os.Create("quotes.json")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jsonBytes, err := json.MarshalIndent(quotes, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	file.Write(jsonBytes)
}
