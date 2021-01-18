package aggregator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/Dilip-Nandakumar/text-scraper/utils"
)

func TestUpdateFrequentWords(t *testing.T) {
	aggregator := NewWordAggregator(1)

	aggregator.updateFrequentWords("a", 1)

	frequentWord := aggregator.frequentWordsStore.Pop().(frequentWord)

	assert.Equal(t, "a", frequentWord.Word)
	assert.Equal(t, 1, frequentWord.Frequency)
}

func TestAggregateWords(t *testing.T) {
	utils.InitLogger()
	aggregator := NewWordAggregator(5)

	aggregator.AggregateWords([]string {
		"a", "a", "a", "b", "b", "c", "c", "c", "c", "d", "d", "e", "f",
	})

	frequentWords := aggregator.GetFrequentWords()

	// for _, frequentWord := range frequentWords {
	// 	t.Logf("%s__%d", frequentWord.word, frequentWord.frequency)
	// }

	expectedWords := []frequentWord{
		frequentWord{"f", 1},
		frequentWord{"d", 2},
		frequentWord{"b", 2},
		frequentWord{"a", 3},
		frequentWord{"c", 4},
	}

	assert.Equal(t, expectedWords, frequentWords)
}
