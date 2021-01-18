package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWords(t *testing.T) {
	words := GetWords("<div>word1 word2   word3<div>")
	expectedWords := []string{"word1", "word2", "word3"}

	assert.Equal(t, expectedWords, words)
}

func TestGetWordPairs(t *testing.T) {
	words := GetWordPairs("<div>word1 word2   word3<div>")
	expectedWords := []string{"word1___word2", "word2___word3"}

	assert.Equal(t, expectedWords, words)
}
