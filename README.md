# WebScraping
Using Colly, a Go library for web scraping, and saving the data to a PostgreSQL database inside a Docker container.

## What is Web-Scraping?

Web scraping is when you use a computer program to extract information or data from websites. It works by downloading the HTML pages and pulling out the important information you want. People often use web scraping for things like finding prices, doing market research, analyzing data, or gathering content from different sources. However, it's important to remember that some websites don't allow web scraping or have rules about how it can be used, so it's important to check that you're not breaking any laws or being unethical.

## Technologies used

The construction of the algorithm was entirely done in Go, using the following libraries:

- [colly](https://github.com/gocolly/colly)

- [encoding/json](https://pkg.go.dev/encoding/json)

- [Docker](https://www.docker.com/)

- [PostgreSQL](https://www.postgresql.org/)

- [database/sql](https://github.com/lib/pq)
