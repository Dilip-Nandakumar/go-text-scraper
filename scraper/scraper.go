package scraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

func Scrap(response chan []byte, url string, depth int) {
	c := colly.NewCollector(
		colly.MaxDepth(depth),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnResponse(func(r *colly.Response) {
		response <- r.Body
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	log.Info("Closing the response channel...")
	close(response)
}
