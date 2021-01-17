package parser

import (
	"strings"

	"github.com/microcosm-cc/bluemonday"
)

const wordSeperator = " "

func GetWords(htmlBody string) []string {
	policy := bluemonday.NewPolicy()
	sanitizedText := policy.Sanitize(htmlBody)
	formattedText := standardizeSpaces(sanitizedText)
	return strings.Split(formattedText, wordSeperator)
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), wordSeperator)
}