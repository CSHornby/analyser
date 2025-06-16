package services

import (
	"encoding/csv"
	"io"
)

type ExtractCsv struct {
}

type ExtractCsvI interface {
	Extract(file io.Reader) (records [][]string, err error)
}

func (e ExtractCsv) Extract(file io.Reader) (records [][]string, err error) {
	r := csv.NewReader(file)

	records, err = r.ReadAll()
	if err != nil {
		return nil, err
	}

	return
}
