package main

import "gopkg.in/yaml.v2"

type YAMLMarshaler struct{}

func (j *YAMLMarshaler) UnMarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func (j *YAMLMarshaler) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}
