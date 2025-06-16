package services

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractCSV(t *testing.T) {
	file, _ := os.Open("./../test-files/statement.csv")
	extractCsv := ExtractCsv{}

	records, err := extractCsv.Extract(file)
	assert.Nil(t, err)

	expRecords := [][]string{
		{"13-04-2025", "Line 3", " £22.45"},
		{"14-04-2025", "Line 5", " +£23.45"},
	}

	assert.Equal(t, expRecords, records)
}

func TestExtractCSVBad(t *testing.T) {
	file, _ := os.Open("./../test-files/bad-statement.csv")
	extractCsv := ExtractCsv{}

	_, err := extractCsv.Extract(file)
	assert.NotNil(t, err)
}
