package main

import (
	"github.com/Dilip-Nandakumar/word-highlights-scraper/config"
	"github.com/Dilip-Nandakumar/word-highlights-scraper/scraper"
	"github.com/Dilip-Nandakumar/word-highlights-scraper/utils"
	"github.com/Dilip-Nandakumar/word-highlights-scraper/parser"
	"github.com/Dilip-Nandakumar/word-highlights-scraper/aggregator"

	log "github.com/sirupsen/logrus"
)

func main() {
	utils.InitLogger()
	config := config.NewConfig()
	wordAggregator := aggregator.NewWordAggregator(10)

	log.Info("word highlights scraper has started")
	scrapResponse := make(chan string, 1)

	go func() {
		for response := range scrapResponse {
			log.Debugf("Response received: %s", response)
			words := parser.GetWords(response)
			wordAggregator.AggregateWords(words)
		}
	}()

	scraper.Scrap(scrapResponse, config.URL, config.Depth)
	frequentWords := wordAggregator.GetFrequentWords()

	for _, frequentWord := range frequentWords {
		log.Infof("word = %s, frequency = %d", frequentWord.Word, frequentWord.Frequency)
	}

	log.Info("word highlights scraper has completed")
}
