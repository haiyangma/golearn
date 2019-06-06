package engine

import (
	"crawler/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s \n", r.Url)
	body, err := fetcher.Fetch(r.Url)
	//time.Sleep(time.Duration(500) * time.Millisecond)
	if err != nil {
		log.Printf("Fetcher :error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parse(body, r.Url), nil
}
