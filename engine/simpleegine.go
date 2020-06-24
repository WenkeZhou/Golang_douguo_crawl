package engine

import (
	"../fetcher"
	"fmt"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request

	for _, e := range seeds {
		requests = append(requests, e)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching url: %s\n", r.Url)
		body, err := fetcher.Fetch(r.Url)

		if err != nil {
			log.Printf("Fetch Error: %s\n", r.Url)
		}

		pareresult := r.ParseFunc(body)

		requests = append(requests, pareresult.Requests...)

		for _, item := range pareresult.Items {
			fmt.Printf("Got items: %s\n", item)
		}
	}
}
