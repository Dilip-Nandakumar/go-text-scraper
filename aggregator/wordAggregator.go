package aggregator

import (
	"container/heap"

	log "github.com/sirupsen/logrus"
)

type wordAggregator struct {
	frequentWordsStore  FrequentWordStore
	capacity            int
	wordAccumulator     map[string]int
}

func NewWordAggregator(capacity int) wordAggregator {
	return wordAggregator{
		frequentWordsStore:  *new(FrequentWordStore),
		capacity:            capacity,
		wordAccumulator: make(map[string]int),
	}
}

func (aggregator *wordAggregator) AggregateWords(words []string) {
	for _, word := range words {
		if frequency, ok := aggregator.wordAccumulator[word]; ok {
			aggregator.wordAccumulator[word] = frequency + 1
		} else {
			aggregator.wordAccumulator[word] = 1
		}

		aggregator.updateFrequentWords(word, aggregator.wordAccumulator[word])
	}
}

func (aggregator *wordAggregator) GetFrequentWords() []frequentWord {
	var words []frequentWord

	for aggregator.frequentWordsStore.Len() > 0 {
		words = append(words, heap.Pop(&aggregator.frequentWordsStore).(frequentWord))
	}

	return words
}


func (aggregator *wordAggregator) LogFrequentWords() {
	frequentWords := aggregator.GetFrequentWords()

	for _, frequentWord := range frequentWords {
		log.Infof("word = %s, frequency = %d", frequentWord.Word, frequentWord.Frequency)
	}
}

type frequentWord struct {
	Word      string
	Frequency int
}

type FrequentWordStore []frequentWord

func (words *FrequentWordStore) Less(i, j int) bool {
	return (*words)[i].Frequency < (*words)[j].Frequency
}

func (words *FrequentWordStore) Swap(i, j int) {
	(*words)[i], (*words)[j] = (*words)[j], (*words)[i]
}

func (words *FrequentWordStore) Len() int {
	return len(*words)
}

func (words *FrequentWordStore) Pop() (v interface{}) {
	*words, v = (*words)[:words.Len()-1], (*words)[words.Len()-1]
	return
}

func (words *FrequentWordStore) Push(v interface{}) {
	*words = append(*words, v.(frequentWord))
}

func (aggregator *wordAggregator) updateFrequentWords(word string, frequency int) {
	if aggregator.frequentWordsStore.Len() == aggregator.capacity && aggregator.frequentWordsStore[0].Frequency > frequency {
		return
	}

	if aggregator.frequentWordsStore.Len() == aggregator.capacity {
		heap.Pop(&aggregator.frequentWordsStore)
	}

	if index, ok := aggregator.getWordIndexInStore(word); ok {
		aggregator.frequentWordsStore[index].Frequency++
	} else {
		log.Debugf("Pushing word %s with frequency %d to word store", word, frequency)
		fw := frequentWord{word, frequency}
		heap.Push(&aggregator.frequentWordsStore, fw)
	}
}

func (aggregator *wordAggregator) getWordIndexInStore(word string) (int, bool) {
	for i, frequentWord := range aggregator.frequentWordsStore {
		if frequentWord.Word == word {
			return i, true
		}
	}

	return 0, false
}
