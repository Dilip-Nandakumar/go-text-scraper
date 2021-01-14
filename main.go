package main

import (
	"github.com/Dilip-Nandakumar/word-highlights-scraper/config"
	"github.com/Dilip-Nandakumar/word-highlights-scraper/scraper"
	"github.com/Dilip-Nandakumar/word-highlights-scraper/utils"

	log "github.com/sirupsen/logrus"
)

func main() {
	utils.NewLogger()
	config := config.NewConfig()

	log.Info("word highlights scraper has started")
	scrapResponse := make(chan []byte, 1)

	go func() {
		for response := range scrapResponse {
			log.Debugf("Response received: %s", response)
		}
	}()

	scraper.Scrap(scrapResponse, config.URL, config.Depth)

	log.Info("word highlights scraper has completed")
}
