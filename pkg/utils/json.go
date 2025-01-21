package utils

import "github.com/goccy/go-json"

func JSONMarshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func JSONUnMarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
