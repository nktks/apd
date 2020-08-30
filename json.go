package main

import "encoding/json"

type JSONMarshaler struct{}

func (j *JSONMarshaler) UnMarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (j *JSONMarshaler) Marshal(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}
