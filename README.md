# WebScraping
Using Colly, a Go library for web scraping, and saving the data to a Local PostgreSQL database inside a Docker container.

## What is Web-Scraping?

Web scraping is when you use a computer program to extract information or data from websites. It works by downloading the HTML pages and pulling out the important information you want. People often use web scraping for things like finding prices, doing market research, analyzing data, or gathering content from different sources. However, it's important to remember that some websites don't allow web scraping or have rules about how it can be used, so it's important to check that you're not breaking any laws or being unethical.

## Technologies used

The construction of the algorithm was entirely done in Go, using the following libraries:

- [colly](https://github.com/gocolly/colly)

- [encoding/json](https://pkg.go.dev/encoding/json)

- [Docker](https://www.docker.com/)

- [PostgreSQL](https://www.postgresql.org/)

- [database/sql](https://github.com/lib/pq)

## Pull postgres image:
```
 docker pull postgres
```

## Initialize a container

```
docker run --name mypsql-container -d -p 5432:5432 -e POSTGRES_PASSWORD=example postgres
```
This command creates and runs a Docker container named mypsql-container with the official PostgreSQL image postgres. The -d flag runs the container in detached mode, allowing it to run in the background. The -p flag maps the container's internal port 5432 to the host's port 5432, allowing connections to the PostgreSQL database from outside the container.

The -e flag sets an environment variable named POSTGRES_PASSWORD to the value example. This sets the password for the default user postgres in the PostgreSQL instance running inside the container.

## Open bash or Docker Desktop [terminal]

### SQL script

```
CREATE DATABASE scrapdatabase;

\c

CREATE TABLE quotes(
    id SERIAL PRIMARY KEY,
    quote VARCHAR,
    author VARCHAR
);
    
```

## Website to Scrap

http://quotes.toscrape.com/
