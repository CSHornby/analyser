package services

import (
	"main/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClean(t *testing.T) {
	records := [][]string{
		{"01-04-2025", "line 1", " £50.00"},
		{"01-04-2025", "line 2", "£51.01"},
		{"01-04-2025", "line 3", " +£52.00"},
		{"01-04-2025", "line 4", "53.00"},
	}

	expCleaned := []models.Entry{
		{Line: "line 1", Amount: 50.0},
		{Line: "line 2", Amount: 51.01},
		{Line: "line 3", Amount: -52.0},
		{Line: "line 4", Amount: 53.0},
	}

	clean := Clean{}

	cleaned, err := clean.Clean(records)
	assert.Nil(t, err)

	assert.Equal(t, expCleaned, cleaned)
}

func TestCleanMissingField(t *testing.T) {
	records := [][]string{
		{"01-04-2025", "line 1"},
		{"01-04-2025", "line 2", " £51.01"},
		{"01-04-2025", "line 3", " +£52.00"},
		{"01-04-2025", "line 4", "53.00"},
	}

	clean := Clean{}

	_, err := clean.Clean(records)
	assert.NotNil(t, err)
}

func TestCleanBadData(t *testing.T) {
	records := [][]string{
		{"01-04-2025", "line 1", "text"},
		{"01-04-2025", "line 2", "£51.01"},
		{"01-04-2025", "line 3", " +£52.00"},
		{"01-04-2025", "line 4", "53.00"},
	}

	clean := Clean{}

	_, err := clean.Clean(records)
	assert.NotNil(t, err)
}
