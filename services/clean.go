package services

import (
	"errors"
	"main/models"
	"regexp"
	"strconv"
	"strings"
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

		amountRaw = strings.Trim(amountRaw, " ")
		conv := 1.0
		// If amount starts with a + then this means a refund, or negative spending
		if amountRaw[:1] == "+" {
			conv = -1
		}
		amountRaw = regexp.MustCompile(`[^0-9.]+`).ReplaceAllString(amountRaw, "") // Remove non-numeric characters
		line = regexp.MustCompile(`[.*]+`).ReplaceAllString(line, " ")             // Replace . and * with spaces

		var amount float64
		amount, err = strconv.ParseFloat(amountRaw, 64)
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
