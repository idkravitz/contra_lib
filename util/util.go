package util

import (
	"os"
	"encoding/json"
	)

func Qjson(v interface{}) string {
	js, _ := json.Marshal(v)
	return string(js)
}

func GetenvDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return value
}