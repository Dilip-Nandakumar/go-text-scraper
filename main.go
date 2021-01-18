package main

import (
	"sync"

	"github.com/Dilip-Nandakumar/text-scraper/config"
	"github.com/Dilip-Nandakumar/text-scraper/scraper"
	"github.com/Dilip-Nandakumar/text-scraper/utils"
	"github.com/Dilip-Nandakumar/text-scraper/parser"
	"github.com/Dilip-Nandakumar/text-scraper/aggregator"

	log "github.com/sirupsen/logrus"
)

func main() {
	utils.InitLogger()
	config := config.NewConfig()
	wordAggregator := aggregator.NewWordAggregator(10)
	wordPairAggregator := aggregator.NewWordAggregator(10)

	log.Info("Text scraper has started")
	scrapResponse := make(chan string, 1)
	var aggWaitGroup sync.WaitGroup

	go func() {
		aggWaitGroup.Add(1)
		defer aggWaitGroup.Done()
		for response := range scrapResponse {
			log.Debugf("Response received: %s", response)
			words := parser.GetWords(response)
			wordAggregator.AggregateWords(words)
			wordPairs := parser.GetWordPairs(response)
			wordPairAggregator.AggregateWords(wordPairs)
		}
	}()

	scraper.Scrap(scrapResponse, config.URL, config.Depth)
	aggWaitGroup.Wait()

	log.Info("Frequent words:")
	wordAggregator.LogFrequentWords()
	log.Info("Frequent word pairs:")
	wordPairAggregator.LogFrequentWords()

	log.Info("Text scraper has completed")
}
