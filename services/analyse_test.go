package services

import (
	"main/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyse(t *testing.T) {
	tests := []struct {
		name       string
		entries    []models.Entry
		categories map[string]float64
	}{
		{
			name: "One word match",
			entries: []models.Entry{
				{Line: "gas M1", Amount: 30},
				{Line: "tesco port solent", Amount: 12},
				{Line: "lidl north end", Amount: 13},
				{Line: "ma north end", Amount: 5},
			},
			categories: map[string]float64{"Utilities": 30, "Groceries": 25, "ma north end": 5},
		},
		{
			name: "Two word match",
			entries: []models.Entry{
				{Line: "john baker london road", Amount: 15},
				{Line: "tesco petrol havant", Amount: 11},
			},
			categories: map[string]float64{"Entertainment": 15, "Car": 11},
		},
	}

	for _, test := range tests {
		analyse := Analyse{}
		categories := analyse.Analyse(test.entries)
		assert.Equal(t, test.categories, categories, test.name+" failed")
	}
}
