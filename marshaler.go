package main

import (
	"fmt"
	"log"

	"github.com/araddon/dateparse"
)

type Marshaler interface {
	UnMarshal([]byte, interface{}) error
	Marshal(interface{}) ([]byte, error)
}

type Record map[string]interface{}
type Data []Record

func Append(kf *KeyFinder, m Marshaler, input string) (string, error) {
	var result Data
	if err := m.UnMarshal([]byte(input), &result); err != nil {
		return "", err
	}
	fromKey := ""
	toKey := ""
	if len(result) > 0 {
		fromKey = kf.DetectFromKey(result[0].keys())
		toKey = kf.DetectToKey(result[0].keys())
	}
	for i, r := range result {
		if fromKey == "" || toKey == "" {
			continue
		}
		from, err := dateparse.ParseAny(r[fromKey].(string))
		if err != nil {
			log.Print(fmt.Errorf("cant find from key: %w", err))
			continue
		}
		to, err := dateparse.ParseAny(r[toKey].(string))
		if err != nil {
			log.Print(fmt.Errorf("cant find to key: %w", err))
			continue
		}
		duration := to.Sub(from)
		result[i][kf.DurationKey] = duration.String()
	}

	out, err := m.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (r Record) keys() []string {
	ks := []string{}
	for k, _ := range r {
		ks = append(ks, k)
	}
	return ks
}
