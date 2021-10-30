package json

import (
	jsoniter "github.com/json-iterator/go"
)

var (
	Json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func Unmarshal(body []byte, v interface{}) error {
	return Json.Unmarshal(body, v)
}

func Marshal(v interface{}) ([]byte, error) {
	return Json.Marshal(v)
}
