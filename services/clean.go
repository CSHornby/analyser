package services

import (
	"errors"
	"main/models"
	"regexp"
	"strconv"
)

type Clean struct {
}

type CleanI interface {
	Clean(data [][]string) (entries []models.Entry, err error)
}

func (c Clean) Clean(data [][]string) (entries []models.Entry, err error) {
	for _, record := range data {
		if len(record) != 3 {
			return nil, errors.New("incorrect record length, expected 3 fields")
		}

		line := record[1]
		amountRaw := record[2]
		conv := 1.0
		if amountRaw[:1] == "+" {
			conv = -1
		}
		amountRaw = regexp.MustCompile(`[^0-9.]+`).ReplaceAllString(amountRaw, "") // Remove non-numeric characters
		line = regexp.MustCompile(`[.*]+`).ReplaceAllString(line, " ")             // Remove non-numeric characters

		var amount float64
		amount, err = strconv.ParseFloat(amountRaw, 32)
		if err != nil {
			return nil, errors.New("non number found in amount field")
		}

		entry := models.Entry{
			Line:   line,
			Amount: amount * conv,
		}
		entries = append(entries, entry)
	}

	return entries, nil
}
