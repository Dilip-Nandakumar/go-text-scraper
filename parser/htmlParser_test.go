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