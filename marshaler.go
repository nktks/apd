package main

import (
	"fmt"
	"log"
	"reflect"

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
	for i, r := range result {
		appended, ok := appendRecord(kf, m, r)
		if ok {
			result[i] = appended
		}
	}

	out, err := m.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func appendRecord(kf *KeyFinder, m Marshaler, r Record) (Record, bool) {
	fromKey := kf.DetectFromKey(r.keys())
	toKey := kf.DetectToKey(r.keys())
	if fromKey != "" && toKey != "" {
		from, err := dateparse.ParseAny(r[fromKey].(string))
		if err != nil {
			log.Print(fmt.Errorf("cant find from key: %w", err))
			return r, false
		}
		to, err := dateparse.ParseAny(r[toKey].(string))
		if err != nil {
			log.Print(fmt.Errorf("cant find to key: %w", err))
			return r, false
		}
		duration := to.Sub(from)
		r[kf.DurationKey] = duration.String()
	}
	for k, v := range r {
		rv := reflect.ValueOf(v)
		if rv.Type().Kind() == reflect.Map {
			rec, aok := v.(map[string]interface{})
			if aok {
				appended, bok := appendRecord(kf, m, Record(rec))
				if bok {
					r[k] = appended
				}
			}
			irec, iok := v.(map[interface{}]interface{})
			if iok {
				smap := map[string]interface{}{}
				for ik, iv := range irec {
					sk, sok := ik.(string)
					if sok {
						smap[sk] = iv
					}
				}
				appended, bok := appendRecord(kf, m, Record(smap))
				if bok {
					r[k] = appended
				}
			}
		}
		if rv.Type().Kind() == reflect.Slice {
			siface, ok := v.([]interface{})
			if ok {
				for sk, iface := range siface {
					rec, aok := iface.(map[string]interface{})
					if aok {
						appended, bok := appendRecord(kf, m, Record(rec))
						if bok {
							siface[sk] = appended
						}
					}
					irec, iok := iface.(map[interface{}]interface{})
					if iok {
						smap := map[string]interface{}{}
						for ik, iv := range irec {
							ssk, sok := ik.(string)
							if sok {
								smap[ssk] = iv
							}
						}
						appended, bok := appendRecord(kf, m, Record(smap))
						if bok {
							siface[sk] = appended
						}
					}
				}
			}
		}
	}
	return r, true
}

func (r Record) keys() []string {
	ks := []string{}
	for k, _ := range r {
		ks = append(ks, k)
	}
	return ks
}
