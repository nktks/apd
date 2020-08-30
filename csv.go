package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"strings"

	"github.com/araddon/dateparse"
)

func AppendCSV(kf *KeyFinder, input string) (string, error) {

	sr := strings.NewReader(input)
	r := csv.NewReader(sr)
	r.LazyQuotes = true

	records, err := r.ReadAll()
	if err != nil {
		return "", err
	}
	tsv := false
	if len(records) > 0 && strings.Contains(records[0][0], "\t") {
		tsv = true
		sr = strings.NewReader(input)
		r = csv.NewReader(sr)
		r.LazyQuotes = true
		r.Comma = '\t'
		if records, err = r.ReadAll(); err != nil {
			return "", err
		}
	}
	fromIndex := 0
	fromFound := false
	toIndex := 0
	toFound := false
	if len(records) > 0 {
		fromIndex, fromFound = kf.DetectFromIndex(records[0])
		toIndex, toFound = kf.DetectToIndex(records[0])
		if fromFound && toFound {
			records[0] = append(records[0], kf.DurationKey)
		}
	}

	for i, r := range records {
		if i == 0 {
			continue
		}
		if !fromFound || !toFound {
			continue
		}
		from, err := dateparse.ParseAny(r[fromIndex])
		if err != nil {
			log.Print(fmt.Errorf("cant parse time from key: %w", err))
			records[i] = append(records[i], "")
			continue
		}
		to, err := dateparse.ParseAny(r[toIndex])
		if err != nil {
			log.Print(fmt.Errorf("cant parse time to key: %w", err))
			records[i] = append(records[i], "")
			continue
		}
		duration := to.Sub(from)
		records[i] = append(records[i], duration.String())
	}
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if tsv {
		w.Comma = '\t'
	}

	w.WriteAll(records)
	ret := buf.String()
	if err := w.Error(); err != nil {
		return "", err
	}

	return ret, nil
}
