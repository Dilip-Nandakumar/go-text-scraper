package parser

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

const wordSeperator = " "
const wordPairSeperator = "___"

func GetWords(htmlBody string) []string {
	policy := bluemonday.NewPolicy()
	sanitizedText := policy.Sanitize(htmlBody)
	formattedText := standardizeSpaces(sanitizedText)
	return strings.Split(formattedText, wordSeperator)
}

func GetWordPairs(htmlBody string) []string {
	words := GetWords(htmlBody)

	var wordPairs []string
	currentIndex := 0
	lastIndex := len(words) - 1

	for currentIndex < lastIndex {
		wordPairs = append(
			wordPairs,
			strings.Join(
				[]string{ words[currentIndex], words[currentIndex + 1] }, wordPairSeperator,
			),
		)

		currentIndex++
	}

	return wordPairs
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), wordSeperator)
}
