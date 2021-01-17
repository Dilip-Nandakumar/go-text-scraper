package scraper

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

func Scrap(response chan string, url string, depth int) {
	c := colly.NewCollector(
		colly.MaxDepth(depth),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		match, _ := regexp.MatchString(fmt.Sprintf("^%s", url), e.Attr("href"))

		if match {
			e.Request.Visit(e.Attr("href"))
		}
	})

	c.OnResponse(func(r *colly.Response) {
		response <- string(r.Body)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(url)

	log.Info("Closing the response channel...")
	close(response)
}
