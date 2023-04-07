package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
	_ "github.com/lib/pq"
)

type Quote struct {
	Quote  string `json: "quote"`
	Author string `json: "author"`
}

func main() {
	quotes := []Quote{}

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=scrapdatabase  sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

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

	c.Visit("http://quotes.toscrape.com/page/2/")

	for _, q := range quotes {
		_, err = db.Exec("INSERT INTO quotes (quote, author) VALUES ($1, $2)", q.Quote, q.Author)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Dados inseridos com sucesso!")

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
