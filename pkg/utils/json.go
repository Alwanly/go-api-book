package utils

import "github.com/goccy/go-json"

func JsonMarshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func JsonUnMarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
